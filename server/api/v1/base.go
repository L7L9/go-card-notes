package v1

import (
	"github.com/gin-gonic/gin"
	"lqlzzz/go-card-notes/model/common/request"
	"lqlzzz/go-card-notes/model/common/response"
	"lqlzzz/go-card-notes/model/schema"
)

type BaseApi struct{}

// SignUp //
// 注册api
func (api *BaseApi) SignUp(c *gin.Context) {
	var signUpRequest request.SignUpRequest
	err := c.ShouldBindJSON(&signUpRequest)
	if err != nil {
		response.FailedWithMsg(c, err.Error())
		return
	}

	user := schema.User{
		Username: signUpRequest.Username,
		Password: signUpRequest.Nickname,
		Nickname: signUpRequest.Nickname,
		Phone:    signUpRequest.Phone,
	}
	if err = baseService.SignUp(user); err != nil {
		response.FailedWithMsg(c, "注册失败，请稍后")
	}
	response.SuccessWithMsg(c, "登录成功")
}

// SignIn //
// 登录api
func (api *BaseApi) SignIn(c *gin.Context) {

}
