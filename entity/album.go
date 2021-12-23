package entity

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
	Id int
	FilePath string `json:"file_path"`
	FileName string `json:"file_name"`
	Status int64 `json:"status"`
	CreateTime int64 `json:"create_time"`
}
