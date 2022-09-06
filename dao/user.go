package dao

import (
	"log"
	"ms-go-blog/models"
)

func GetUserNameById(userId int) string {
	row := DB.QueryRow("select user_name from blog_user where uid = ?", userId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var userName string
	_ = row.Scan(&userName)
	return userName
}

func GetUser(userName, passwd string) *models.User {
	row := DB.QueryRow(
		"select * from blog_user where user_name = ? and passwd = ? limit 1",
		userName, passwd)
	if row.Err() != nil {
		log.Println(row.Err())
		return nil
	}
	var user models.User
	err := row.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &user
}

