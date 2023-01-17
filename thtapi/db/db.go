package db

import (
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
		modelUserPassword{},
		modelUserInfo{},
		modelGroupURL{},
		modelUserGroup{},
		modelURLWhiteList{},
		modelSvcName2SvcAddr{},
		modelURL2EndPoint{},
		modelTableModified{})
	must(err, "DB.AutoMigrate error, ")
	initTableData()
	common.LogCritical("数据库初始化成功")
	updateCache()
}

func initTableData() {
	mydb.Delete(modelTableModified{}, "1=1")
	mydb.Create(&modelTableModified{Name: modelURL2EndPoint{}.TableName(), Tag: 1})
	mydb.Create(&modelTableModified{Name: modelSvcName2SvcAddr{}.TableName(), Tag: 1})
	mydb.Create(&modelTableModified{Name: modelUserGroup{}.TableName(), Tag: 1})
	mydb.Create(&modelTableModified{Name: modelGroupURL{}.TableName(), Tag: 1})
	mydb.Create(&modelTableModified{Name: modelURLWhiteList{}.TableName(), Tag: 1})
}

func must(err error, msg string) {
	if err != nil {
		common.LogError(msg + err.Error())
		os.Exit(1)
	}
}
