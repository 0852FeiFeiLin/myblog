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
	//beego.Router("/tags",&controllers.TagsController{},"get:ToTags")
	beego.Router("/blogs.html", &controllers.BlogsController{}, "get:ToBlogs")
	beego.Router("/blogs", &controllers.BlogsController{}, "post:Blogs")

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
	//个人相册
	beego.Router("/album.html", &controllers.AlbumController{}, "get:ToAlbum")

	//上传相册文件
	beego.Router("/uploadAlbum",&controllers.AlbumController{},"post:UploadAlbum")
	beego.Router("/showAllAlbum",&controllers.AlbumController{},"post:ShowAllAlbum")

	//个人相册测试
	beego.Router("/albumTest.html",&controllers.AlbumTestController{},"get:ToAlbumTest")
	beego.Router("/albumTest",&controllers.AlbumTestController{},"post:AlbumTest")

	//个人信息
	beego.Router("/myInfo.html",&controllers.MyInfoController{},"get:ToMyInfo")
	beego.Router("/myInfo",&controllers.MyInfoController{},"get:MyInfo")

}
