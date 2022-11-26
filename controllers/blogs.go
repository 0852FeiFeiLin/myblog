package controllers

import (
	"fmt"
	"myBlog/service"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/5 9:40
 **/
/*
	这个是展示那些博文集合的函数
*/

type BlogsController struct {
	BaseController
}

func (b *BlogsController) ToBlogs() {
	b.TplName = "blogs.html"
	fmt.Println("blogs..........")
}

/*
	就是按tag查询，分类展示在blogs页面中，就是有很多小窗口的页面，
	然后我们能通过这些小窗口，然后根据具体的url进入到blog详情中，
*/
func (b *BlogsController) Blogs() {
	b.TplName = "blogs.html"
	//到时候显示具体内容也写在这个函数，就是跳转页面后也显示内容
	//通过该方法找到url?后面的值
	value := b.Ctx.Request.Referer()
	fmt.Println(value)
	//方式1
	tag := b.GetString("tag")
	/*//方式2
	var tag2 = b.Input().Get("tag")*/
	fmt.Println(tag)
	//利用tag去数据库中group查找,返回结果集
	Blogs, err := service.QueryBlogsByTags(tag)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//创建map
	/*	Blogs :=map[string]interface{}{
		"Blogs" : blogs,
	}*/
	//把blogs结果集结构体发送给前端
	b.Data["Blogs"] = Blogs
}
