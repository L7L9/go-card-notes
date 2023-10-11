package v1

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"lqlzzz/go-card-notes/model"
	"lqlzzz/go-card-notes/model/common/request"
	"lqlzzz/go-card-notes/model/common/response"
	"lqlzzz/go-card-notes/utils"
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

	user := &model.User{
		Username: signUpRequest.Username,
		Password: signUpRequest.Password,
		Nickname: signUpRequest.Nickname,
		UUID:     uuid.NewV4(),
	}
	if err = baseService.SignUp(user); err != nil {
		response.FailedWithMsg(c, err.Error())
		return
	}
	response.SuccessWithMsg(c, "注册成功")
}

// SignIn //
// 登录api
func (api *BaseApi) SignIn(c *gin.Context) {
	var signInRequest request.SignInRequest
	err := c.ShouldBind(&signInRequest)
	if err != nil {
		response.FailedWithMsg(c, err.Error())
		return
	}

	user := &model.User{
		Username: signInRequest.Username,
		Password: signInRequest.Password,
	}

	if user, err = baseService.SignIn(user); err != nil {
		response.FailedWithMsg(c, err.Error())
		return
	}

	// 签发token
	jwt := utils.NewJwt()
	claims := jwt.GenerateClaims(user.UUID, user.ID, user.BaseRoleID, user.Username)
	token, err := jwt.GenerateToken(claims)
	if err != nil {
		response.FailedWithMsg(c, "获取token失败")
		return
	}
	response.SuccessWithDetail(c, "登陆成功", response.SignInResponse{
		User:  *user,
		Token: token,
	})
}
