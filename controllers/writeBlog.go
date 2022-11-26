package controllers

import (
	"fmt"
	"myBlog/entity"
	"myBlog/service"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/27 22:21
 **/

type WriteBlogController struct {
	//beego.Controller
	BaseController
}
/*
	跳转到写博客的页面
*/
func (w *WriteBlogController) ToWriteBlog(){
	w.TplName = "writeBlog.html"
}

/*
	处理写完博客的内容提交到数据库
*/

func (w *WriteBlogController)WriteBlog(){
	//解析前端数据，然后反序列化存入结构体，然后insert进数据库，返回插入结果
/*	var article entity.Article
	w.ParseForm(&article)*/
	var title = w.GetString("title")
	var tags = w.GetString("tags")
	var short = w.GetString("short")
	var content = w.GetString("content")

	//获取到持续连接的对象用户名当作作者
	var author = w.GetSession("userName").(string)
	fmt.Println(author)

	 article := entity.Article{
	 	Title: title,
	 	Author: author, //作者
	 	Tags: tags,
	 	Short: short,
	 	Content: content,
	 }

	/*这个好像会受到限制
	*/
	fmt.Println(article)
	//将结构体数据插入数据库，然后路由重定向到tags页面，并显示添加的内容。
	_, _, err := service.InsertBlog(article)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("发布成功!")
	//执行那个页面，然后显示所有文章，并把新插入的那条显示在最上面
	/*w.TplName = "bloginfo.html"*/
	//w.Redirect("/bloginfo.html",302)
	w.Data["info"] = "发布成功"
	w.Redirect("/blog.html",302)
}