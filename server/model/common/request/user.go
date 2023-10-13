package request

import "mime/multipart"

// SignUpRequest //
// 注册请求
type SignUpRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
}

// SignInRequest //
// 登录请求
type SignInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ChangePasswordRequest //
// 修改密码请求
type ChangePasswordRequest struct {
	Password    string `json:"password" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

// UpdateUserInformationRequest //
// 修改用户信息请求
type UpdateUserInformationRequest struct {
	Nickname string                `form:"nickname" json:"nickname"`
	HeadImg  *multipart.FileHeader `form:"headImg" json:"headImg"`
	Phone    string                `form:"phone" json:"phone"`
	Email    string                `form:"email" json:"email"`
}

// FollowRequest //
// 关注或者取关的请求
type FollowRequest struct {
	UserID   uint `json:"userId"`
	IsFollow bool `json:"isFollow"`
}
