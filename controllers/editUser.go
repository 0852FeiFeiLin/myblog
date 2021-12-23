package controllers

import "fmt"

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/27 22:07
 **/
/*
	关于我操作，里面可以修改信息，状态，
*/
type EditUserController struct {
	//beego.Controller
	BaseController
}
/*
	跳转到关于个人信息页面
*/
func (e *EditUserController)ToEditUser()  {
	e.TplName = "editUser.html"
	e.Data["IsLogin"] = e.IsLogin
	fmt.Println("这里的session编辑用户",e.IsLogin)
	fmt.Println(e.LoginUser.(string))
}
/*
	处理用户编辑关于个人信息页面的控制器
*/
func (c *EditUserController)EditUser()  {

}