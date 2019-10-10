package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

//https://www.yoytang.com/go-gin-doc.html#Gin%20的介绍

func main() {
	engine := gin.Default()
	engine.POST("/uploads", func(context *gin.Context) {
		//获取请求中文件的名字为 file
		file, header, err := context.Request.FormFile("file")
		name := context.DefaultPostForm("name", "基本密码")
		log.Print(name)
		if err != nil {
			context.String(http.StatusBadRequest, "Bad request")
			return
		}
		filename := header.Filename

		fmt.Println(file, err, filename)
		// 将传递的文件保存到本地
		out, err := os.Create("C:/Users/songjiabin1/Desktop/" + filename)
		defer out.Close()
		io.Copy(out, file)
		context.String(http.StatusOK, "upload successful")

	})
	engine.Run()
}
