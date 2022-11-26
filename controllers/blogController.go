package controllers

import (
	"myBlog/entity"
	"myBlog/service"
	"strconv"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/1 23:06
 **/
type BlogController struct {
	//beego.Controller
	BaseController
}
/*
	这个控制器和tags是控制请是挂钩的，他的页面内容也就是tags，所以他的有一些功能在tags控制器里面
		1.显示count数量吗，写在了blogcontroller中，
	    2.显示最新文章功能，鞋子啊tags控制器中
		3.显示文章总条目
*/
func (b *BlogController)ToBlog(){
	b.TplName = "blog.html"

	var gro entity.GroupClass

	//返回查询到的键值对集合
	groMap := gro.SelectGroup()
	//定义集合，输出给前端
	/* res := map[string]interface{}{
		"Tags" : groMap,
	}*/
	// 遍历map。然后利用模板发送值到前端
	for k, _ := range groMap {
		b.Data[k] = groMap[k]
		//前端: b.Data["AJAX"] = groMap["Ajax"]
	}
	b.Data["tags"] = groMap
	b.Data["IsLogin"]= b.IsLogin
	/*b.Data["json"] = groMap*/

	//调用显示最新文章功能
	b.bestBlogs()


}
/*
	Blog控制器内部方法，小写权限控制，最新文章
*/
func (b *BlogController) bestBlogs()  {
	blogs, err := service.QueryBestNews()
	if err != nil {
		return
	}
	for i, blog := range blogs {
		//fmt.Println(i,blog.Short)
		//前端渲染
		b.Data["short"+strconv.Itoa(i)] = blog.Short
	}

}