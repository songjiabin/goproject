package main

import (
	_ "DDFruit/models"
	_ "DDFruit/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
