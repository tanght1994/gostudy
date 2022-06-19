package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	res, err := client.Do(req)
	fmt.Println(res)
}
