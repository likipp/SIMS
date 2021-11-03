package global

import (
	"gopkg.in/boj/redistore.v1"
	"gorm.io/gorm"
)

type user struct {
	Username string `json:"username"`
	NickName string `json:"nickname" gorm:"default:'匿名用户'"`
	DeptID   uint   `json:"deptID"`
}

var (
	GDB *gorm.DB
	GRedis          *redistore.RediStore
	GUser           user
)
