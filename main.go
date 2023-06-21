package main

import (
	"QRcode/Code"

	"github.com/gin-gonic/gin"
)

func main() {
	// 开启Gin引擎
	r := gin.Default()
	// 使用该结构体下的函数,生成二维码图片
	app := Code.NewCodeController()
	r.LoadHTMLFiles("template/index.html")
	r.Static("/static", "./static")
	r.GET("/", app.GetCode)
	r.POST("/", app.GetCode)

	r.Run(":8080")
}
