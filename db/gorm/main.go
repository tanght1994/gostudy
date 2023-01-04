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
	Query(db)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// ConnectDB 连接数据库
func ConnectDB() *gorm.DB {
	dsn := "root:123456@tcp(www.tanght.xyz:3306)/test?charset=utf8mb4&parseTime=true&loc=Local"
	cfg := gorm.Config{}
	// cfg.CreateBatchSize = 100          // 当insert一组数据的时候(1000个), 每次插入100个, 插入10次
	// cfg.PrepareStmt = true             // 每次遇到新的SQL语句, 就将他编译成预编译语句, 以后使用预编译语句
	// cfg.SkipDefaultTransaction = false // GORM的Update Create等操作是否需要在Transaction中执行, 默认是在Transaction中执行的. 在不在Transaction中执行, 对我们写代码来说没有任何影响
	// cfg.QueryFields = true             // Find(SELECT)的时候不用指定字段, 直接查询所有字段
	db, err := gorm.Open(mysql.Open(dsn), &cfg)
	must(err)
	return db
}

// AutoMigrate 自动创建或修改表
func AutoMigrate(db *gorm.DB) {
	// 删除表
	db.Migrator().DropTable(&tangHongTao{})

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
	// 插入一条数据 (使用struct)
	db.Create(&tangHongTao{B: "a"})

	// 插入多条数据 (使用struct list)
	db.Create([]tangHongTao{
		{B: "b"},
		{B: "c"},
	})

	// 插入一条数据 (使用map) 使用map插入数据时要指定表名 db.Model(tangHongTao{}) 或 db.Table("tht")
	db.Model(tangHongTao{}).Create(map[string]interface{}{"b": "d"})

	// 插入多条数据 (使用map list)
	db.Table("tht").Create([]map[string]interface{}{
		{"b": "e"},
		{"b": "f"},
	})
}

// Transaction 事务
// Transaction中的函数, 如果返回error, 则自动rollback, 如果没返回error, 则自动commit
func Transaction(db *gorm.DB) {
	// func中没有返回错误
	db.Transaction(func(tx *gorm.DB) error {
		tx.Create([]tangHongTao{
			{B: "g"},
			{B: "h"},
		})

		tx.Create([]tangHongTao{
			{B: "i"},
			{B: "j"},
		})
		// 没有return error, 所以会commit
		// g h i j 四行数据被成功插入到数据库
		return nil
	})

	// func中返回错误
	db.Transaction(func(tx *gorm.DB) error {
		tx.Create([]tangHongTao{
			{B: "k"},
			{B: "l"},
		})
		tx.Create([]tangHongTao{
			{B: "m"},
			{B: "n"},
		})
		// 返回了error, 所以会rollback
		// k l m n 四行数据不会被提交, 会被数据库rollback
		return fmt.Errorf("123")
	})
}

// Query 查询
// Find 	将结果提取到[]res中
// First 	将结果提取到res中 (且自动添加 ORDER BY pk LIMIT 1)
// Last 	将结果提取到res中 (且自动添加 ORDER BY pk DESC LIMIT 1)
// Take		将结果提取到res中 (不添加 ORDER BY 只添加 LIMIT 1)
// Count	将结果提取到int中 (且自动添加count(*))
func Query(db *gorm.DB) {
	var tht tangHongTao
	var thtList []tangHongTao

	// 1.First() (1.自动添加 ORDER BY 和 LIMIT 2.取值到struct)
	tht = tangHongTao{}
	// SELECT * FROM `tht` ORDER BY `tht`.`a` LIMIT 1
	// 即使我们没有使用Limit()和Order()函数来设置 ORDER BY 和 LIMIT, 但是First会自动添加
	db.Debug().Model(tangHongTao{}).First(&tht) // tht 一定要是zero值, 否则会影响SQL语句
	fmt.Println(tht)

	// 2.Last() (1.自动添加 ORDER BY 和 LIMIT 2.取值到struct)
	tht = tangHongTao{}
	// SELECT * FROM `tht` ORDER BY `tht`.`a` DESC LIMIT 1
	// 即使我们没有使用Limit()和Order()函数来设置 ORDER BY 和 LIMIT, 但是First会自动添加
	db.Debug().Model(tangHongTao{}).Last(&tht)
	fmt.Println(tht)

	// 3.Take() (1.自动添加 LIMIT 2.取值到struct)
	tht = tangHongTao{}
	// SELECT * FROM `tht` LIMIT 1
	// Take 会自动添加 LIMIT 1
	db.Debug().Model(tangHongTao{}).Take(&tht)
	fmt.Println(tht)

	// 4.Find() (1.取值到struct list 或 struct)
	thtList = []tangHongTao{}
	// SELECT * FROM `tht`
	// Find 不会添加任何东西
	// Find 接受一个 struct 数组, 将结果放到数组中
	// Find 也能接受一个单一的 struct, 那么Find就会把结果中的第一个值放到这个 struct 中, 然后close(rows)
	db.Debug().Model(tangHongTao{}).Find(&thtList)
	fmt.Println(thtList)

	// 5.Where() 上述所有函数前, 都可以添加 Where 进行条件过滤
	tht = tangHongTao{}
	// SELECT * FROM `tht` WHERE b > 'c' ORDER BY `tht`.`a` LIMIT 1
	db.Debug().Model(tangHongTao{}).Where("b > ?", "c").First(&tht)
	fmt.Println(tht)
}
