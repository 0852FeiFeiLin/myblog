package controllers

import (
	"fmt"
	"log"
	"myBlog/entity"
	"os"
	"path/filepath"
	"strconv"

	"time"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/27 22:03
 **/
/*
	个人相册页面操作
*/
type AlbumController struct {
	BaseController
}

func (a *AlbumController) ToAlbum() {
	a.TplName = "album.html"
}

/*
	上传相册
*/
func (a *AlbumController) UploadAlbum() {
	file, header, err := a.GetFile("file") //返回文件，文件信息头，错误信息
	if err != nil {
		a.Ctx.WriteString("File retrieval failure")
		return
	}
	defer file.Close()          //记得延迟关闭文件上传，否则会出现临时文件不清除的情况，
	filename := header.Filename //将文件头信息赋值给filename变量
	fmt.Println("filename", filename)
	//也可以通过判断文件名后缀来判断文件类型，
	ext := filepath.Ext(filename)
	fmt.Println("ext:", ext)
	fileType := "other"
	if ext == ".jpg" || ext == ".png" || ext == ".gif" || ext == ".jpeg" || ext == ".img" {
		fileType = "images" //文件夹类型
	}

	now := time.Now()
	//文件夹路径
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	//ModePerm是0777，这样拥有该文件夹路径的执行权限
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
		a.Ctx.WriteString(err.Error())
		return
	}
	//文件路径
	timeStamp := time.Now().Unix() //时间戳
	//文件按照时间戳进行命名
	phoneName := fmt.Sprintf("%d%s", timeStamp, ext)
	fmt.Println("文件name:", phoneName)
	//拼接文件夹和路径，然后创建
	filePathStr := filepath.Join(fileDir, phoneName)
	fmt.Println("文件名：", filePathStr)
	err = a.SaveToFile("file", filePathStr) //fromFile是提交的时候的html表单中的name,参数2是保存文件的路径+文件名
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//将浏览器客户端上传的文件拷贝到本地路径的文件里面

	//将上传的文件保存到数据库中，  id path name  status createTime
	time := strconv.FormatInt(timeStamp, 10)
	album := entity.Album{
		0,
		filePathStr,
		phoneName,
		0,
		time,
	}
	count, err := album.InsertAlbum(album)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("受影响的行数：", count)
	a.TplName = "album.html"

}

func (a *AlbumController) ShowAllAlbum(){
	fmt.Println("是否是ajax请求：",a.IsAjax())
	album := entity.Album{}
	Albums, err := album.FindAllAlbum()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("查询到的图片数量：", len(Albums))
	imgPath :=	make([]string,0)
	for _, a2 := range Albums {
		fmt.Println("路径：", a2.FilePath)
		imgPath = append(imgPath, a2.FilePath)
	}

	fmt.Println(imgPath)
	//a.Data["Album"] = Albums
	a.Data["ImgPath"] = imgPath
	a.Data["json"] = strconv.FormatBool(true)
	a.TplName = "album.html"
}