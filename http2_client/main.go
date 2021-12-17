package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/http2"
)

func main() {
	client := new(http.Client)
	ca, err := ioutil.ReadFile("cacert.pem")
	if err != nil {
		fmt.Println(err)
		return
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(ca)
	tlsConf := &tls.Config{RootCAs: pool}
	client.Transport = &http2.Transport{TLSClientConfig: tlsConf}

	wg := &sync.WaitGroup{}
	wg.Add(6)
	go get(client, wg)
	time.Sleep(time.Duration(time.Second))
	go get(client, wg)
	time.Sleep(time.Duration(time.Second))
	go get(client, wg)
	time.Sleep(time.Duration(time.Second))
	go post(client, wg)
	time.Sleep(time.Duration(time.Second))
	go post(client, wg)
	time.Sleep(time.Duration(time.Second))
	go post(client, wg)
	time.Sleep(time.Duration(time.Second))
	wg.Wait()
}

func get(client *http.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := client.Get("https://localhost:8000/")
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
	resp, err := client.Post("https://localhost:8000/", "application/json", strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
	} else {
		b := make([]byte, 2048)
		n, _ := resp.Body.Read(b)
		fmt.Println(string(b[:n]))
	}
}
