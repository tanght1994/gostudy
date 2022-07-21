package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ARG struct {
	T ARG2
	A string
	B string
}

type ARG2 struct {
	C string
	D string
}



func main() {
	dsn := "root:Tht940415,./@tcp(www.tanght.xyz:3306)/tanght"
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Println("sqlx.Open error, ", err)
		os.Exit(1)
	}
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(100 * time.Minute)
	db.SetConnMaxLifetime(100 * time.Hour)

	db.Get()
	db.Select()
	db.QueryRowx()
	db.Queryx()
	

	arg := ARG{ARG2{"9", "9"}, "8", "8"}

	// res, err := db.NamedExec(`INSERT INTO test (a, b) VALUES (:a, :b)`, map[string]interface{}{
	// 	"a": "6",
	// 	"b": "6",
	// })

	res, err := db.NamedExec(`INSERT INTO test (a, b) VALUES (:a, :b)`, arg)
	if err != nil {
		fmt.Println("db.NamedExec error, ", err)
		os.Exit(1)
	}
	last_id, _ := res.LastInsertId()
	affected, _ := res.RowsAffected()
	fmt.Println(last_id, affected)
}
