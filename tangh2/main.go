package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var count = 1000

func tanght(wg *sync.WaitGroup, name string) {
	defer wg.Done()
	url := ""
	url = "http://gtf.ai.xingzheai.cn/v2.0/game_chat_ban/detect_text"        // 国内
	url = "http://gtf.ai-abroad.xingzheai.cn/v2.0/game_chat_ban/detect_text" // 海外
	payload := map[string]interface{}{
		"token":        "VESGY9UAP5LMDW7H",
		"context":      "大家好，我是唐洪涛，我今年18岁。",
		"user_id":      "123456789",
		"context_type": "chat",
		"data_id":      "1",
	}
	b, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(name, "json.Marshal err, byebye!!!!!!!!!", err)
		return
	}
	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(b))
	var data map[string]interface{}
	if resp != nil {
		respBytes, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(respBytes, &data)
		if err != nil {
			fmt.Println(name, "json.Unmarshal err, ", err, " ---- ", string(respBytes))
		}
	}

	fmt.Println(name, " start")

	t0 := time.Now().UnixNano()
	for i := 0; i < count; i++ {
		resp, err = http.Post(url, "application/json", bytes.NewBuffer(b))
		if err != nil {
			fmt.Println(name, "Post err, ", err)
			continue
		}
		if resp != nil {
			respBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(name, "ioutil.ReadAll err, ")
				continue
			}

			err = json.Unmarshal(respBytes, &data)
			if err != nil {
				// fmt.Println(name, "json.Unmarshal err, ", err, " ---- ", string(respBytes))
				fmt.Print("1")
			}
			fmt.Print("0")
		}
	}
	t1 := time.Now().UnixNano()
	fmt.Println(name, " ", (t1-t0)/1000/1000/int64(count), "ms")
}

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go tanght(wg, strconv.Itoa(i))
	}
	wg.Wait()
}
