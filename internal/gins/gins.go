package gins

import (
	"SIMS/utils/msg"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	UserIDKey = "/user-id"
)

type PageList interface {
	GetList(interface{}) (error, []interface{}, int64)
}

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

func ParseJSONWithPath(c *gin.Context, obj interface{}) error {
	fmt.Println(c.FullPath(), "fullPath")
	fmt.Println(c.Request.URL, "URL")
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
		msg.Result(nil, msg.QueryParamsFail, 1, false, c)
	}
	msg.Result(nil, msg.GetSuccess, 1, true, c)
}

func ParamQuery(c *gin.Context, obj interface{}) (err error, action string) {
	if action = c.DefaultQuery(obj.(string), "*"); action == "" {
		return msg.QueryParamsFail, ""
	}
	return nil, action
}

func GetUserID(c *gin.Context) string {
	return c.GetString(UserIDKey)
}

func SetUserID(c *gin.Context, userID string) {
	c.Set(UserIDKey, userID)
}
