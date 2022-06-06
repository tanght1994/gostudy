package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func main() {
	wc, res, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8000/ws", nil)
	must(err, "websocket.DefaultDialer.Dial error")
	fmt.Println(res.Header)
	go func() {
		for {
			stcmsg := STCMsg{}
			wc.ReadJSON(&stcmsg)
			fmt.Println("收到消息:", stcmsg)
		}
	}()

	for {
		target := 0
		msg := ""
		fmt.Println("请输入对方ID")
		fmt.Scanf("%d", &target)
		fmt.Println("请输入消息")
		fmt.Scanf("%s", &msg)
		wc.WriteJSON(NewCTSMsg(CTSMD_PTPChat{To: uint64(target), Msg: msg}))
	}
}

func must(e error, msg string) {
	if e != nil {
		fmt.Println(msg, e)
		panic(e)
	}
}
