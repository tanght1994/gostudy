package main

import (
	"fmt"
	"sync"
)

func session1() {
	fmt.Println("session start")
	defer fmt.Println("session end")

	var thtabcd string
	db1 := create_db(1)
	db2 := create_db(1)
	defer db1.Close()
	defer db2.Close()

	// 两个连接分别设置@thtabcd变量
	// 验证 每条连接 是 一个会话
	db1.Exec("SET @thtabcd = 100")
	db2.Exec("SET @thtabcd = 200")

	// 连接1读取 @thtabcd
	row1 := db1.QueryRow("SELECT @thtabcd")
	must(row1.Err())
	row1.Scan(&thtabcd)
	fmt.Println(thtabcd) // 100

	// 连接2读取 @thtabcd
	row2 := db2.QueryRow("SELECT @thtabcd")
	must(row2.Err())
	row2.Scan(&thtabcd)
	fmt.Println(thtabcd) // 200

	// 所以每条连接是一个session
}

func session2() {
	fmt.Println("session2 start")
	defer fmt.Println("session2 end")
	db := create_db(3)

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		fmt.Println("1 start")
		defer fmt.Println("1 end")
		defer wg.Done()
		db.Exec("select sleep(5)")
	}()

	go func() {
		fmt.Println("2 start")
		defer fmt.Println("2 end")
		defer wg.Done()
		db.Exec("select sleep(5)")
	}()

	go func() {
		fmt.Println("3 start")
		defer fmt.Println("3 end")
		defer wg.Done()
		db.Exec("select sleep(5)")
	}()

	wg.Wait()
}
