package cookies

import (
	"SIMS/config"
	"SIMS/global"
	"SIMS/utils/msg"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"gopkg.in/boj/redistore.v1"
)

//var RS *redistore.RediStore

func InitSession(admin config.Redis) {
	store, err := redistore.NewRediStore(10, "tcp", admin.Path, admin.Password, []byte("secret-key"))
	if err != nil {
		panic("Redis启动异常")
	}
	//store.SetMaxAge(10 * 24 * 3600)

	//store.Pool.IdleTimeout = 60*60*24*7
	//store.Options.MaxAge = 60*60*24*7
	store.SetMaxAge(60 * 60 * 24 * 7)
	store.Options.Secure = false
	store.Options.HttpOnly = true
	global.GRedis = store
}

func GetSession(c *gin.Context) (*sessions.Session, error) {
	session, err := global.GRedis.Get(c.Request, "session")
	if err != nil {
		//msg.Auth(http.StatusBadRequest, nil, msg.GetSessionFail, 2, false, c)
		//c.Abort()
		return nil, err
	}
	return session, nil
}

func SaveSession(c *gin.Context) {
	//err := sessions.Save(c.Request, c.Writer)
	if err := sessions.Save(c.Request, c.Writer); err != nil {
		msg.Result(nil, msg.SaveSessionFail, 2, false, c)
	}
}

func DeleteSession(c *gin.Context) {
	session, err := GetSession(c)
	if err != nil {
		msg.Result(nil, msg.DeleteSessionFail, 2, false, c)
		c.Abort()
		return
	}
	session.Options.MaxAge = -1
	if err := sessions.Save(c.Request, c.Writer); err != nil {
		msg.Result(nil, msg.DeleteSessionFail, 2, false, c)
		c.Abort()
		return
	}
	msg.Result(nil, msg.LoginOutSuccess, 0, true, c)
}
