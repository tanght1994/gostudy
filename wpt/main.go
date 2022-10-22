package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

// wpt 是 web performance test 的缩写

func main() {
	c := flag.Int("c", 0, "并发量")
	p := flag.String("p", "req.txt", "request文件的路径")
	flag.Parse()
	data := parseReqFile(*p)

	// 单独协程每秒打印请求情况
	wg := sync.WaitGroup{}
	stop := make(chan struct{})
	totalCnt := *c * len(data)
	wg.Add(1)
	go printProcessingInfo(&wg, stop, int64(totalCnt))

	// 开始压测
	t0 := time.Now()
	allPeople(*c, data)
	StatTotalTime = time.Since(t0).Seconds()

	// 通知 printProcessingInfo 结束
	close(stop)
	// 等待 printProcessingInfo 结束
	wg.Wait()

	// 输出统计结果
	printStatInfo(totalCnt)
}

type ReqData struct {
	method string
	url    string
	body   string
	header map[string]string
}

var (
	StatMutex       = sync.Mutex{}
	StatStatusCode  = map[int]int{}
	StatTimeCost    = []int64{}
	StatCompleteCnt = int64(0)
	StatTotalTime   = float64(0)
	StatDataSize    = int64(0)
)

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

func onePeople(requests []ReqData, wg *sync.WaitGroup, start <-chan struct{}) {
	defer wg.Done()
	// 准备客户端
	tp := http.Transport{}
	tp.DisableKeepAlives = true
	tlscfg := tls.Config{InsecureSkipVerify: true}
	tp.TLSClientConfig = &tlscfg
	client := http.Client{}
	client.Transport = &tp

	// 准备统计
	costsTimeList := []int64{}
	statusCodeMap := make(map[int]int)

	//等待开始
	<-start

	for i := 0; i < len(requests); i++ {
		// 制作请求
		req := requests[i]
		body := bytes.NewBufferString(req.body)
		request, err := http.NewRequest(req.method, req.url, body)
		request.Close = true // 不支持keep-alive
		must("http.NewRequest error", err)
		for k, v := range req.header {
			request.Header.Add(k, v)
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
			if res.StatusCode != 200 {
				// nothing to do
			} else {
				data, _ := ioutil.ReadAll(res.Body)
				atomic.AddInt64(&StatDataSize, int64(len(data)))
				// 只统计状态码为200的耗时
				costTime := time.Since(t0).Milliseconds()
				costsTimeList = append(costsTimeList, costTime)
			}
			statusCodeMap[res.StatusCode]++
			res.Body.Close()
		}

		// 完成个数+1
		atomic.AddInt64(&StatCompleteCnt, 1)
	}

	// 添加到全局统计
	StatMutex.Lock()
	defer StatMutex.Unlock()
	for k, v := range statusCodeMap {
		StatStatusCode[k] += v
	}
	StatTimeCost = append(StatTimeCost, costsTimeList...)
}

func allPeople(peopleCnt int, reqDataList []ReqData) {
	wg := sync.WaitGroup{}
	wg.Add(peopleCnt)
	start := make(chan struct{})
	for i := 0; i < peopleCnt; i++ {
		go onePeople(reqDataList, &wg, start)
	}
	close(start)
	wg.Wait()
}

func printProcessingInfo(wg *sync.WaitGroup, stop <-chan struct{}, totalReqCnt int64) {
	// 每秒输出已完成的个数
	defer wg.Done()
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-stop:
			return
		case <-ticker.C:
			fmt.Println(atomic.LoadInt64(&StatCompleteCnt), "/", totalReqCnt)
		}
	}
}

func printStatInfo(expectCnt int) {
	actualCnt := 0
	codes := []int{}
	for k, v := range StatStatusCode {
		codes = append(codes, k)
		actualCnt += v
	}
	sort.Ints(codes)

	fmt.Printf("计划发送%d个请求\n", expectCnt)
	fmt.Printf("实际发送%d个请求\n", actualCnt)
	fmt.Println("")
	fmt.Println("状态码    个数           百分比%")
	for _, code := range codes {
		cnt := StatStatusCode[code]
		per := float64(cnt) / float64(actualCnt) * 100
		fmt.Printf("%-10d%-15d%.2f\n", code, cnt, per)
	}
	fmt.Println("")

	sort.Slice(StatTimeCost, func(i, j int) bool { return StatTimeCost[i] < StatTimeCost[j] })

	l := len(StatTimeCost)
	if l < 10 {
		for _, v := range StatTimeCost {
			fmt.Printf("%dms\n", v)
		}
	} else {
		// 对l向下取整  l=123  =>  l=120
		l = l / 10 * 10
		a := l / 10
		for i := 1; i < 10; i++ {
			fmt.Printf("%3d%% 的请求响应时间小于 %dms\n", i*10, StatTimeCost[a*i])
		}
		fmt.Printf("%3d%% 的请求响应时间小于 %dms\n", 100, StatTimeCost[len(StatTimeCost)-1])
	}
}
