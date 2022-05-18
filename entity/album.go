package entity

import (
	"myBlog/utils"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/27 16:03
 **/
/*
	这个对应的是个人相册实体，album
*/
type Album struct {
	Id         int `json:"id"`
	FilePath   string `json:"file_path"`
	FileName   string `json:"file_name"`
	Status     int64  `json:"status"`
	CreateTime string `json:"create_time"`
}

/*
	将图片上传到数据库，并返回结果
*/
func (a *Album) InsertAlbum(album Album) (int64, error) {
	sql := "INSERT INTO album(filepath,filename,`status`,createtime) VALUES(?,?,?,?)"
	//状态为0是正常图片，
	_, count, err := utils.ExecDemo(sql, album.FilePath, album.FileName, 0, album.CreateTime)
	return count, err
}

/*
	查询所有图片
*/
func (a *Album) FindAllAlbum()(album []Album, err error) {
	sql := "select * from album"
	results, err := utils.QueryResults(sql)
	var albums []Album
	for results.Next() {
		id := 0
		var filepath, filename ,createTime string
		var status int64
		err := results.Scan(&id,&filepath,&filename,&status,&createTime)
		if err != nil {
			return nil,err
		}
		album := Album{
			Id: id,
			FilePath: filepath,
			FileName: filename,
			Status:  status,
			CreateTime: createTime,
		}
		//添加到切片中
		albums = append(albums, album)
	}
	return albums,nil
}
