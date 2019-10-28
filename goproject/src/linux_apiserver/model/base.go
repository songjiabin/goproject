package model

import "time"

type BaseModel struct {
	Id        int        `gorm:"column:primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt time.Time  `gorm:"column:createdAt" json:"createdtime"`
	UpdatedAt time.Time  `gorm:"column:updatedAt" json:"updatedtime"`
	DeletedAt *time.Time `gorm:"column:deletedAt" json:"deletedtime"`
}
