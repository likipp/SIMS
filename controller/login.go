package controller

import (
	"SIMS/config"
	"SIMS/global"
	"SIMS/init/cookies"
	"SIMS/internal/gins"
	"SIMS/middleware"
	"SIMS/models"
	"SIMS/services"
	"SIMS/utils/msg"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	UserID   string `json:"userid"`
	NickName string `json:"nickname"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Access   string `json:"access"`
}

type LoginResult struct {
	Status  string `json:"status"`
	Success bool   `json:"success"`
	Type    string `json:"type"`
}

func Login(c *gin.Context) {
	var l models.Login
	var u User
	var ls LoginResult
	err := gins.ParseJSON(c, &l)
	if err != nil {
		msg.Result(nil, msg.QueryParamsFail, 1, false, c)
		return
	}
	err, success := services.SLogin(l)
	if success {
		u.Name = "Eva"
		u.Access = "admin"
		u.NickName = "周环环"
		u.UserID = "10000"
		GetTokenAndSession(c, u)
		ls.Success = true
		ls.Status = "ok"
		ls.Type = "account"
		//msg.Result(ls, err, 0, true, c)
		c.JSONP(http.StatusOK, ls)
		return
	}
	ls.Status = "error"
	ls.Success = false
	ls.Type = "account"
	msg.Result(nil, err, 1, false, c)
	return
}

func GetCurrentUser(c *gin.Context) {
	session, err := cookies.GetSession(c)
	if err != nil {
		msg.Auth(http.StatusExpectationFailed, nil, msg.GetSessionFail, 0, false, c)
		return
	}
	var user User
	user.Avatar = session.Values["avatar"].(string)
	user.UserID = session.Values["id"].(string)
	user.NickName = session.Values["nickname"].(string)
	user.Access = "admin"
	user.Name = session.Values["name"].(string)
	msg.Result(user, errors.New("获取成功"), 0, true, c)
	return
}

func GetTokenAndSession(c *gin.Context, user User) {
	j := &middleware.JWT{
		SigningKey: []byte(config.AdminConfig.JWT.SigningKey),
	}
	id, _ := strconv.ParseUint(user.UserID, 0, 64)
	clams := models.CustomClaims{
		ID:         uint(id),
		NickName:   user.NickName,
		Username:   user.Name,
		BufferTime: 60 * 60 * 24, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 7天
			Issuer:    "xiao",                         // 签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		msg.Auth(http.StatusBadRequest, nil, msg.GetSessionFail, 2, false, c)
		return
	}
	session, _ := cookies.GetSession(c)
	session.Values["nickname"] = user.NickName
	session.Values["name"] = user.Name
	session.Values["avatar"] = "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"
	session.Values["id"] = user.UserID
	session.Values["access"] = user.Access
	session.Values["token"] = token
	global.GUser.Username = user.Name
	global.GUser.NickName = user.NickName
	cookies.SaveSession(c)
}

func Logout(c *gin.Context) {
	cookies.DeleteSession(c)
}
