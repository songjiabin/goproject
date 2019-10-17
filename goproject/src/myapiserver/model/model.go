package model

import (
	"time"
)

type BaseModel struct {
	Id        uint     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"` //主键
	CreatedAt time.Time  `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updatedAt" json:"-"`
	DeletedAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
}
