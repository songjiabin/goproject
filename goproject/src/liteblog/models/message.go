package models

import "github.com/astaxie/beego/orm"

type Message struct {
	Id      int
	Content string `json:"content"`                 //评论的内容
	Key     string `orm:"unique" json:"key"`        //评论的key唯一标识
	User    *User  `orm:"rel(fk)" json:"user"`      //当前评论是哪个用户评论的
	Note    *Note  `orm:"rel(fk)" json:"note"`      //当前评论属于哪个文章
	Praise  int    `orm:"default(0)" json:"praise"` //点赞的数量
}

func SaveMessage(message Message) error {
	newOrm := orm.NewOrm()
	_, err := newOrm.Insert(&message)
	return err
}

func QueryMessageByKey(key string) (message Message, err error) {
	newOrm := orm.NewOrm()
	message = Message{
		Key: key,
	}
	err = newOrm.Read(&message, "Key")
	return
}

//得到所有的评论信息根据文章的key
func QueryMessage(key string) (messages []Message, err error) {
	newOrm := orm.NewOrm()
	messages = []Message{}
	_, err = newOrm.QueryTable("Message").RelatedSel("User").Filter("Note__Key", key).All(&messages)
	return

}




