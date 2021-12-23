package entity

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/27 16:02
 **/
/*
	这个对应写博客和博客详情内容的实体，article
*/
type Article struct {
	Id int `json:"id"`
	Title string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
	Tags string `json:"tags" form:"tags"`
	Short string `json:"short" form:"short"`
	Content string `json:"content" form:"content"`
	CreateTime int64 `json:"create_time"`
	CreateTime1 string `json:"create_time_1"`
}
