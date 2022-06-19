package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int, 5)
	go send2ch(ch)
	go recv4ch(ch)
	time.Sleep(100 * time.Hour)
}

func unblocksend(ch chan<- int) {
	num := rand.Int()%1000 + 1000
	select {
	case ch <- num:
		fmt.Println("send ", num)
	default:
		fmt.Println("!!!send ", num)
	}
}

func send2ch(ch chan<- int) {
	for {
		unblocksend(ch)
		time.Sleep(400 * time.Millisecond)
	}
}

func unblockrecv(ch <-chan int) {
	select {
	case x := <-ch:
		fmt.Println("                recv ", x)
	default:
		fmt.Println("                !!!recv")
	}
}

func recv4ch(ch <-chan int) {
	for {
		unblockrecv(ch)
		time.Sleep(1 * time.Second)
	}
}
