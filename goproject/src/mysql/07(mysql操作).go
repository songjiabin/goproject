package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	//db, err := sql.Open("mysql", "root:123456@/china?charset=utf8")
	// ⽤户名：密码@/数据库名称？编码⽅式
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/mydb")
	//db, err := sql.Open("mysql", "root:@/china?charset=utf8")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	//增删改查

	//增
	//insert(db)

	//删除
	//del(db)

	//改
	//update(db)

	//查
	selectdb()

}

//增
func insert(db *sql.DB) {
	result, error_insert := db.Exec("INSERT INTO person(name,age,rmb,gender,birthday) VALUES (?,?,?,?,?)", "宋佳宾", 23, 13343423, 1, 19920117)
	if error_insert != nil {
		fmt.Println(error_insert)
		return
	}

	lastInseriD, _ := result.LastInsertId() //返回最后受影响的id
	rows, _ := result.RowsAffected()        //多少行受影响
	fmt.Println(lastInseriD, rows)

}

//删
func del(db *sql.DB) {
	lastInseriDd, _ := db.Exec("delete from person WHERE name not LIKE ?;", "%宾")
	fmt.Println(lastInseriDd)
}

//改
func update(db *sql.DB) {
	result, error := db.Exec("UPDATE person set name =? WHERE name =?", "张设", "李永奇")
	if (error != nil) {
		fmt.Println(result)
		return
	}
	lastInseriD, _ := result.LastInsertId() //返回最后受影响的id
	rows, _ := result.RowsAffected()        //多少行受影响

	fmt.Println(lastInseriD, rows)
}

//查
func selectdb() {

	db, error := sqlx.Open("mysql", "root:123456@tcp(localhost:3306)/mydb")
	if error != nil {
		fmt.Println(nil)
		return
	}

	var p [] Person

	e := db.Select(&p, "select * from person;")
	if (e != nil) {
		fmt.Println(e)
		return
	}

	fmt.Println("数据为:", p)

}

//定义切片从而保存查询的结果

type Person struct {
	Id       int     `db:"id"`
	Name     string  `db:"name"`
	Age      int     `db:"age"`
	Money    float64 `db:"rmb"`
	Gender   bool    `db:"gender"`
	Birthday string  `db:"birthday"`
}
