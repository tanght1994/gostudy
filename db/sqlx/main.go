package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db := create_db(1)
	defer db.Close()
	create_table(db)
	Select(db)
	Get(db)
	Named(db)
	SliceScan(db)
	MapScan(db)
}

func create_db(size int) *sqlx.DB {
	dsn := "root:123456@tcp(www.tanght.xyz:3306)/haha?interpolateParams=True&multiStatements=True"
	db, err := sqlx.Open("mysql", dsn)
	must(err)
	db.SetMaxOpenConns(size)
	db.SetConnMaxIdleTime(100 * time.Minute)
	db.SetConnMaxLifetime(100 * time.Hour)
	return db
}

func create_table(db *sqlx.DB) {
	fmt.Println("create_table start")
	defer fmt.Println("create_table end")
	sql := `
	DROP TABLE IF EXISTS t1;
	CREATE TABLE t1 (id INT NOT NULL AUTO_INCREMENT, a VARCHAR(50), b INT, PRIMARY KEY (id));
	INSERT INTO t1 (a, b) VALUES ('a', 1), ('b', 2), ('c', 3), ('d', 4);
	`
	_, err := db.Exec(sql)
	must(err)
}

func Select(db *sqlx.DB) {
	fmt.Println("Select start")
	defer fmt.Println("Select end")

	var data []struct {
		Id int
		A  string
		B  int
	}
	var err error

	// select 的第一个参数必须是*slice
	// slice 中的元素是 struct, struct中的字段与 SQL 中的字段相对应(不区分大小写)
	// select 函数会把SQL的结果放在slice中
	err = db.Select(&data, "SELECT id, a, b FROM t1")
	must(err)
	fmt.Println(data)

	// 查不到结果不会报错
	// 查到的结果为空, 所以data会被设置为空slice
	err = db.Select(&data, "SELECT id, a, b FROM t1 WHERE id > 99999")
	must(err)
	fmt.Println(data)

	// 如果 select 只查询1个字段, 那么slice中的元素不必是struct, 可以只是一个基础类型
	var dddd []int
	err = db.Select(&dddd, "SELECT id FROM t1")
	must(err)
	fmt.Println(dddd)
}

func Get(db *sqlx.DB) {
	fmt.Println("Get start")
	defer fmt.Println("Get end")

	var data struct {
		Id int
		A  string
		B  int
	}
	var err error

	// Get的第一个参数是struct, SQL的结果将保存到data处
	err = db.Get(&data, "SELECT * FROM t1 WHERE id=1")
	must(err)
	fmt.Println(data)

	err = db.Get(&data, "SELECT * FROM t1 WHERE id=2")
	must(err)
	fmt.Println(data)

	// 如果SQL只查询1个字段, 那么Get的第一个参数可以不是struct
	var id int
	err = db.Get(&id, "SELECT id FROM t1 WHERE id=1")
	must(err)
	fmt.Println(id)

	// 如果查不到结果, Get返回sql.ErrNoRows
	err = db.Get(&data, "SELECT * FROM t1 WHERE id=99999")
	if err == sql.ErrNoRows {
		fmt.Println(err)
	} else {
		panic("aaaaaaa")
	}

	// 如果SQL语句的结果为多条, Get只获取第一条, 自动丢弃剩余数据
	err = db.Get(&data, "SELECT * FROM t1")
	must(err)
	fmt.Println(data)
}

func Named(db *sqlx.DB) {
	fmt.Println("Named start")
	defer fmt.Println("Named end")

	var err error
	var res sql.Result

	// 批量插入, 使用struct数组
	data1 := []struct {
		A string
		B int
	}{
		{"python", 1},
		{"C++", 2},
		{"JAVA", 3},
		{"Python", 4},
	}
	res, err = db.NamedExec("INSERT INTO t1 (a, b) VALUES (:a, :b)", data1)
	must(err)
	fmt.Println(res.RowsAffected())

	// 批量插入, 使用map[string]interface{}数组
	data2 := []map[string]interface{}{
		{"a": "python", "b": 4},
		{"a": "C++", "b": 4},
		{"a": "JAVA", "b": 2},
		{"a": "Python", "b": 1},
		{"a": "中国", "b": 1},
	}
	res, err = db.NamedExec("INSERT INTO t1 (a, b) VALUES (:a, :b)", data2)
	must(err)
	fmt.Println(res.RowsAffected())

	// 可以批量插入, 当然也可以单个插入啊
	data3 := map[string]interface{}{"a": "Golang", "b": 1111}
	res, err = db.NamedExec("INSERT INTO t1 (a, b) VALUES (:a, :b)", data3)
	must(err)
	fmt.Println(res.RowsAffected())

	data4 := struct {
		A string
		B int
	}{"GOGOGO", 1}
	res, err = db.NamedExec("INSERT INTO t1 (a, b) VALUES (:a, :b)", data4)
	must(err)
	fmt.Println(res.RowsAffected())
}

func IN(db *sqlx.DB) {}

func SliceScan(db *sqlx.DB) {
	a, err := db.Queryx("SELECT * FROM t1")
	must(err)
	for a.Next() {
		data, err := a.SliceScan()
		must(err)
		fmt.Println(data)
	}
}

func MapScan(db *sqlx.DB) {
	a, err := db.Queryx("SELECT * FROM t1")
	must(err)
	data := map[string]interface{}{}
	for a.Next() {
		a.MapScan(data)
		fmt.Println(string(data["a"].([]uint8)), string(data["b"].([]uint8)), string(data["id"].([]uint8)))
	}
}
