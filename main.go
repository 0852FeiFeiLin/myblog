package main

import (
	"myBlog/utils"
	_ "myBlog/routers"
	"github.com/astaxie/beego"
)

func main2() {
	//连接数据库
	utils.ConnectDB()

	//设置静态资源
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}

