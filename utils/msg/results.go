package msg

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//type Response struct {
//	Code        int           `json:"code"`
//	Success     bool          `json:"success"`
//	Msg         string        `json:"msg"`
//	Timestamp   int64         `json:"timestamp"`
//	Result      interface{}   `json:"result"`
//}

type Response struct {
	ErrorCode    int         `json:"errorCode"`
	Success      bool        `json:"success"`
	ErrorMessage string      `json:"errorMessage"`
	Timestamp    int64       `json:"timestamp"`
	ShowType     int         `json:"showType"`
	Data         interface{} `json:"data"`
	Host         string      `json:"host"`
}

type PageInfo struct {
	Response
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
}

//func (r *Response) Error() string {
//	return r.ErrorMessage.Error()
//}

func Result(data interface{}, msg error, showType int, success bool, c *gin.Context) {
	var r = &Response{
		ErrorCode:    http.StatusBadRequest,
		Success:      false,
		ErrorMessage: msg.Error(),
		ShowType:     showType,
		Timestamp:    time.Now().Unix(),
		Data:         nil,
		Host:         c.ClientIP(),
	}
	if success {
		r.Success = true
		r.ErrorCode = http.StatusOK
		c.JSON(http.StatusOK, r)
	}
	r.Data = data
	c.JSON(http.StatusBadRequest, r)
}

func ResultWithPageInfo(data interface{}, msg error, showType int, success bool, total int64, page, size int, c *gin.Context) {
	c.JSON(http.StatusOK, &PageInfo{
		Response: Response{
			ErrorCode:    http.StatusOK,
			Success:      success,
			ErrorMessage: msg.Error(),
			ShowType:     showType,
			Timestamp:    time.Now().Unix(),
			Data:         data,
			Host:         c.ClientIP(),
		},
		Total:    total,
		Page:     page,
		PageSize: size,
	})
}

//func Result(code int, data interface{}, msg string, success bool, c *gin.Context) {
//	c.JSON(code, &Response{
//		Code: code,
//		Success: success,
//		Msg:  msg,
//		Timestamp: time.Now().Unix(),
//		Result: data,
//	})
//}

//func Success(c *gin.Context) {
//	Result(http.StatusOK, map[string]interface{}{}, "操作成功", true, c)
//}
//
//func SuccessWithMessage(message string, c *gin.Context) {
//	Result(http.StatusOK, map[string]interface{}{}, message, true, c)
//}
//
//func SuccessWithData(data interface{}, c *gin.Context) {
//	Result(http.StatusOK, data, "操作成功", true, c)
//}
//
//func SuccessDetailed(data interface{}, message string, c *gin.Context) {
//	Result(http.StatusOK, data, message, true, c)
//}
//
//func Fail(c *gin.Context) {
//	Result(http.StatusBadRequest, map[string]interface{}{}, "操作失败", false, c)
//}
//
//func FailWithMessage(message string, c *gin.Context) {
//	Result(http.StatusBadRequest, map[string]interface{}{}, message, false, c)
//}
//
//func FailWithDetailed(data interface{}, message string, c *gin.Context) {
//	Result(http.StatusBadRequest, data, message, false, c)
//}
