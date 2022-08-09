package main

import (
	"crypto/md5"
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

type LogData struct {
	Id      int
	Log     string
	Errsign string
}

func main() {
	dsn := "root:1234@tcp(192.168.7.45:3309)/hw_flog?interpolateParams=True&multiStatements=True"
	db, err := sqlx.Open("mysql", dsn)
	must(err)
	db.SetMaxOpenConns(1)
	db.SetConnMaxIdleTime(100 * time.Minute)
	db.SetConnMaxLifetime(100 * time.Hour)

	errsign := make(map[[16]byte]int)
	log := make(map[[16]byte]int)

	id := 634782
	data := make([]LogData, 0)
	for {
		data = data[0:0]
		db.Select(&data, "SELECT id, errsign, log FROM client_log WHERE id>? LIMIT 10000", id)
		if len(data) == 0 {
			break
		}
		id = data[len(data)-1].Id
		fmt.Println(id)
		for _, d := range data {
			errsign[md5.Sum([]byte(d.Errsign))]++
			log[md5.Sum([]byte(d.Log))]++
		}
	}
	fmt.Println("-------")
	fmt.Println(len(errsign))
	fmt.Println(len(log))
}


func getid(db *sqlx.DB, date string) int {
	id := 0
	err := db.Get(&id, "SELECT id FROM client_log WHERE gmt_create >= ? LIMIT 1", date)
	must(err)
	return id
}

type aggregationData struct {
	id int
	data string
}

func aggregation(db *sqlx.DB, field string, start int) {
	data := make([]LogData, 0)
	for {
		data = data[0:0]
		db.Select(&data, "SELECT id, errsign, log FROM client_log WHERE id>? LIMIT 10000", id)
		if len(data) == 0 {
			break
		}
		id = data[len(data)-1].Id
		fmt.Println(id)
		for _, d := range data {
			errsign[md5.Sum([]byte(d.Errsign))]++
			log[md5.Sum([]byte(d.Log))]++
		}
	}
	fmt.Println("-------")
	fmt.Println(len(errsign))
	fmt.Println(len(log))
}