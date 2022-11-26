package service

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
/*
	用户对数据库的交互操作
*/
const (
	//正常状态= 0  非正常= 1  删除状态2
	STATUS = 0
	LOCK   = 1
	DEL    = 2
)

/*
	登录
*/
func Login(users entity.Users) entity.Users {
	//因为我们插入数据库用的是md5加密过后的密码数据，所以我们验证的时候需要也先将密码进行加密，然后再进行查询
	pwd := utils.MD5(users.Password)
	fmt.Println("加密后的密码：", pwd)
	//根据用户名和密码进行查询数据， 返回查询到的单行数据
	sql := "select * from users where username =? and password=? and status =0 "
	//查询
	row := utils.QueryObject(sql, users.Username, pwd) //将加密过后的密码和数据库进行查询
	var user entity.Users

	row.Scan(&user.Id, &user.Username, &user.Password, &user.Status, &user.CreateTime)

	//返回查询到的单行数据的结构体
	return user

}

/*
	注册
*/
func Register(user entity.Users) (id, count int64, err error) {
	/*
		建议实际开发中，我们的私密数据不要明文传输，这不安全，
		然后我们把用户插入数据库的时候把密码加密，然后加密到数据库里面，这样就算权限再高的管理员也看不见你的密码
		后面验证登录的时候，也把数据库进行加密，然后比对和数据库加密过后的密码是否一致
		本项目采用md5加密算法加密
	*/
	sql := "insert into users(username,password,status, createtime )values(?,?,?,?)"
	//hash加密后的密码
	pwd := utils.MD5(user.Password)
	//把结构体数据和加密后的密码插入数据库
	id, count, err = utils.ExecDemo(sql, user.Username, pwd, STATUS, time.Now().Unix())
	if err != nil {
		return 0, 0, err
	}
	return id, count, nil
}

/*
	根据用户名返回查询到的结果
*/
func QueryByUserName(user entity.Users) bool {
	isExist := false
	sql := "select * from users where username = ?"
	//根据用户名进行查询
	fmt.Println("用户名:", user.Username)
	row := utils.QueryObject(sql, user.Username)
	/*
		 QueryRow执行一个最多返回一行的查询。
		注意：QueryRow总是返回一个非空值。如果没有值，错误被延迟到调用Row的Scan方法。
			如果查询没有选择行，*Row的Scan将返回ErrNoRows。
		所以我们通过判断row。Scan的方法来判断值是否存在
	*/
	/*if row == nil { //不能这样判断，因为row一般是返回的非空值，
		fmt.Println("我没查询到了数据,说明该用户名不存在")
	}*/

	//正确方法：通过row.Scan写入结构体，然后判断是否出现错误，不存在数据的话他是会报错的
	user = entity.Users{}
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Status, &user.CreateTime)

	//err不为空，就说明没有数据行，说明用户不存在
	if err != nil {
		return isExist //不存在false
	}
	//row.Scan 没有产生err的话，为空的话，就说明该用户不存在
	isExist = true //存在true
	fmt.Println("该用户名已经存在")
	return isExist
}

/*
	展示，
*/

/*
	根据用户名查询用户
*/
func QueryUserInfoByName(name string) (*entity.Users, error) {
	sql := "select * from users where username = ?"
	row := utils.QueryObject(sql, name)
	user := entity.Users{}
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Status, &user.CreateTime)

	//err不为空，就说明没有数据行，说明用户不存在
	if err != nil {
		return nil, err //不存在false
	}

	return &user, nil
}
