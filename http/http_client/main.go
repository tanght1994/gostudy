package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

func main() {
	client := new(http.Client)
	client.Transport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxConnsPerHost:       30,
	}

	wg := &sync.WaitGroup{}
	wg.Add(10)
	// get(client, wg)
	// time.Sleep(time.Duration(time.Second))
	// get(client, wg)
	// time.Sleep(time.Duration(time.Second))
	// get(client, wg)
	// time.Sleep(time.Duration(time.Second))
	post(client, wg)
	// time.Sleep(time.Duration(time.Second))
	post(client, wg)
	// time.Sleep(time.Duration(time.Second))
	post(client, wg)
	post(client, wg)
	post(client, wg)
	post(client, wg)
	post(client, wg)
	post(client, wg)
	post(client, wg)
	post(client, wg)
	// time.Sleep(time.Duration(time.Second))
	wg.Wait()
}

func get(client *http.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	req, _ := http.NewRequest("GET", "http://localhost:9876/test", nil)
	// req.Header.Set("Connection", "keep-alive")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	} else {
		b := make([]byte, 2048)
		n, _ := resp.Body.Read(b)
		fmt.Println(string(b[:n]))
	}
}

func post(client *http.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	s := `{"Name": "tanght", "Age": 18}`
	resp, err := client.Post("http://localhost:9876/test", "application/json", strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
	} else {
		// b := make([]byte, 2048)
		// n, _ := resp.Body.Read(b)
		// fmt.Println(string(b[:n]))
		resp.Body.Close()
		fmt.Println(resp.Status)
	}
}
