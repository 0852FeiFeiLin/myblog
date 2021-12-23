package main

import (
	"fmt"
	"myBlog/entity"
	"myBlog/utils"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/3 16:00
 **/

func main1() {
	utils.ConnectDB()
	var arts []entity.Article
	var art  entity.Article
	//bestNews := make(map[string]interface{})
	sql := "select author,short from article  order by id desc limit 0,3"
	//调用方法
	rows, err := utils.QueryResults(sql)
	defer rows.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if rows.Scan()!= nil{
		fmt.Println("我查询到了数据")
	}

	//解析返回结果集
	for rows.Next() {
		author := ""
		short := ""
		err := rows.Scan(&author,&short)
		if err != nil {
			fmt.Println(err.Error())
		}
		art.Author = author
		art.Short = short
		fmt.Println(art)

		arts = append(arts,art)
	}
	fmt.Println(arts)
}