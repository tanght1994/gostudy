package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"sort"
	"sync"
	"time"

	// _ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func abc() {
	a := make([]int, 1024*1024)
	for i := range a {
		a[i] = rand.Int()
	}
}

func memsetLoop(a []byte, v byte) {
	for i := range a {
		a[i] = v
	}
}

var (
	data1 = `
	{
		"game_package_id":"Fishing",
		"plat_id":"1",
		"country_id":"9999",
		"channel_id":"999",
		"client_app_version":"1.5.0",
		"client_res_version":"504"
	}
	`
)

type PostData struct {
	url  string
	data string
}

var postData []PostData = []PostData{
	{"http://18.139.76.91:9091/v2/oversea_app_update/app_update_info/", `{"game_package_id":"Fishing","plat_id":"1","country_id":"9999","channel_id":"999","client_app_version":"1.5.0","client_res_version":"504"}`},
	{"http://18.139.76.91:9091/v2/app_mobile/get_package_depend_info/", `{"game_package_id": "DominoCS.game","plat_id": 8,"country_id": 8,"channel_id": 8,"app_version": "0"}`},
	{"http://18.139.76.91:9091/v1/oversea_app_update/apk_update_info/", `{"apk_id": "预留字段","plat_id": "1","country_id": "0360","channel_id": "001"}`},
	{"http://18.139.76.91:9091/v1/oversea_app_update/game_server_addr/", `{"app_version":"8","res_version":"8","country_id":"8","channel_id":"8"}`},
}

var done int64
var start chan struct{}
var costs map[int64]int
var mutex sync.Mutex

func onepeople(cnt int, wg *sync.WaitGroup) {
	defer wg.Done()
	<-start
	for i := 0; i < cnt; i++ {
		idx := rand.Intn(4)
		url := postData[idx].url
		data := postData[idx].data
		bf := bytes.NewBufferString(data)
		t0 := time.Now()
		res, err := http.Post(url, "application/json", bf)
		if err != nil {
			fmt.Println(err)
			time.Sleep(50 * time.Millisecond)
			continue
		}
		cost := time.Since(t0).Milliseconds()
		time.Sleep(50 * time.Millisecond)
		mutex.Lock()
		costs[cost]++
		mutex.Unlock()
		if res.StatusCode != 200 {
			fmt.Println(res.Status)
		} else {
			a, err := ioutil.ReadAll(res.Body)
			must(err)
			if a[0] != '{' {
				fmt.Println("!= {")
			}
		}
		res.Body.Close()
	}
}

func allpeople(people, percnt int) {
	costs = make(map[int64]int)
	start = make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(people)
	for i := 0; i < people; i++ {
		go onepeople(percnt, &wg)
	}
	close(start)
	wg.Wait()
	keys := []int64{}
	for k := range costs {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	for _, v := range keys {
		fmt.Printf("%d %d\n", v, costs[v])
	}
}

func main() {
	t0 := time.Now()
	allpeople(100, 50)
	fmt.Printf("cost %v", time.Now().Sub(t0))
	return

	conn, err := net.Dial("tcp", "www.tanght.xyz:20001")
	must(err)
	// fmt.Println(conn.Write([]byte("GET /hello HTTP/1.1\r\nConnection: keep-alive\r\n\r\n")))
	fmt.Println(conn.Write([]byte("GET /hello HTTP/1.1\r\n\r\n")))
	data := make([]byte, 1000)
	fmt.Println(conn.Read(data))
	fmt.Println("")
	fmt.Println(string(data))

	time.Sleep(900 * time.Millisecond)

	fmt.Println(conn.Write([]byte("GET /hello1 HTTP/1.1\r\n\r\n")))
	memsetLoop(data, 0)
	fmt.Println(conn.Read(data))
	fmt.Println("")
	fmt.Println(string(data))

	time.Sleep(900 * time.Millisecond)

	fmt.Println(conn.Write([]byte("GET /hello2 HTTP/1.1\r\n\r\n")))
	memsetLoop(data, 0)
	fmt.Println(conn.Read(data))
	fmt.Println("")
	fmt.Println(string(data))

	return
	go http.ListenAndServe("0.0.0.0:8080", nil)
	// for i := 0; i < 10; i++ {
	// 	abc()
	// 	time.Sleep(time.Second)
	// }
	engine := gin.Default()
	engine.POST("/test", h_test)
	engine.Run("localhost:9876")
}

type User struct {
	Name    string `binding:"required" json:"name"`
	Age     int    `json:"age"`
	Schools []struct {
		Name  string `binding:"required" json:"name"`
		Level int    `binding:"required" json:"level"`
	} `binding:"gt=0,dive" json:"schools"`
}

func h_test(ctx *gin.Context) {
	fmt.Println("---------------------------")
	fmt.Println("RemoteAddr:" + ctx.Request.RemoteAddr)
	fmt.Println("Proto:" + ctx.Request.Proto)
	// user := &User{}
	// err := ctx.BindJSON(user)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(user)
	// }
	for k, v := range ctx.Request.Header {
		fmt.Printf("%s: %v\n", k, v)
	}
	ctx.JSON(200, map[string]string{"data": "ok"})
}
