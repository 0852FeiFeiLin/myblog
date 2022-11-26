package utils

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)

/**
 * @author: linfeifei
 * @email: 2778368047@qq.com
 * @phone: 18170618733
 * @DateTime: 2021/11/27 9:54
 **/
/*
	这一层是操作数据库的代码，里面负责连接数据库，和数据库的CRUD
	主要负责查询数据库，然后返回结果集，具体对结构集进行什么操作modules层负责。

	目的：user --> service --> utils --> DB
	提高代码复用性和重构代码，让代码看起来更优雅，灵活使用
*/
//全局变量
var sqldb *sql.DB //让程序连接一次就好了，全局变量 = 连接的局部变量

/*
	连接数据库
*/
func ConnectDB() {
	//通过键值对获取到配置文件里面的信息 conf
	config := beego.AppConfig
	//驱动
	driver := config.String("db_driverName")
	//用户名
	user := config.String("db_user")
	//密码
	pwd := config.String("db_pwd")
	//连接ip
	ip := config.String("db_ip")
	//端口 3306端口
	port := config.String("db_port")
	//数据库名
	dbname := config.String("db_name")
	//连接数据库,获取连接对象
	contStr := user + ":" + pwd + "@tcp("+ip + ":" + port+")/" + dbname
	db, err := sql.Open(driver,contStr)
	if err != nil {
		fmt.Println("连接指定数据库失败",err.Error())
		return
	}
	//让局部变量 = 全局变量，这样就程序运行中就不需要多次连接数据库
	sqldb = db
}

/*
	根据指定的sql语句，和条件，查询，返回结果集和错误  Query()
*/
func QueryResults(sql string, args ...interface{}) (*sql.Rows, error) {
	//ConnectDB()就不需要了
	/*
		query是可返回N多行的结果集
		QueryRow是返回单行数据
	*/
	rows, err := sqldb.Query(sql, args...)
	//延迟关闭
	if err != nil {
		return nil, err
	}
	//返回结果集
	return rows, nil

}

/*
	根据指定的条件，返回单行数据   QueryRow()
*/
func QueryObject(sql string, args ...interface{}) (*sql.Row){
	row := sqldb.QueryRow(sql, args...)
	return row
}

/*
	增删改操作，返回受影响的行数和最新的id，Exec()
*/
func ExecDemo(sql string,args ...interface{})(id,count int64,err error) {
	//Exec() :可以实现对数据的增删改的操作

	result, err := sqldb.Exec(sql, args...)
	if err != nil {
		fmt.Println("执行指定的sql语句输出错误：",err.Error())
		return 0,0,err
	}
	//如果是插入语句，他会返回最新的主键id
	//判断受影响的行数
	if count,err = result.RowsAffected();err != nil{
		fmt.Println("查询受影响的行数发生错误",err)
		return 0, count, nil
	}

	//组合写法
	if id, err = result.LastInsertId(); err != nil {
		fmt.Println("查询指定的ID时发生错误：", err)
		return 0, count, nil
	}

	return id, count, err
}

/*
	创建用户表：CreateModify
*/
func CreateModify(){

}