package service

import (
	"errors"
	"ms-go-blog/dao"
	"ms-go-blog/models"
	"ms-go-blog/utils"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	passwd = utils.Md5Crypt(passwd, "mszlu")
	user := dao.GetUser(userName, passwd)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}
	uid := user.Uid
	//生成 token，采用 jwt 进行生成
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token 未能生成")
	}
	var userInfo models.UserInfo
	userInfo.Uid = uid
	userInfo.UserName = userName
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginRes{
		Token: token,
		UserInfo: userInfo,
	}
	return lr, nil
}
