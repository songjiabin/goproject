package main

import (

	_ "myweb/models"
	_ "myweb/routers"

	"github.com/astaxie/beego"

)

func main() {
	beego.Run()

}
