package response

import (
	"lqlzzz/go-card-notes/model"
	"lqlzzz/go-card-notes/model/common/dto"
)

// SignInResponse //
// 登录响应
type SignInResponse struct {
	Token string
	User  model.User
}

// GetUserListResponse //
// 获取用户列表
type GetUserListResponse struct {
	UserList []dto.UserInfo
}
