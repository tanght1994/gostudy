package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

// Exit ...
func Exit(num int, format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func main() {
	engine := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	engine.POST("/tanght", func(ctx *gin.Context) {
		data, err := ioutil.ReadAll(ctx.Request.Body)
		fmt.Println(err, string(data))
		ctx.Request.Body.Close()
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		a := struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{}
		err = ctx.ShouldBindJSON(&a)
		fmt.Println(err, a)
	})
	fmt.Println(http.ListenAndServe(":13999", engine.Handler()))
}

func gintest() {

}

func fmctest() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("1.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		Exit(1, "firebase.NewApp error %s", err.Error())
	}
	client, err := app.Messaging(ctx)
	if err != nil {
		Exit(1, "app.Messaging error %s", err.Error())
	}
	fmt.Println("fcm client init ok")
	ttl := time.Second * 600
	message := messaging.MulticastMessage{
		Tokens: []string{
			"eUjuEaGrT628dyUQ4MRxLA:APA91bEtYVZ89eU4n5n19P1G8JZ3NaLFHH_favEBrB7WVMXTNRSPfId9HWE9Hb2cM7kr3J9rg0TS12WrrBtkM0tLRfupYLT7aPL1ayQ9qNCFoIlXQpXbui5YzQAnuLOKwjwAjLtlA_xM",
			"e-qIbDdQSNyWIl_1CNjZ2j:APA91bE6qjiPjxhERRhzky5byzqgtyaWzs1VRzKUdOzeZDmCInNNG-elMMHTgJiYyz49O6q_sv_B9P-NVyyBKYkokNVr5u4l51wc9CqvROGYMNa38pGzrdqVs7tNTYf7Y2uMoTlXEx-z",
		},
		Notification: &messaging.Notification{
			Title: "this is test title",
			Body:  "this is test body",
		},
		Android: &messaging.AndroidConfig{
			TTL: &ttl,
		},
	}
	if res, err := client.SendMulticast(ctx, &message); err != nil {
		fmt.Printf("FCMPush Run client.SendMulticast error %s\n", err.Error())
	} else {
		fmt.Printf("FCM Success:%d Failure:%d\n", res.SuccessCount, res.FailureCount)
		for idx, v := range res.Responses {
			if messaging.IsRegistrationTokenNotRegistered(v.Error) {
				fmt.Println(idx)
				// 将Tokens[idx]删掉, 这个token已经无效了
			}
			fmt.Println("error", messaging.IsRegistrationTokenNotRegistered(v.Error), v.Error)
		}
	}
}
