package model

import (
	"myapiserver/util"
	"sync"
)

type BaseModel struct {
	Id        uint       `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"` //主键
	CreatedAt util.LocalTime  `gorm:"column:createdAt" json:"-"`
	UpdatedAt util.LocalTime  `gorm:"column:updatedAt" json:"-"`
	DeletedAt *util.LocalTime `gorm:"column:deletedAt" sql:"index" json:"-"`
}

type UserInfo struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	SayHello  string `json:"sayHello"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint]*UserInfo
}
