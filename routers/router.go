package routers

import (
	"myBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//跳转登录
	beego.Router("/", &controllers.IndexController{}, "get:ToIndex")
	beego.Router("/blog.html", &controllers.BlogController{}, "get:ToBlog")

	//登录
	beego.Router("/login.html", &controllers.IndexController{}, "get:ToIndex")
	beego.Router("/login", &controllers.IndexController{}, "post:Index")
	//注册
	beego.Router("/register.html", &controllers.RegisterController{}, "get:ToRegister")
	beego.Router("/register", &controllers.RegisterController{}, "post:Register")

	//标签
	beego.Router("/tags.html", &controllers.TagsController{}, "get:ToTags")
	//问题：a标签每次提交，跳转是get请求，怎么让他变为post请求
	//beego.Router("/tags",&controllers.TagsController{},"get:ToTags")
	beego.Router("/blogs.html", &controllers.BlogsController{}, "get:ToBlogs")
	beego.Router("/blogs", &controllers.BlogsController{}, "post:Blogs")
	//个人相册
	beego.Router("/album.html", &controllers.AlbumController{}, "get:ToAlbum")
	//beego.Router("/album",&controllers.AlbumController{},"post:Album")

	//关于我
	beego.Router("/editUser.html", &controllers.EditUserController{}, "get:ToEditUser")
	beego.Router("/editUser", &controllers.EditUserController{}, "post:EditUser")

	//博客列表
	beego.Router("/blogs", &controllers.BlogsController{}, "get:Blogs")
	//博客详情
	beego.Router("/bloginfo", &controllers.BlogInfoController{}, "get:BlogInfo")
	//写博客
	beego.Router("/writeBlog.html", &controllers.WriteBlogController{}, "get:ToWriteBlog")
	beego.Router("/writeBlog", &controllers.WriteBlogController{}, "post:WriteBlog")
	//退出写博客页面
	beego.Router("/exit", &controllers.ExitController{}, "get:Exit")

	//编辑博客
	/*	beego.Router("/editBlog.html",&controllers.EditBlogController{},"get:ToEditBlog")*/
	beego.Router("/editBlog", &controllers.EditBlogController{}, "post:EditBlog")
	//删除博客
	beego.Router("/deleteBlog", &controllers.BlogInfoController{}, "get:DeleteBlog")

}
