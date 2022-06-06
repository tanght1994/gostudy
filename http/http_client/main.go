package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

func main() {
	client := new(http.Client)

	wg := &sync.WaitGroup{}
	wg.Add(6)
	get(client, wg)
	time.Sleep(time.Duration(time.Second))
	get(client, wg)
	time.Sleep(time.Duration(time.Second))
	get(client, wg)
	time.Sleep(time.Duration(time.Second))
	// go get(client, wg)
	// time.Sleep(time.Duration(time.Second))
	// go get(client, wg)
	// time.Sleep(time.Duration(time.Second))
	// go get(client, wg)
	// time.Sleep(time.Duration(time.Second))
	// wg.Wait()
}

func get(client *http.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	req, _ := http.NewRequest("GET", "http://localhost:8000/tanght/", nil)
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
	resp, err := client.Post("http://localhost:8000/tanght/", "application/json", strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
	} else {
		b := make([]byte, 2048)
		n, _ := resp.Body.Read(b)
		fmt.Println(string(b[:n]))
	}
}
