package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lqlzzz/go-card-notes/model"
	"lqlzzz/go-card-notes/model/common/dto"
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

// FollowOrNot //
// 关注或者取消关注
func (api *UserApi) FollowOrNot(c *gin.Context) {
	var followRequest request.FollowRequest
	err := c.ShouldBindJSON(&followRequest)
	if err != nil {
		response.FailedWithMsg(c, err.Error())
		return
	}

	userID := utils.GetUserID(c)
	userFollow := &model.UserFollow{
		UserID:   userID,
		FollowID: followRequest.UserID,
		Status:   followRequest.IsFollow,
	}
	if err = userService.OperateFollow(userFollow); err != nil {
		response.FailedWithMsg(c, "关注失败")
		return
	}
	response.SuccessWithMsg(c, "操作成功")
}

// GetFollowerList //
// 获取粉丝列表
func (api *UserApi) GetFollowerList(c *gin.Context) {
	// 获取用户id
	userID := utils.GetUserID(c)
	if userList, err := userService.GetFollowerList(userID); err != nil {
		response.FailedWithMsg(c, "获取粉丝列表失败")
	} else {
		var getUserListResponse response.GetUserListResponse
		for _, v := range userList {
			userInfo := dto.UserInfo{
				Username: v.Username,
				Nickname: v.Nickname,
				HeadImg:  v.HeadImg,
			}
			getUserListResponse.UserList = append(getUserListResponse.UserList, userInfo)
		}
		response.SuccessWithDetail(c, "操作成功", getUserListResponse)
	}
}

// GetFollowList //
// 获取关注列表
func (api *UserApi) GetFollowList(c *gin.Context) {
	// 获取用户id
	userID := utils.GetUserID(c)
	if userList, err := userService.GetFollowList(userID); err != nil {
		response.FailedWithMsg(c, "获取粉丝列表失败")
	} else {
		var getUserListResponse response.GetUserListResponse
		for _, v := range userList {
			userInfo := dto.UserInfo{
				Username: v.Username,
				Nickname: v.Nickname,
				HeadImg:  v.HeadImg,
			}
			getUserListResponse.UserList = append(getUserListResponse.UserList, userInfo)
		}
		response.SuccessWithDetail(c, "操作成功", getUserListResponse)
	}
}
