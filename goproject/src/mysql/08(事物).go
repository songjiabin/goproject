package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//批量完成sql的时候 要是有一个sql有问题 那么据全部不能完成
func main() {

	//连接数据库
	db, error := sqlx.Open("mysql", "root:123456@tcp(localhost:3306)/mydb")
	handError(error)

	//开始事物
	tx, error_begin := db.Begin()
	handError(error_begin)

	//进行增删改查
	res1, e1 := tx.Exec("INSERT INTO  person(name,age,birthday,gender,rmb) VALUES (?,?,?,?,?)", "赵立康", "34", "19960426", true, 233454)

	res2, e2 := tx.Exec("DELETE FROM person WHERE name LIKE ?", "%设")

	res3, e3 := tx.Exec("UPDATE   person SET  name =? WHERE id=?;", "许美琳", 1)

	if e1 != nil || e2 != nil || e3 != nil {
		//任意的一个语句出现错误
		fmt.Println(e1, e2, e3)
		//事物回滚 相当于啥都没做
		tx.Rollback()
	} else {

		//否则提交事物
		tx.Commit()

		ra1, _ := res1.RowsAffected()
		ra2, _ := res2.RowsAffected()
		ra3, _ := res3.RowsAffected()
		fmt.Println("事物提交成功，受影响的行：", ra1, ra2, ra3)
	}

}

func handError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
