package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//db, err := sql.Open("mysql", "root:123@/mydb?charset=utf8")
	// ⽤户名：密码@/数据库名称？编码⽅式
	//db, err :=sql.Open("mysql","root:123456:@tcp(localhost:3306)/china")
	db, err := sql.Open("mysql", "root:@/china?charset=utf8")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err1 := db.Query("SELECT * FROM t_city")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(rows)
}
