package db

import (
	"fmt"
	"os"
	"sync"
	"thtapi/common"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var mydb *gorm.DB

func init() {
	synclock = sync.RWMutex{}
	var err error
	mydb, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	must(err, "gorm.Open error, ")
	mydb.Set("gorm:table_options", "CHARSET=utf8")
	mydb.Set("gorm:table_options", "collation=utf8_unicode_ci")
	err = mydb.AutoMigrate(
		&UserPassword{},
		&UserInfo{},
		&GroupURL{},
		&UserGroup{},
		&URLWhiteList{},
		&ServerAddr{},
		&URL2URL{})
	must(err, "DB.AutoMigrate error, ")
	// if !mydb.Migrator().HasTable(&UserPassword{}) {
	// 	err = mydb.Migrator().CreateTable(UserPassword{})
	// 	must(err, "创建表失败,")
	// }
	// if !mydb.Migrator().HasTable(&UserInfo{}) {
	// 	err = mydb.Migrator().CreateTable(UserInfo{})
	// 	must(err, "创建表失败,")
	// }
	// if !mydb.Migrator().HasTable(&GroupURL{}) {
	// 	err = mydb.Migrator().CreateTable(GroupURL{})
	// 	must(err, "创建表失败,")
	// }
	// if !mydb.Migrator().HasTable(&UserGroup{}) {
	// 	err = mydb.Migrator().CreateTable(UserGroup{})
	// 	must(err, "创建表失败,")
	// }
	// if !mydb.Migrator().HasTable(&URLWhiteList{}) {
	// 	err = mydb.Migrator().CreateTable(URLWhiteList{})
	// 	must(err, "创建表失败,")
	// }
	// if !mydb.Migrator().HasTable(&ServerAddr{}) {
	// 	err = mydb.Migrator().CreateTable(ServerAddr{})
	// 	must(err, "创建表失败,")
	// }
	// if !mydb.Migrator().HasTable(&URL2URL{}) {
	// 	err = mydb.Migrator().CreateTable(URL2URL{})
	// 	must(err, "创建表失败,")
	// }
	common.LogCritical("数据库初始化成功")

	syncdb()
}

func must(err error, msg string) {
	if err != nil {
		common.LogError(msg + err.Error())
		os.Exit(1)
	}
}

// A ...
func A() {
	fmt.Println("A")
}
