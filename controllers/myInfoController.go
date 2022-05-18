package controllers

import (
	"fmt"
	"myBlog/models"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/5/8 16:58
 * @Description:
 **/
type MyInfoController struct {
	BaseController
}

//get
func (m *MyInfoController) ToMyInfo() {
	m.TplName = "myInfo.html"
	//请求转发，这样就是post请求了，就不会还是get请求响应数据
	m.Redirect("/myInfo",307)
	//m.Ctx.Redirect(302,"/myInfo")
}

//post
func (m *MyInfoController) MyInfo() { //获取个人信息
	name := m.GetSession("userName")
	userInfo, _ := models.QueryUserInfoByName(name.(string))

	fmt.Println(userInfo)
	m.Data["Users"] = userInfo

	m.TplName = "myInfo.html"

}
