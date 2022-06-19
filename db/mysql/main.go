package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	tanght2()
}

func tanght1() {
	// [user[:password]@][net[(addr)]]/dbname[?param1=value1&paramN=valueN]
	db, err := sql.Open("mysql", `tanght:Tht940415,./@tcp(www.tanght.xyz:3306)/test?interpolateParams=True`)
	if err != nil {
		fmt.Println("sql.Open error, ", err)
		os.Exit(1)
	}

	// 设置连接池参数
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxIdleTime(500 * time.Second)
	db.SetConnMaxLifetime(500 * time.Second)

	// Ping一下试试
	err = db.Ping()
	if err != nil {
		fmt.Println("db.Ping() error, ", err)
		os.Exit(1)
	}
	fmt.Println("ping ok")

	rows1, err := db.Query("SELECT b, c FROM `test` WHERE `c`>? LIMIT 10", "4 OR 1")
	if err != nil {
		fmt.Println("db.Query error, ", err)
		os.Exit(1)
	}
	b := ""
	c := 0
	// rows1.Next() 返回 false 时, 会顺便将此 Rows 所持有的 conn 放回到连接池中
	for rows1.Next() {
		rows1.Scan(&b, &c)
		fmt.Println(b, c)
	}

	// db.Query 会从 DB的连接池中找一个空闲的连接, 如果没有连接可用, 则阻塞等待
	// 所以 ROWS 必须保证 CLOSE, 不然将连接用完的话, 就该阻塞了
	rows2, err := db.Query("SELECT b, c FROM `test` WHERE `c`>? LIMIT 10", "4 OR 1")
	if err != nil {
		fmt.Println("db.Query error, ", err)
		os.Exit(1)
	}

	fmt.Println(rows2)
}

func tanght2() {
	// [user[:password]@][net[(addr)]]/dbname[?param1=value1&paramN=valueN]
	db, err := sql.Open("mysql", `tanght:Tht940415,./@tcp(www.tanght.xyz:3306)/test?interpolateParams=True`)
	if err != nil {
		fmt.Println("sql.Open error, ", err)
		os.Exit(1)
	}

	// 设置连接池参数
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxIdleTime(500 * time.Second)
	db.SetConnMaxLifetime(500 * time.Second)

	// Ping一下试试
	err = db.Ping()
	if err != nil {
		fmt.Println("db.Ping() error, ", err)
		os.Exit(1)
	}
	fmt.Println("ping ok")

	stmt1, err := db.Prepare("SELECT b, c FROM `test` WHERE `c`>=18 LIMIT 10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stmt2, err := db.Prepare("SELECT b, c FROM `test` WHERE `c`>=18 LIMIT 10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// rows1 不 close 则 stmt2.Query()会阻塞
	rows1, err := stmt1.Query()
	if err != nil {
		fmt.Println("stmt1.Query() error, ", err)
		os.Exit(1)
	}
	fmt.Println(rows1)

	rows2, err := stmt2.Query()
	if err != nil {
		fmt.Println("stmt2.Query() error, ", err)
		os.Exit(1)
	}
	fmt.Println(rows2)
}
