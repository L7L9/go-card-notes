package router

import (
	"github.com/gin-gonic/gin"
	v1 "lqlzzz/go-card-notes/api/v1"
)

// BaseRouter //
// 用于注册基础路由
// 登录注册
type BaseRouter struct{}

func (r *BaseRouter) Initialize(fatherGroup *gin.RouterGroup) {
	baseApi := v1.ApiOuter.BaseApi

	baseRouter := fatherGroup.Group("/user")
	{
		baseRouter.POST("/signIn/", baseApi.SignIn)
		baseRouter.POST("/signUp/", baseApi.SignUp)
	}
}
