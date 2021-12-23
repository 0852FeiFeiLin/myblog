package entity

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/27 15:57
 **/
/*
	这个对应user实体类
*/
type Users struct{
	Id int
	Username string `form:"username"`
	Password string `form:"password"`
	Status int64 `json:"status"`
	CreateTime int64 `json:"create_time"`
}