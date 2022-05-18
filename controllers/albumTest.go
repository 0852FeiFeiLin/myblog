package controllers

import (
	"fmt"
	"log"
	"myBlog/entity"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2022/5/7 19:37
 * @Description:
 **/

type AlbumTestController struct {
	BaseController
}

func (a *AlbumTestController) ToAlbumTest() {
	a.TplName = "albumTest.html"
}

func (a *AlbumTestController) AlbumTest() {

	fmt.Println("是否是ajax请求：", a.IsAjax())
	album := entity.Album{}
	Albums, err := album.FindAllAlbum()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("查询到的图片数量：", len(Albums))
	imgPath := make([]string, 0)
	for _, a2 := range Albums {
		fmt.Println("路径：", a2.FilePath)
		imgPath = append(imgPath, a2.FilePath)
	}
	fmt.Println(imgPath)
	a.Data["Album"] = Albums  //结构体切片
	a.Data["json"] = imgPath  //切片
	a.TplName = "albumTest.html"
}

/*
 */
