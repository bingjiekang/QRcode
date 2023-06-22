package Code

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	// "log"
)

type CodeController struct{}

func NewCodeController() CodeController {
	return CodeController{}
}

// 生成一个二维码
func (g *CodeController) GetCode(c *gin.Context) {
	// 提交待转化信息
	if c.Request.Method == "POST" {
		// 保存path的名字
		name := strconv.FormatInt(time.Now().UnixNano(), 10)

		// 图片保存的路径
		pngpath := "./static/codeimage/" + name + ".png"
		// 静态文件访问目录
		path := "." + pngpath
		// 根据获取的内容 生成二维码
		var requestMap = make(map[string]string)
		json.NewDecoder(c.Request.Body).Decode(&requestMap)
		fmt.Println("这是传过来的json数据需要:", requestMap["data"])
		// sult := "https://iris.kangbingjie.com"
		sult := requestMap["data"]
		qrCode, err := qrcode.New(sult, qrcode.Highest)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("ok")
		}

		// 保存成文件
		err = qrCode.WriteFile(256, pngpath)
		if err != nil {
			fmt.Println(err)
		}
		// 返回对应的页面
		c.HTML(http.StatusOK, "index.html", gin.H{
			"context": requestMap["data"],
			"path":    path,
		})

	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"context": "",
			"path":    "",
		})
	}
	return
}
