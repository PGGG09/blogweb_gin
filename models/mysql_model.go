package models

import (
	"blogweb_gin/utils"
	"fmt"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int
	Createtime int64
}

func InsertUser(user User) (int64, error) {
	return utils.ModifyDB("insert into user(username, password, status, create_time) values (?,?,?,?)", user.Username, user.Password, user.Status, user.Createtime)
}

// 按条件查询
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}

func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	return QueryUserWightCon(sql)
}
