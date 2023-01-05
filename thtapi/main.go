package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"thtapi/common"
	"thtapi/db"
	"time"
)

// Requester ...
var Requester http.Client

func main() {
	common.SetLogLevel(common.LevelDebug)
	common.LogCritical("thtapi start")
	db.A()
	transport := http.Transport{
		IdleConnTimeout: 10 * time.Second,
	}
	Requester = http.Client{
		Transport: &transport,
	}
}

// func aaaa() {
// 	listen, err := net.Listen("tcp", "0.0.0.0:8080")
// 	must(err)
// 	close(START)
// 	for {
// 		conn, err := listen.Accept()
// 		must(err)
// 		go dealconn(conn)
// 	}
// }

// func dealconn(conn net.Conn) {
// 	ddl := time.Now().Add(3 * time.Second)
// 	data := make([]byte, 0)
// 	for {
// 		conn.SetDeadline(ddl)
// 		tmp := make([]byte, 1024*4)
// 		_, err := conn.Read(tmp)
// 		if err != nil {
// 			conn.Close()
// 			break
// 		}
// 		data = append(data, tmp...)
// 	}
// 	fmt.Println("----------------------------")
// 	fmt.Println(string(data))
// 	fmt.Println("----------------------------")
// }

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ccccc() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var user, targetURL, originURL string
		var serverAddrs []string
		var err error
		originURL = r.URL.Path

		// url是否免验证
		if !db.InWhitelist(originURL) {
			// 验证登录
			user, err = db.ParseToken(r.Header.Get("token"))
			if err != nil {
				w.WriteHeader(400)
				return
			}

			// 验证权限
			if !db.HavePermission(user, originURL) {
				w.WriteHeader(400)
				return
			}
		}

		serverAddrs, targetURL, err = db.GetTargetURL(originURL)
		if err == db.ErrNotFindTargetURL {
			w.WriteHeader(404)
			return
		} else if err == db.ErrNotFindServerAddr {
			w.WriteHeader(500)
			return
		}

		serverAddr := serverAddrs[rand.Intn(len(serverAddrs))]
		url := "http://" + serverAddr + targetURL
		req, _ := http.NewRequest(r.Method, url, r.Body)
		req.URL.RawQuery = r.URL.RawQuery
		req.Header = r.Header.Clone()
		req.Header.Del("Connection")
		res, err := Requester.Do(req)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		defer res.Body.Close()
		header := w.Header()
		for k, v := range res.Header {
			for _, x := range v {
				header.Add(k, x)
			}
		}
		header.Del("Connection")
		data, _ := ioutil.ReadAll(res.Body)
		w.Write(data)
	})
	http.ListenAndServe("0.0.0.0:8000", nil)
}
