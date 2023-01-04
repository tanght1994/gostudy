package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type tangHongTao struct {
	A uint64 `gorm:"column:a;type:bigint unsigned;primary_key;auto_increment;not null" json:"a"`
	B string `gorm:"column:b;type:varchar(255)" json:"b"`
}

func (tangHongTao) TableName() string {
	// 指定表名, 如果不指定的话, 默认是 tang_hong_taos
	return "tht"
}

func main() {
	db := ConnectDB()
	// db.Session(&gorm.Session{}) // 返回db, 只是叫session而已, 其实就是db, session是db的拷贝, 为什么要这样? 因为这样的话我们可以重新设置db的gorm.Config, 且不影响原始的db对象
	// db.Table("your_table_name") // 返回db, 顺便将 db.Statement.Table 设置为你想操作的表名
	// db.Model(tangHongTao{})     // 返回db, 顺便将 db.Statement.Table 设置为你想操作的表名, Model()会从结构体中提取表名
	// db.Begin(&sql.TxOptions{})  // 返回db, 只是叫tx而已, 顺便将 db.Statement.ConnPool 设置为事务的连接, 然后使用这个tx进行增删改查, 完事之后记得commit
	AutoMigrate(db)
	Insert(db)
	Transaction(db)
	var tht tangHongTao
	db.First(&tht)
	fmt.Println(tht)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// ConnectDB 连接数据库
func ConnectDB() *gorm.DB {
	dsn := "root:1234@tcp(www.tanght.xyz:3306)/test?charset=utf8mb4&parseTime=true&loc=Local"
	cfg := gorm.Config{}
	// cfg.CreateBatchSize = 100          // 当insert一组数据的时候(1000个), 每次插入100个, 插入10次
	// cfg.PrepareStmt = true             // 每次遇到新的SQL语句, 就将他编译成预编译语句, 以后使用预编译语句
	// cfg.SkipDefaultTransaction = false // 是否每次 create update delete 都需要手动commit
	// cfg.QueryFields = true             // Find(SELECT)的时候不用指定字段, 直接查询所有字段
	db, err := gorm.Open(mysql.Open(dsn), &cfg)
	must(err)
	return db
}

// AutoMigrate 自动创建或修改表
func AutoMigrate(db *gorm.DB) {
	// 自动创建或修改表
	// 如果数据库中无此表, 则创建
	// 如果数据库中有此表, 则对比Go代码中表的定义与数据库中的表结构是否一样, 如果不一样则修改数据库
	// 修改原则:
	//    Go中新增的字段, 则数据库也要相应的增加字段
	//    Go中删除的字段, 数据库中的字段不删除!!!
	//    Go中字段类型发生变化, 数据库中的类型也要跟着变化
	//    比如说将 varchar 修改为 int  那么对于已经存在的记录aaaaa, 就会报错
	//    *** 主键字段不能修改类型 ***
	err := db.AutoMigrate(&tangHongTao{})
	must(err)
}

// Insert 插入数据
func Insert(db *gorm.DB) {
	// 插入一条数据
	db.Create(&tangHongTao{B: "10"}).Commit()

	// 插入多条数据
	db.Table("tht").Create([]map[string]interface{}{
		{"b": "b1"},
		{"b": "b2"},
		{"b": "b3"},
		{"b": "b4"},
		{"b": "b5"},
	}).Commit()
}

// Transaction 事务
func Transaction(db *gorm.DB) {
	// Transaction 中的函数, 如果返回error, 则自动rollback
	// 如果没返回error, 则自动commit
	db.Transaction(func(tx *gorm.DB) error {
		val := []tangHongTao{
			{B: "1"},
			{B: "2"},
			{B: "3"},
			{B: "4"},
			{B: "5"},
		}
		tx.Create(val)
		// tx.Table("tht").Create([]map[string]interface{}{
		// 	{"b": 11},
		// 	{"b": 12},
		// 	{"b": 13},
		// 	{"b": 14},
		// })
		return nil
	})
}
