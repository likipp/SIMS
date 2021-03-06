package msg

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Response struct {
	Code        int           `json:"code"`
	Success     bool          `json:"success"`
	Msg         string        `json:"msg"`
	Timestamp   int64         `json:"timestamp"`
	Result      interface{}   `json:"result"`
}

func (r *Response) Error() string {
	return r.Msg
}

func Result(code int, data interface{}, msg string, success bool, c *gin.Context) {
	c.JSON(code, &Response{
		Code: code,
		Success: success,
		Msg:  msg,
		Timestamp: time.Now().Unix(),
		Result: data,
	})
}

func Success(c *gin.Context) {
	Result(http.StatusOK, map[string]interface{}{}, "操作成功", true, c)
}

func SuccessWithMessage(message string, c *gin.Context) {
	Result(http.StatusOK, map[string]interface{}{}, message, true, c)
}

func SuccessWithData(data interface{}, c *gin.Context) {
	Result(http.StatusOK, data, "操作成功", true, c)
}

func SuccessDetailed(data interface{}, message string, c *gin.Context) {
	Result(http.StatusOK, data, message, true, c)
}

func Fail(c *gin.Context) {
	Result(http.StatusBadRequest, map[string]interface{}{}, "操作失败", false, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(http.StatusBadRequest, map[string]interface{}{}, message, false, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(http.StatusBadRequest, data, message, false, c)
}