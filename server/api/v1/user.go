package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lqlzzz/go-card-notes/model"
	"lqlzzz/go-card-notes/model/common/request"
	"lqlzzz/go-card-notes/model/common/response"
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
	user := &model.User{Model: gorm.Model{ID: userID}, Password: changePasswordRequest.Password}
	if err = userService.ChangePassword(user, changePasswordRequest.NewPassword); err != nil {
		response.FailedWithMsg(c, err.Error())
		return
	}
	response.SuccessWithMsg(c, "修改密码成功")
}

// UpdateUserInformation //
// 更改用户信息
func (api *UserApi) UpdateUserInformation(c *gin.Context) {
	// 接收参数
	var updateUserInformationRequest request.UpdateUserInformationRequest
	err := c.ShouldBindJSON(&updateUserInformationRequest)
	if err != nil {
		response.FailedWithMsg(c, err.Error())
		return
	}

	userID := utils.GetUserID(c)
	user := &model.User{
		Model:    gorm.Model{ID: userID},
		Nickname: updateUserInformationRequest.Nickname,
		HeadImg:  updateUserInformationRequest.HeadImg,
		Phone:    updateUserInformationRequest.Phone,
		Email:    updateUserInformationRequest.Email,
	}
	newUser, err := userService.UpdateUserInformation(user)
	if err != nil {
		response.FailedWithMsg(c, "更改用户信息失败")
		return
	}
	response.SuccessWithDetail(c, "修改信息成功", newUser)
}
