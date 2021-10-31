package models

import (
	"SIMS/config"
	"SIMS/utils"
	"errors"
)

type Login struct {
	Username string `form:"UserName" json:"username" binding:"required"`
	Password string `form:"Password" json:"password" binding:"required"`
}

func (l *Login) GetUser() (user Login, err error) {
	user = Login{
		Username: "周环环",
	}
	return user, nil

}

func UserLogin(l *Login) (err error, success bool) {
	var login = config.AdminConfig.Login
	//password := utils.PasswordHash("eva628.")
	//fmt.Println(password, "password")
	if l.Username != login.Username {
		return errors.New("用户错误"), false
	}
	if utils.PasswordVerify(login.Password, l.Password) != true {
		return errors.New("密码错误"), false
	}
	return errors.New("登录成功"), true
}
