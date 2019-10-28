package model

//和数据库中的表进行映射的实体类

type User struct {
	BaseModel
	Username string `gorm:"username" json:"username" binding:"required" validate="min=1,max=32"`
	Password string `gorm:"password" json:"password" binding:"required" validate="min=5,max=128"`
}

func CreateUser() {



}
