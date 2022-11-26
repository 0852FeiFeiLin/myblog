package controllers

import (
	"encoding/json"
	"fmt"
	"myBlog/entity"
	"myBlog/service"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/27 18:11
 **/
type BlogInfoController struct {
	//beego.Controller
	BaseController
}

/*
	显示博客详情的逻辑，就是通过小窗口就定位到博客详情页面。
	注意，blog详情只有一个页面，通过查询数据库内容进行动态展示博客详情内容，然后也能
*/
func (b *BlogInfoController) BlogInfo() {
	b.TplName = "bloginfo.html"
	//通过该方法找到url?前面的值
	value := b.Ctx.Request.Referer()
	fmt.Println(value)
	//方式1
	tag := b.GetString("tag")
	title := b.GetString("title")
	author := b.GetString("author")
	fmt.Println(tag+","+title+",", author)
	/*//方式2
	var tag2 = b.Input().Get("tag")*/

	//根据tag,title,author去数据库中查找,返回结果集
	Blogs, err := service.QueryBlog(tag, title, author)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//创建map
	/*	Blogs :=map[string]interface{}{
		"Blogs" : blogs,
	}*/
	//把blogs结果集结构体发送给前端
	b.Data["BlogInfo"] = Blogs
	b.Data["IsLogin"] = b.IsLogin
}

func (b *BlogInfoController) DeleteBlog() {
	b.TplName = "bloginfo.html"
	//通过该方法找到url?后面的值
	value := b.Ctx.Request.Referer()
	fmt.Println(value)
	//获取url后面跟的参数
	tag := b.GetString("tag")
	title := b.GetString("title")

	fmt.Println(b.GetSession("userName").(string))
	//查询数据库，然后删除操作,
	count, err := service.DeleteBlog(tag, title, b.GetSession("userName").(string))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if count == 0 {

		fmt.Println("删除失败，您没有权限")
		return
	}
	b.Redirect("blog.html",302)
	fmt.Println("删除成功")

}

/*
	删除博客：z这个是利用Ajax回传数据，然后想实现页面内容传递，但是由于太长了，没成功好像
*/
func (b *BlogInfoController) DeleteBlog0() {
	b.TplName = "bloginfo.html"
	ajax := b.IsAjax()
	fmt.Println("是否为AJAX:", ajax)
	//获取从ajax传过来的数据
	data := b.Ctx.Input.RequestBody
	fmt.Println("前端ajax传递的数据:", string(data))
	var art entity.Article
	//将数据解析到结构体里面
	json.Unmarshal(data, &art)
	fmt.Println(art)
}
/*
	编辑博客：
*//*
func (b *BlogInfoController) EditBlog() {

	b.Redirect("bloginfo",302)


}*/
