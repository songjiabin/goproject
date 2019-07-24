package models

//来试试一对一的操作

//一个人一个外貌
//一个外貌对应一个人
// 所以是 一对一

//外貌
type PersonName struct {
	Id     int
	Title  string
	Person *Person `orm:"reverse(one)"`
}

//用户
type Person struct {
	Id         int
	Name       string
	PersonName *PersonName `orm:"rel(one)"`
}

//手动设置表的名字
func (u *PersonName) TableName() string {
	return "name"
}
