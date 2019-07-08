package main

import (
	"database/sql"
	"fmt"

)

func main() {
	db, err := sql.Open("mysql", "root:123456@/mytest?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err1 := db.Query("SELECT * FROM user_info")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(rows)
}
