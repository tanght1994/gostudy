package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"sort"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

// wpt 是 web performance test 的缩写

var (
	g_Concurrency = int64(0) // 并发量, 命令行参数 -c
	g_CountLimit  = int64(0) // 总请求数, 命令行参数 -n
	g_TimeOut     = int64(0) // http请求超时时间, 命令行参数 -t
	g_ReqFilePath = ""       // 请求文件的路径, 命令行参数 -p
)

var (
	g_StatMutex       = sync.Mutex{}
	g_StatStatusCode  = map[int]int{} // 保存每次请求的状态码 {状态码: 出现次数}
	g_StatRespTimes   = []int64{}     // 保存每次请求的响应时间, 单位是毫秒
	g_StatBeginCnt    = int64(0)      // 请求开始数量
	g_StatCompleteCnt = int64(0)      // 请求完成数量
	g_StatTotalTime   = float64(0)    // 压测开始至结束的总时间
	g_StatDataSize    = int64(0)      // 从服务器收到的数据总量, 单位是字节
)

func main() {
	flag.Int64Var(&g_Concurrency, "c", 1, "并发量")
	flag.Int64Var(&g_CountLimit, "n", 1, "总请求数, 到达此数量时压测会停止")
	flag.Int64Var(&g_TimeOut, "t", 3000, "请求超时时间, 单位为毫秒")
	flag.StringVar(&g_ReqFilePath, "p", "req.txt", "request文件的路径")
	flag.Parse()

	wg := sync.WaitGroup{}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// 监听ctrl+c
	wg.Add(1)
	go signalListen(ctx, func() {
		cancel()
	}, &wg)

	// 解析请求文件
	data := parseReqFile(g_ReqFilePath)

	// 单独协程每秒打印请求情况
	wg.Add(1)
	go printProcessingInfo(ctx, &wg)

	// 开始压测
	t0 := time.Now()
	allPeople(ctx, int(g_Concurrency), data)
	g_StatTotalTime = time.Since(t0).Seconds()

	// 压测结束, 让所有辅助协程都停止
	cancel()
	// 等待 所有协程 结束
	wg.Wait()

	// 输出统计结果
	printStatInfo()
}

type ReqData struct {
	method string
	url    string
	body   string
	header map[string]string
}

func must(msg string, err error) {
	if err != nil {
		fmt.Printf("%s, error is %s\n", msg, err.Error())
		os.Exit(1)
	}
}

func readLines(path string) []string {
	lines := make([]string, 0)
	f, e := os.Open(path)
	must(fmt.Sprintf("open file error, path is %s", path), e)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func makeRequests(data []string) []ReqData {
	data = append(data, "", "", "", "")
	regexp1, err := regexp.Compile(`\S`)
	must("regexp compile error", err)

	var stat int = 0
	var statFirstLine int = 0
	var statHeader int = 1
	var statBody int = 2

	requests := make([]ReqData, 0)
	reqdata := ReqData{}
	reqdata.header = map[string]string{}

	for i := range data {
		// 如果这一行全是不可见字符, 直接把它变成空字符串
		if !regexp1.Match([]byte(data[i])) {
			data[i] = ""
		}

		line := []byte(data[i])

		switch stat {
		case statFirstLine:
			if len(line) == 0 {
				continue
			}
			tmp := bytes.Split(line, []byte{' '})
			if len(tmp) < 2 {
				fmt.Println("request data parse error 1")
				os.Exit(1)
			}
			reqdata.method = string(tmp[0])
			reqdata.url = string(tmp[1])
			stat = statHeader
			continue
		case statHeader:
			if len(line) == 0 {
				stat = statBody
				continue
			}
			tmp := bytes.SplitN(line, []byte{':'}, 2)
			if len(tmp) < 2 {
				fmt.Println("request data parse error 2")
				os.Exit(1)
			}
			reqdata.header[string(tmp[0])] = string(tmp[1])
		case statBody:
			if len(line) == 0 {
				stat = statFirstLine
				requests = append(requests, reqdata)
				reqdata = ReqData{}
				reqdata.header = map[string]string{}
				continue
			}
			if len(reqdata.body) == 0 {
				reqdata.body = reqdata.body + string(line)
			} else {
				reqdata.body = reqdata.body + "\n" + string(line)
			}
		}
	}
	return requests
}

func parseReqFile(path string) []ReqData {
	lines := readLines(path)
	return makeRequests(lines)
}

func onePeople(ctx context.Context, requests []ReqData, wg *sync.WaitGroup, start <-chan struct{}) {
	defer wg.Done()
	// 准备客户端
	tp := http.Transport{}
	tp.DisableKeepAlives = true
	tlscfg := tls.Config{InsecureSkipVerify: true}
	tp.TLSClientConfig = &tlscfg
	client := http.Client{}
	client.Transport = &tp
	client.Timeout = time.Duration(g_TimeOut) * time.Millisecond

	// 准备统计
	costsTimeList := []int64{}
	statusCodeMap := make(map[int]int)

	//等待开始
	<-start

LOOP1:
	for i := 0; true; i++ {

		select {
		case <-ctx.Done():
			break LOOP1
		default:
		}

		if i == len(requests) {
			i = 0
		}
		// 制作请求
		req := requests[i]
		body := bytes.NewBufferString(req.body)
		request, err := http.NewRequest(req.method, req.url, body)
		request.Close = true // 不支持keep-alive
		must("http.NewRequest error", err)
		for k, v := range req.header {
			request.Header.Add(k, v)
		}

		// 是否达到结束条件, 并且请求开始个数+1
		beginCnt := atomic.AddInt64(&g_StatBeginCnt, 1)
		if beginCnt > g_CountLimit {
			atomic.AddInt64(&g_StatBeginCnt, -1)
			break
		}
		// 发送请求
		t0 := time.Now()
		res, err := client.Do(request)

		// 处理响应
		if err != nil {
			// 状态码-1代表未知错误
			// 可能是客户端的错误也可能是服务端的错误
			statusCodeMap[-1]++
		} else {
			dataSize := 0
			// 统计header的长度
			for k, v := range res.Header {
				dataSize += len(k)
				for _, tmp := range v {
					dataSize += len(tmp)
				}
			}

			if res.StatusCode != 200 {
				// nothing to do
			} else {
				data, _ := ioutil.ReadAll(res.Body)
				dataSize += len(data)
				// 只统计状态码为200的耗时
				costTime := time.Since(t0).Milliseconds()
				costsTimeList = append(costsTimeList, costTime)
			}
			atomic.AddInt64(&g_StatDataSize, int64(dataSize))
			statusCodeMap[res.StatusCode]++
			res.Body.Close()
		}

		// 完成个数+1
		atomic.AddInt64(&g_StatCompleteCnt, 1)
	}

	// 添加到全局统计
	g_StatMutex.Lock()
	defer g_StatMutex.Unlock()
	for k, v := range statusCodeMap {
		g_StatStatusCode[k] += v
	}
	g_StatRespTimes = append(g_StatRespTimes, costsTimeList...)
}

func allPeople(ctx context.Context, peopleCnt int, reqDataList []ReqData) {
	wg := sync.WaitGroup{}
	wg.Add(peopleCnt)
	start := make(chan struct{})
	for i := 0; i < peopleCnt; i++ {
		go onePeople(ctx, reqDataList, &wg, start)
	}
	close(start)
	wg.Wait()
}

func printProcessingInfo(ctx context.Context, wg *sync.WaitGroup) {
	// 每秒输出已完成的个数
	defer wg.Done()
	ticker := time.NewTicker(time.Second)
	lastCompletCnt := int64(0)

	printInfo := func() {
		beginCnt := atomic.LoadInt64(&g_StatBeginCnt)
		completCnt := atomic.LoadInt64(&g_StatCompleteCnt)
		ingCnt := beginCnt - completCnt
		diffCnt := completCnt - lastCompletCnt
		remainCnt := atomic.LoadInt64(&g_CountLimit) - beginCnt
		now := time.Now().Format("15:04:05.000")
		format := "%v   已开始%-10d已完成%-10d进行中%-10d比上一次多完成%-10d剩余%d\n"
		fmt.Printf(format, now, beginCnt, completCnt, ingCnt, diffCnt, remainCnt)
		lastCompletCnt = completCnt
	}

	printInfo()

	goon := true
	for goon {
		select {
		case <-ctx.Done():
			goon = false
		case <-ticker.C:
			// nothing
		}
		printInfo()
	}
}

func printStatInfo() {
	actualCnt := 0
	codes := []int{}
	for k, v := range g_StatStatusCode {
		codes = append(codes, k)
		actualCnt += v
	}
	sort.Ints(codes)

	fmt.Println("")
	fmt.Printf("计划发送%d个请求\n", g_CountLimit)
	fmt.Printf("实际发送%d个请求\n", actualCnt)
	fmt.Printf("总耗时%.2f秒\n", g_StatTotalTime)
	fmt.Printf("平均QPS%d\n", int(float64(actualCnt)/g_StatTotalTime))
	fmt.Printf("总接收数据量%s(大致)\n", formatDataSize(int(g_StatDataSize))) // header数据可能不准确, 且未统计FirstLine
	fmt.Println("")
	fmt.Println("状态码    个数           百分比%")
	for _, code := range codes {
		cnt := g_StatStatusCode[code]
		per := float64(cnt) / float64(actualCnt) * 100
		fmt.Printf("%-10d%-15d%.2f\n", code, cnt, per)
	}
	fmt.Println("")
	printTimeCostProportion()
}

func signalListen(ctx context.Context, action func(), wg *sync.WaitGroup) {
	defer wg.Done()
	// 接收到SIGINT信号时执行action函数
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	select {
	case <-c:
		action()
	case <-ctx.Done():
		// nothing
	}
}

func formatDataSize(size int) string {
	// byte 转 KB MB GB
	s := float64(size)
	if s < 1024.0 {
		return fmt.Sprintf("%.0fB", s)
	}
	s = s / 1024.0
	if s < 1024.0 {
		return fmt.Sprintf("%.2fKB", s)
	}
	s = s / 1024.0
	if s < 1024.0 {
		return fmt.Sprintf("%.2fMB", s)
	}
	s = s / 1024.0
	return fmt.Sprintf("%.2fGB", s)
}

func printTimeCostProportion() {
	// 打印响应时间的占比
	sort.Slice(g_StatRespTimes, func(i, j int) bool { return g_StatRespTimes[i] < g_StatRespTimes[j] })
	proportions := []float64{0.5, 0.6, 0.7, 0.8, 0.9}
	proportions = append(proportions, 0.91, 0.92, 0.93, 0.94, 0.95, 0.96, 0.97, 0.98, 0.99)
	proportions = append(proportions, 0.991, 0.992, 0.993, 0.994, 0.995, 0.996, 0.997, 0.998, 0.999)
	l := len(g_StatRespTimes)
	for _, v := range proportions {
		idx := int(float64(l)*v) - 1
		if idx < 0 {
			idx = 0
		}
		respTime := g_StatRespTimes[idx]
		fmt.Printf("%.1f%% 响应时间小于 %dms\n", v*100, respTime)
	}
	fmt.Printf("100%% 响应时间小于 %dms\n", g_StatRespTimes[len(g_StatRespTimes)-1])
}
