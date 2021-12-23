package models

import (
	"fmt"
	"myBlog/entity"
	"myBlog/utils"
	"time"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/27 10:17
 **/

/*有关个人博客的操作*/

/*
	写博客：
		把结构体的字段，一一插入数据库中
*/
func InsertBlog(art entity.Article) (id, count int64, err error) {
	return addArticle(art)
}

/*（权限控制，小写外部不可访问）*/
func addArticle(art entity.Article) (id, count int64, err error) {
	fmt.Println(art.Title)
	//sql语句
	sql := "insert into article(`title`,`author`,`tags`,`short`,`content`,`createtime`)values(?,?,?,?,?,?)"
	//结构体数据
	//调用插入方法
	id, count, err = utils.ExecDemo(sql, art.Title, art.Author, art.Tags, art.Short, art.Content, time.Now().Unix())
	if err != nil {
		return 0, 0, err
	}
	return id, count, nil
}

/*
	个人最新博客：
		根据数据库进行查找，返回博客数组对象，显示最新的博客简介在页面中
		sql思路： 先按id进行降序排序，然后limit分页，选择前3条
*/
func QueryBestNews() ([]entity.Article, error) {
	//最新文章数组对象
	var arts []entity.Article
	sql := "select id,author,short from article  order by id desc limit 0,3"
	rows, err := utils.QueryResults(sql)
	//延迟关闭资源
	defer rows.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	//解析返回结果集
	for rows.Next() {
		id := 0
		author := ""
		short := ""
		err := rows.Scan(&id, &author, &short)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		art := entity.Article{Id: id, Author: author, Short: short}
		//数组对象添加值
		arts = append(arts, art)
	}
	fmt.Println(arts)
	return arts, nil
}

/*
	根据具体类别tags，然后返回所有数据，展示在页面上，
*/
func QueryBlogsByTags(tag string) ([]entity.Article, error) {
	var blogs []entity.Article
	sql := "select * from article where tags = ?"
	rows, err := utils.QueryResults(sql, tag)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	//解析结果集
	for rows.Next() {
		//注意，row.Scan如果没有查询到行，是会报错的，所以这里要判断err
		id := 0
		var title, author, tags, short, content string
		var createTime int64

		err = rows.Scan(&id, &title, &author, &tags, &short, &content, &createTime)
		if err != nil {
			return nil, err
		}
		cretime := utils.SwitchTimeStampToData(createTime)
		fmt.Println(cretime)
		blog := entity.Article{Id: id, Title: title, Author: author, Tags: tags, Short: short, Content: content, CreateTime: createTime, CreateTime1: cretime}
		blogs = append(blogs, blog)
	}
	fmt.Println(blogs)
	return blogs, nil
}

/*
	根据Blogs传递的信息，去查找数据库，然后返回具体的详情信息结果
	这里暂时是根据标题，类型，作者进行查询的
*/
func QueryBlog(tag string, tit string, aut string) ([]entity.Article, error) {
	var blogs []entity.Article
	sql := "select * from article where tags = ? and title =? and author = ? "
	row := utils.QueryObject(sql, tag, tit, aut)
	id := 0
	var title, author, tags, short, content string
	var createTime int64
	err := row.Scan(&id, &title, &author, &tags, &short, &content, &createTime)
	if err != nil {
		return nil, err
	}
	cretime1 := utils.SwitchTimeStampToData(createTime)

	blog := entity.Article{Id: id, Title: title, Author: author, Tags: tags, Short: short, Content: content, CreateTime: createTime, CreateTime1: cretime1}
	blogs = append(blogs, blog)
	return blogs, nil
}

/*
	删除博客：根究tags，title，author，
*/
func DeleteBlog(tag string, tit string, aut string) (int64, error) {
	sql := "delete from article where tags = ? and title = ? and author = ?"
	_, count, err := utils.ExecDemo(sql, tag, tit, aut)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	//返回受影响的行数
	return count, nil
}

/*
	更新操作：
*/
func UpdateBlog(art entity.Article) (id, count int64, err error) {
	sql := "update article set title =? ,tags=?,short = ?,content=?,createtime=? where id=?"
	i, c, err := utils.ExecDemo(sql, art.Title, art.Tags, art.Short, art.Content, time.Now().Unix(), art.Id)
	if err != nil {
		return 0, 0, nil
	}
	return i, c, err

}
