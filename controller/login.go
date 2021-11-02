package controller

import (
	"SIMS/internal/gins"
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils/msg"
	"errors"
	"github.com/gin-gonic/gin"
)

type User struct {
	UserID   int `json:"userid"`
	NickName string `json:"nickname"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Access   string `json:"access"`
}

func Login(c *gin.Context) {
	var l models.Login
	err := gins.ParseJSON(c, &l)
	if err != nil {
		msg.Result(nil, msg.QueryParamsFail, 1, false, c)
		return
	}
	err, success := services.SLogin(l)
	if success {
		msg.Result(nil, err, 1, true, c)
		return
	}
	msg.Result(nil, err, 1, false, c)
	return
}

func GetCurrentUser(c *gin.Context) {

	var user = User{
		Access:   "admin",
		NickName: "周环环",
		Name:     "Eva",
		UserID: 10000,
	}
	msg.Result(user, errors.New("获取成功"), 1, true, c)
	return
}
