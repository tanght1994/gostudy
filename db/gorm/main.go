package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TangHongTao struct {
	Account        string `gorm:"column:account;primary_key;type:char(255)" json:"account"`
	Password       string `gorm:"column:password;type:varchar(255)" json:"password"`
	Ex             string `gorm:"column:ex;size:1024" json:"ex"`
	UserID         string `gorm:"column:user_id;type:varchar(255)" json:"userID"`
	AreaCode       string `gorm:"column:area_code;type:varchar(255)"`
	InvitationCode string `gorm:"column:invitation_code;type:varchar(255)"`
	RegisterIP     string `gorm:"column:register_ip;type:varchar(255)"`
	HahaHehe       int32  `gorm:"column:haha_hehe"`
}

func (TangHongTao) TableName() string {
	return "abc"
}

type WangQing struct {
	InvitationCode string    `gorm:"column:invitation_code;primary_key;type:varchar(32)"`
	CreateTime     time.Time `gorm:"column:create_time"`
	UserID         string    `gorm:"column:user_id"`
	LastTime       time.Time `gorm:"column:last_time"`
	Status         int32     `gorm:"column:status"`
}

func (WangQing) TableName() string {
	return "def"
}

func main() {
	dsn := "root:1234@tcp(192.168.7.45:3309)/test?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	must(err)
	db.AutoMigrate(
		&TangHongTao{},
		&WangQing{},
	)
	if !db.Migrator().HasTable(&TangHongTao{}) {
		db.Migrator().CreateTable(TangHongTao{})
	}
	if !db.Migrator().HasTable(&WangQing{}) {
		db.Migrator().CreateTable(WangQing{})
	}

	// db.Create(&TangHongTao{Account: "001", Password: "lalala"})
	var tht TangHongTao
	db.First(&tht, "001")
	fmt.Println(tht)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
