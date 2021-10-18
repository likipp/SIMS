package gins

import (
	"SIMS/utils/msg"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	UserIDKey = "/user-id"
)

func Parse(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBind(obj); err != nil {
		return msg.QueryParamsFail
	}
	return nil
}

func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(&obj); err != nil {
		return msg.QueryParamsFail
	}
	return nil
}

func ParseQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		return msg.QueryParamsFail
	}
	return nil
}

func ParseForm(c *gin.Context, obj interface{}) {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		msg.Result(nil, msg.QueryParamsFail.Error(), 1, false, c)
	}
	msg.Result(nil, msg.GetSuccess.Error(), 1, true, c)
}

func GetUserID(c *gin.Context) string {
	return c.GetString(UserIDKey)
}

func SetUserID(c *gin.Context, userID string) {
	c.Set(UserIDKey, userID)
}
