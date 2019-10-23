package main

import (
	"github.com/gin-gonic/gin"
	"linux_apiserver/router"
)

func main() {
	//new engine
	engine := gin.New()
	router.Load(engine)
	engine.Run()
}
