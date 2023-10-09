package router

import (
	"github.com/gin-gonic/gin"
	v1 "lqlzzz/go-card-notes/api/v1"
)

// UserRouter //
// 用户路由
type UserRouter struct {
}

func (router *UserRouter) Initialize(fatherRouter *gin.RouterGroup) {
	userApi := v1.ApiOuter.UserApi

	userRouter := fatherRouter.Group("/user")
	{
		userRouter.PUT("/changePassword/", userApi.ChangePassword)
		userRouter.PUT("/updateUserInformation/", userApi.UpdateUserInformation)
		userRouter.POST("/operateFollow/", userApi.FollowOrNot)
		userRouter.GET("/getFollowList/", userApi.GetFollowList)
		userRouter.GET("/getFollowerList/", userApi.GetFollowerList)
	}
}
