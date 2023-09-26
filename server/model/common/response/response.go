package response

import (
	"github.com/gin-gonic/gin"
	"lqlzzz/go-card-notes/constant"
	"net/http"
)

// Response //
// 普通响应的结构体
type Response struct {
	StatusCode int         `json:"statusCode"`
	StatusMsg  string      `json:"statusMsg"`
	Data       interface{} `json:"data"`
}

func Result(c *gin.Context, response *Response) {
	c.JSON(http.StatusOK, response)
}

func Success(c *gin.Context) {
	Result(c, &Response{StatusCode: constant.OperateOK, StatusMsg: "操作成功"})
}

func Failed(c *gin.Context) {
	Result(c, &Response{StatusCode: constant.OperateFailed, StatusMsg: "操作失败"})
}

func SuccessWithMsg(c *gin.Context, msg string) {
	Result(c, &Response{StatusCode: constant.OperateOK, StatusMsg: msg})
}

func FailedWithMsg(c *gin.Context, msg string) {
	Result(c, &Response{StatusCode: constant.OperateFailed, StatusMsg: msg})
}

func SuccessWithData(c *gin.Context, data interface{}) {
	Result(c, &Response{StatusCode: constant.OperateOK, Data: data})
}

func FailedWithData(c *gin.Context, data interface{}) {
	Result(c, &Response{StatusCode: constant.OperateFailed, Data: data})
}

func SuccessWithDetail(c *gin.Context, msg string, data interface{}) {
	Result(c, &Response{StatusCode: constant.OperateOK, StatusMsg: msg, Data: data})
}

func FailedWithDetail(c *gin.Context, msg string, data interface{}) {
	Result(c, &Response{StatusCode: constant.OperateFailed, StatusMsg: msg, Data: data})
}
