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
 * @DateTime: 2021/11/27 17:26
 **/
type RegisterController struct {
	beego.Controller
}

/*
	跳转到注册页面
*/
func (r *RegisterController) ToRegister() {
	r.TplName = "register.html"
}

/*
	注册用户逻辑，解析用户form表单中的数据，然后先判断用户名是否已经存在，不存在让他插入成功，
	注册成功跳转会login登录页面。
*/
func (r *RegisterController) Register() {
	//1、接收客户端提交的请求参数
	/*userName := r.GetString("username")
	passWord := r.GetString("password")
	rePwd := r.GetString("password2")
	//查看前端传递的值
	fmt.Println(userName)
	fmt.Println(passWord)
	fmt.Println(rePwd)
*/

	var user entity.Users
	r.ParseForm(&user)
	rePwd := r.GetString("password2")

	fmt.Println("前端user:",user)
	fmt.Println("前端username:",user.Username)
	fmt.Println("前端password:",user.Password)
	fmt.Println("前端rePassword:",rePwd)


	//判断两次输入的密码一不一样
	if user.Password != rePwd {
		r.Data["errInfo"] = "输入密码两次不一致"
		return
	}

	//2、检查所提交的请求参数进行合法性检查(是否为空，长度匹配)
	/*
		用前端技术来实现，
	*/

	//3、调用modules的查询方法，查看用户是佛已经存在
	isExist := models.QueryByUserName(user)
	if isExist == true  { //存在
		r.Data["errInfo"] = "该用户名已经存在，请重新输入"
		return
	}
	/*//判断用户id是否不为0 。不为0代表用户存在，就不能插入
	if quser.Id > 0 { //注册的用户已经存在
		r.Data["errInfo"] = "用户名已经存在，请重新输入"
		return
	}*/

	//插入方法把user插入
	_, _, err := models.Register(user) //根据具体需要判断是否要id和受影响的行数
	if err != nil {
		fmt.Println(err.Error())
		r.Data["errInfo"] = "注册发生错误，请重新注册"
		//这里也可以设置样式渲染，就是加上标签
		return
	}
	//没出错，就是注册成功，跳转到login  这个方法只能跳转静态页面
	//r.TplName = "login.html"

	//请求重定向  这样就能根据指定的url匹配到响应的控制文件
	//跳转地址模式  get   请求重定向(http://localhost:8080/index)
	//可以将请求提交到/login的URL中进行处理,
	fmt.Println("注册成功")
	r.Redirect("/login.html",302)

}
