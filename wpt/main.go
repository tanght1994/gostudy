package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// wpt 是 web performance test 的缩写

func main() {
	req, err := http.NewRequest("GET", "http://www.tanght.xyz:10004/", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Close = true
	// req.Header.Add("Connection", "close")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	res.Body.Close()
	fmt.Println(string(data))

	req, err = http.NewRequest("GET", "http://www.tanght.xyz:10004/", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Close = true
	// req.Header.Add("Connection", "close")
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	res.Body.Close()
	fmt.Println(string(data))

	req, err = http.NewRequest("GET", "http://www.tanght.xyz:10004/", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Close = true
	// req.Header.Add("Connection", "close")
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	res.Body.Close()
	fmt.Println(string(data))
}

type ReqData struct {
	method string
	url    string
}
