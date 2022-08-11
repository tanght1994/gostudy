package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := create_db(1)
	create_table(db)
	insert(db)
	query(db)
	query_with_many_statement(db)
	queryrow(db)
	interpolate(db)
}

// func tanght1() {
// 	// [user[:password]@][net[(addr)]]/dbname[?param1=value1&paramN=valueN]
// 	db, err := sql.Open("mysql", `tanght:Tht940415,./@tcp(www.tanght.xyz:3306)/test?interpolateParams=True`)
// 	if err != nil {
// 		fmt.Println("sql.Open error, ", err)
// 		os.Exit(1)
// 	}

// 	// 设置连接池参数
// 	db.SetMaxOpenConns(1)
// 	db.SetMaxIdleConns(1)
// 	db.SetConnMaxIdleTime(500 * time.Second)
// 	db.SetConnMaxLifetime(500 * time.Second)

// 	// Ping一下试试
// 	err = db.Ping()
// 	if err != nil {
// 		fmt.Println("db.Ping() error, ", err)
// 		os.Exit(1)
// 	}
// 	fmt.Println("ping ok")

// 	rows1, err := db.Query("SELECT b, c FROM `test` WHERE `c`>? LIMIT 10", "4 OR 1")
// 	if err != nil {
// 		fmt.Println("db.Query error, ", err)
// 		os.Exit(1)
// 	}
// 	b := ""
// 	c := 0
// 	// rows1.Next() 返回 false 时, 会顺便将此 Rows 所持有的 conn 放回到连接池中
// 	for rows1.Next() {
// 		rows1.Scan(&b, &c)
// 		fmt.Println(b, c)
// 	}

// 	rows1.Columns()

// 	// db.Query 会从 DB的连接池中找一个空闲的连接, 如果没有连接可用, 则阻塞等待
// 	// 所以 ROWS 必须保证 CLOSE, 不然将连接用完的话, 就该阻塞了
// 	rows2, err := db.Query("SELECT b, c FROM `test` WHERE `c`>? LIMIT 10", "4 OR 1")
// 	if err != nil {
// 		fmt.Println("db.Query error, ", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println(rows2)
// }

// func tanght2() {
// 	db := create_db(1)

// 	stmt1, err := db.Prepare("SELECT b, c FROM `test` WHERE `c`>=18 LIMIT 10")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	stmt2, err := db.Prepare("SELECT b, c FROM `test` WHERE `c`>=18 LIMIT 10")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	// rows1 不 close 则 stmt2.Query()会阻塞
// 	rows1, err := stmt1.Query()
// 	if err != nil {
// 		fmt.Println("stmt1.Query() error, ", err)
// 		os.Exit(1)
// 	}
// 	fmt.Println(rows1)

// 	rows2, err := stmt2.Query()
// 	if err != nil {
// 		fmt.Println("stmt2.Query() error, ", err)
// 		os.Exit(1)
// 	}
// 	fmt.Println(rows2)
// }

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func create_db(size int) *sql.DB {
	fmt.Println("create_db start")
	defer fmt.Println("create_db end")
	// 创建数据库连接池
	// 第二个参数格式: [user[:password]@][net[(addr)]]/dbname[?param1=value1&paramN=valueN]
	// interpolateParams=True 允许在sql中使用 ? 进行插值  sql驱动会处理插值
	// multiStatements=True 允许一次Exec中有多条sql语句, 多个语句用;分割
	// 最好不要用multiStatements, 因为这会增加SQL注入的风险
	db, err := sql.Open("mysql", `root:123456@tcp(www.tanght.xyz:3306)/haha?interpolateParams=True&multiStatements=True`)
	must(err)

	// 设置连接池参数
	db.SetMaxOpenConns(size)
	db.SetMaxIdleConns(size)
	db.SetConnMaxIdleTime(500 * time.Second)
	db.SetConnMaxLifetime(500 * time.Second)

	// Ping一下试试
	err = db.Ping()
	must(err)
	return db
}

func create_table(db *sql.DB) {
	fmt.Println("create_table start")
	defer fmt.Println("create_table end")
	if db == nil {
		db = create_db(1)
	}
	sql := `
	DROP TABLE IF EXISTS t1;
	CREATE TABLE t1 (id INT NOT NULL AUTO_INCREMENT, a VARCHAR(50), b INT, PRIMARY KEY (id));
	DROP TABLE IF EXISTS t2;
	CREATE TABLE t2 (id INT NOT NULL AUTO_INCREMENT, c VARCHAR(50), d INT, e INT, PRIMARY KEY (id));
	`
	_, err := db.Exec(sql)
	must(err)
}

func insert(db *sql.DB) {
	fmt.Println("insert start")
	defer fmt.Println("insert end")
	if db == nil {
		db = create_db(1)
	}
	var sql string
	var err error

	// 用 ? 进行插值, 完全是在sql驱动中处理, 不增加mysql服务的负担
	sql = "INSERT INTO t1 (a, b) VALUES (?, ?)"
	_, err = db.Exec(sql, "s1", 10)
	must(err)
	_, err = db.Exec(sql, "s2", 20)
	must(err)
	_, err = db.Exec(sql, "s3", 30)
	must(err)

	// 一次执行多条插入语句
	sql = `
	INSERT INTO t2 (c, d, e) VALUES ('s3', 40, 100);
	INSERT INTO t2 (c, d, e) VALUES ('s4', 50, 200);
	INSERT INTO t2 (c, d, e) VALUES ('s5', 60, 300);
	INSERT INTO t2 (c, d, e) VALUES ('中国', 1111, 2222);
	`
	_, err = db.Exec(sql)
	must(err)
}

// Query 查询多条数据, rows.Next 自动关闭 Rows
// 通过 Next + Scan 循环获取这些数据
func query(db *sql.DB) {
	fmt.Println("query start")
	defer fmt.Println("query end")
	if db == nil {
		db = create_db(1)
	}
	rows, err := db.Query("SELECT * FROM t1")
	must(err)
	var int1, int2 int
	var str1 string
	columns, err := rows.Columns()
	must(err)
	fmt.Println(columns)
	for rows.Next() {
		rows.Scan(&int1, &str1, &int2)
		fmt.Println(int1, str1, int2)
	}
}

// 一次 Query 执行多条SQL语句, 多个SQL用 ; 分割
// 使用rows.NextResultSet()进入到下一条SQL语句的结果集
func query_with_many_statement(db *sql.DB) {
	fmt.Println("query_with_many_statement start")
	defer fmt.Println("query_with_many_statement end")
	if db == nil {
		db = create_db(1)
	}
	sql := `
	SELECT * FROM t1;
	SELECT * FROM t2;
	`
	rows, err := db.Query(sql)
	must(err)
	var int1, int2, int3 int
	var str1 string
	fmt.Println("第一条SQL语句的结果")
	columns, err := rows.Columns()
	must(err)
	fmt.Println(columns)
	for rows.Next() {
		rows.Scan(&int1, &str1, &int2)
		fmt.Println(int1, str1, int2)
	}
	if !rows.NextResultSet() {
		return
	}
	fmt.Println("第二条SQL语句的结果")
	columns, err = rows.Columns()
	must(err)
	fmt.Println(columns)
	for rows.Next() {
		rows.Scan(&int1, &str1, &int2, &int3)
		fmt.Println(int1, str1, int2, int3)
	}
	if !rows.NextResultSet() {
		return
	}
	// 我们没有第三条语句, 所以代码运行不到这里
	fmt.Println("第三条SQL语句的结果")
	panic("hahahaha")
}

// QueryRow 查询1条数据 row.Scan 自动关闭 Row
// 即使你的SQL语句会查出多条数据, 那么也只能 Scan 到第1条数据, 其它数据会自动丢弃
// 如果你的SQL语句没查到数据, 那么返回 ErrNoRows
func queryrow(db *sql.DB) {
	fmt.Println("queryrow start")
	defer fmt.Println("queryrow end")
	if db == nil {
		db = create_db(1)
	}
	row := db.QueryRow("SELECT * FROM t1")
	var int1, int2 int
	var str1 string
	// 如果没有数据, 则返回ErrNoRows
	// 如果有1条数据, 则将这条数据 Scan 出来
	// 如果数据多余1条, 则只能 Scan 出来第一条, 剩下的数据会自动丢弃
	err := row.Scan(&int1, &str1, &int2)
	if err != nil {
		if err != sql.ErrNoRows {
			panic(err)
		}
	}
	fmt.Println(int1, str1, int2)
}

func interpolate(db *sql.DB) {
	// 使用?进行参数占位(插值)
	// sql驱动负责将?替换为真实的值, 然后生成真实的SQL语句发送给MySQL服务器
	// sql驱动根据你传递的参数的类型来决定是否用单引号包裹你的参数
	fmt.Println("interpolate start")
	defer fmt.Println("interpolate end")
	if db == nil {
		db = create_db(1)
	}
	// 生成的SQL语句为 INSERT INTO t1 (a, b) VALUES ('tanght', 100)
	_, err := db.Exec("INSERT INTO t1 (a, b) VALUES (?, ?)", "tanght", 100)
	must(err)

	// 生成的SQL语句为 INSERT INTO 't1' (a, b) VALUES ('wangqing', 500)
	_, err = db.Exec("INSERT INTO ? (a, b) VALUES (?, ?)", "t1", "wangqing", 500)
	fmt.Println(err)
}
