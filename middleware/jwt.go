package middleware

import (
	"SIMS/config"
	"SIMS/init/cookies"
	"SIMS/models"
	"SIMS/utils/msg"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

var (
	TokenExpired     error = errors.New("token is expired")
	TokenNotValidYet error = errors.New("token not active yet")
	TokenMalformed   error = errors.New("that's not even a token")
	TokenInvalid     error = errors.New("couldn't handle this token")
)

type JWT struct {
	SigningKey []byte
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := cookies.GetSession(c)
		if err != nil {
			c.Abort()
			return
		}
		if session.Options.MaxAge < 0 {
			msg.Auth(http.StatusExpectationFailed, gin.H{
				"reload": true,
			}, msg.ExpectationFailed, 2, false, c)
			c.Abort()
			return
		}
		token, ok := session.Values["token"].(string)
		if !ok {
			msg.Auth(http.StatusExpectationFailed, gin.H{
				"reload": true,
			}, msg.ExpectationFailed, 2, false, c)
			c.Abort()
			return
		}

		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				msg.Auth(http.StatusExpectationFailed, gin.H{
					"reload": true,
				}, msg.SessionTimeout, 2, false, c)

				c.Abort()
				return
			}
			msg.Auth(http.StatusExpectationFailed, gin.H{
				"reload": true,
			}, msg.ExpectationFailed, 2, false, c)
			c.Abort()
			return
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + 60*60*24*7
			newToken, _ := j.CreateToken(*claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			c.Set("claims", claims)
		}
		c.Set("claims", claims)
	}
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(config.AdminConfig.JWT.SigningKey),
	}
}

func (j *JWT) CreateToken(claims models.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*models.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
			return nil, TokenNotValidYet
		}
		return nil, TokenNotValidYet
	}
	if token != nil {
		if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}
}
