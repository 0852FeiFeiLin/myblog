package entity

import (
	"fmt"
	"myBlog/utils"
	"reflect"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/12/3 21:48
 **/
/*
	这个类用于储存那些标签：html,ajax,java。。。
		然后后面我们通过遍历group的值，然后
*/
type GroupClass struct{
	AJAX int
	DOM int
	JavaScript string
	Java string
	RPC string
	HTML string
	Css string
	Jsp string
	Go string
	Golang string
	Http string
	JQuery string
	MySQL string
}
/*
	该函数用于操作数据库，然后返回一个map。。。。
*/
/*func (g *GroupClass)init()map[string]int{


}*/
func (g *GroupClass) SelectGroup()map[string]int{
	var group GroupClass
	/*这个结构体用来存数据库tags 和 count的值，键值对*/
	var remap = make(map[string]int)

	//遍历结构体，啊然后依次查询tags值到数据库 group by
	var typeInfo = reflect.TypeOf(group)  //获取的是类型相关
	/*	var valInfo = reflect.ValueOf(group) //获取的是值相关*/
	num := typeInfo.NumField()
	//遍历
	for i := 0; i < num; i++ {
		//fmt.Println(i, typeInfo.Field(i).Name, valInfo.Field(i))
		tname := typeInfo.Field(i).Name
		//fmt.Println(i, typeInfo.Field(i).Name)
		//查询数据库count
		//sql := "select count(1) from article  where tags = ?"
		//sql优化
		sql := "select count(1) from article group by tags having tags = ?"
		row := utils.QueryObject(sql,tname)
		var count int
		row.Scan(&count)
		//fmt.Println(tname,count)
		remap[tname] = count
	}
	fmt.Println(remap)
	//返回键值对集合
	return remap
}