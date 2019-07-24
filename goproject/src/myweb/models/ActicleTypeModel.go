package models

//文章的类型
// 一个类型可以有多个文章    一个文章只能对应一个类型
type ArticleType struct {
	Id       int        `orm:"pk;auto"`
	TypeName string     `orm:"size(20);unique"`
	Acticle  []*Acticle `orm:"reverse(many)"` //所对应的文章
}
