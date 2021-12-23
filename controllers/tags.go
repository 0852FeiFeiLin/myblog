package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/27 18:22
 **/
type TagsController struct {
	//BaseController
	beego.Controller  //这个删掉了就不报错了
}

/*
	标签页面
*/
func (t *TagsController) ToTags(){
	t.TplName = "tags.html"
	t.Redirect("/tags",302)
	fmt.Println("我要去获取count数据了......")//放到了blog控制器里面，不好意思
}
/*
	处理标签:
		1.显示count数量吗，写在了blogcontroller中，
		2.显示最新文章功能，在tags里面
*/
func (t *TagsController)Tags()  {

}