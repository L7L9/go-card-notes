package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lqlzzz/go-card-notes/model/common/request"
	"lqlzzz/go-card-notes/model/common/response"
	"lqlzzz/go-card-notes/model/schema"
	"lqlzzz/go-card-notes/utils"
)

type UserApi struct {
}

// ChangePassword //
// 修改密码
func (api *UserApi) ChangePassword(c *gin.Context) {
	// 接收参数
	var changePasswordRequest request.ChangePasswordRequest
	err := c.ShouldBindJSON(&changePasswordRequest)
	if err != nil {
		response.FailedWithMsg(c, err.Error())
		return
	}
	// 获取用户id
	userID := utils.GetUserID(c)
	user := &schema.User{Model: gorm.Model{ID: userID}, Password: changePasswordRequest.Password}
	if err = userService.ChangePassword(user, changePasswordRequest.NewPassword); err != nil {
		response.FailedWithMsg(c, err.Error())
		return
	}
	response.SuccessWithMsg(c, "修改密码成功")
}
