package controllers

import "fmt"

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/1 20:55
 **/
type ExitController struct {
	BaseController
}
/*
用户点击退出exit，删除session
*/
func (this *ExitController) Exit(){
	//这里退出应该是返回到博客详情页面

	//删除session
	this.DelSession("userName")
	fmt.Println("Session:",this.GetSession("userName"))

	//请求重定向
	this.Redirect("/",302)
}