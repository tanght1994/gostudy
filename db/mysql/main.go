package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// [user[:password]@][net[(addr)]]/dbname[?param1=value1&paramN=valueN]
	dsn := `root:123456@tcp(localhost:3306)/test`
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 设置连接池参数
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(500 * time.Second)
	db.SetConnMaxLifetime(500 * time.Second)

	// Ping一下试试
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 假设test表结构如下 id:int,name:string,age:int,hobby:string
	rows, err := db.Query("SELECT name, age FROM `test` WHERE `age`>=18 LIMIT 10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	name := ""
	age := 0
	for rows.Next() {
		rows.Scan(&name, &age)
		fmt.Println(name, age)
	}

	stmt1, err := db.Prepare("SELECT name, age FROM `test` WHERE `age`>=18 LIMIT 10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(stmt1)
	stmt2, err := db.Prepare("SELECT name, age FROM `test` WHERE `age`>=18 LIMIT 10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(stmt2)
	fmt.Println(123)

	time.Sleep(3 * time.Second)
	rows, err = stmt1.Query()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
