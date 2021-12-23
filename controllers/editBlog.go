package controllers

import (
	"encoding/json"
	"fmt"
	"myBlog/entity"
	"myBlog/models"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/27 22:29
 **/

type EditBlogController struct {
	//beego.Controller
	BaseController
}

/*
	跳转到编辑博客页面
*/
func (e *EditBlogController) ToEditBlog() {

	//e.TplName = "editBlog.html"
	e.TplName = "bloginfo.html"
/*	tag := e.GetString("tag")
	title := e.GetString("title")
	author := e.GetSession("userName").(string)
	short := e.GetString("short")
	content := e.GetString("content")
	fmt.Println(tag+","+title+",", author+","+short,",",content)
	art := entity.Article{0,title,author,tag,title,content,time.Now().Unix()}
	_, count, err := models.UpdateBlog(art)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if count ==0{
		fmt.Println("没有修改到受影响的行数")
		return
	}
	fmt.Println("内容更新成功")
*/
}
/*
	这里是利用Ajax提交form表单，然后进行update数据操作，将用户修改的内容上传到数据库中。
*/
func (e *EditBlogController) EditBlog() {
	ajax := e.IsAjax()
	fmt.Println("是否为Ajax请求：",ajax)
	//获取前端发送的数据
	blogs := e.Ctx.Input.RequestBody
	//打印查看
	fmt.Println(string(blogs))
	//存入结构体中
	var art entity.Article
	json.Unmarshal(blogs,&art)
	fmt.Println("成功序列化到结构体中：\n",art)
	//调用数据库update方法
	id, count, err := models.UpdateBlog(art)
	if err != nil {
		fmt.Println(err.Error())
		e.Data["json"] = false
		return
	}
	fmt.Println("更新博客内容成功")
	fmt.Println("受影响的id：",id," ","受影响的行数",count)
	//如果编辑成功，那么返回true给前端
	e.Data["json"] = "true"
	e.ServeJSON()//发送数据给前端，发送的是一个json数据

}
