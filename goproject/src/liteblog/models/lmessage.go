package models

import "github.com/astaxie/beego/orm"

//留言板
type LMessage struct {
	Id      int
	Content string `json:"content"`                 //评论的内容
	Key     string `orm:"unique" json:"key"`        //评论的key唯一标识
	User    *User  `orm:"rel(fk)" json:"user"`      //当前评论是哪个用户评论的
	Praise  int    `orm:"default(0)" json:"praise"` //点赞的数量
}

func SaveLMessage(message LMessage) error {
	newOrm := orm.NewOrm()
	_, err := newOrm.Insert(&message)
	return err
}

func QueryLMessageByKey(key string) (message LMessage, err error) {
	newOrm := orm.NewOrm()
	message = LMessage{
		Key: key,
	}
	err = newOrm.Read(&message, "Key")
	return
}

//得到所有的评论信息
func QueryLMessage() (messages []LMessage, err error) {
	newOrm := orm.NewOrm()
	messages = []LMessage{}
	_, err = newOrm.QueryTable("LMessage").RelatedSel("User").OrderBy("-Id").All(&messages)
	return

}
