package controller

import (
	"SIMS/internal/gins"
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils/msg"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
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
		Name:     "eva",
	}
	c.JSONP(http.StatusOK, user)
}
