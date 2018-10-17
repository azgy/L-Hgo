package model

import (
	"db/mysql"
	"time"
	"strconv"
)

type User struct {
	Id 			int 	`json:"id" xorm:"notnull pk int(11) comment('id')"`
	Name 		string 	`json:"name" xrom:"varchar(20) comment('用户名')"`
	Password 	string 	`json:"password" xrom:"varchar(20) comment('登录密码')"`
	CreateTime 	string 	`json:"createTime" xrom:"datetime created"`
}

/**
 * 保存用户
 */
func (user *User) Save() (int64, error) {
	engine, err := mysql.Client()
	if err != nil {
		return -1, err
	}
	user.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return engine.Insert(user)
}

/**
 * 查询用户信息
 */
func (user *User) Query() (bool, error) {
	engine, err := mysql.Client()
	if err != nil {
		return false, err
	}
	return engine.Get(user)
}

func (user *User) QueryByName() (bool, error) {
	engine, err := mysql.Client()
	if err != nil {
		return false, err
	}
	return engine.Cols("name").Get(user)
}

func (user *User) ToString() string {
	return "{" +
				"\"id\": " + strconv.Itoa(user.Id) + ", " +
				"\"name\": \"" + user.Name + "\", " +
				"\"password\": \"" + user.Password + "\", " +
				"\"createTime\": \"" + user.CreateTime + "\"" +
			"}"
}