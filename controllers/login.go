package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myBlog/entity"
	"myBlog/models"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/27 12:42
 **/

type IndexController struct {
	beego.Controller
}
//跳转到登录页面
func (i *IndexController)ToIndex(){
	i.TplName = "login.html"
}
/*
	登录逻辑：解析用户表单，数据传输至数据库进行查询，存在登陆成功，并且设置id为sessionId，.SetSession("userId",Id)
			不存在error,挑战到错误页面
*/
func(i *IndexController) Index(){
	//登录成功跳转到tags页面
	//1、解析前端表单
	var user entity.Users
	err := i.ParseForm(&user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//输出前端数据
	fmt.Println("前端username:",user.Username)
	fmt.Println("前端password:",user.Password)
	fmt.Println("传递：",user)
	//传入数据库查询方法验证
	users := models.Login(user)
	fmt.Println("数据库查询到的数据：",users)

	//通过判断id值来判断数据库是否查询到值
	if users.Id <= 0{
		i.Data["errInfo"] = "您输入的用户名和密码输入错误，请重新输入"
		i.TplName = "login.html"
		return
	}
	fmt.Println("登录成功！")
	i.SetSession("userName",user.Username)
	fmt.Println("Session:",i.GetSession("userName"))
	//登录成功跳转到tags页面
	i.TplName = "tags.html"

/*	//请求重定向
	i.Redirect("/tags.html",302)*/
	i.Redirect("blog.html",302)
}