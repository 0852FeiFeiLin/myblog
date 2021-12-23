package controllers

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

func (a *AlbumController)ToAlbum()  {
	a.TplName = "album.html"
}