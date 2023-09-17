package request

// SignUpRequest //
// 注册请求
type SignUpRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
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
	Nickname string `json:"nickname"`
	HeadImg  string `json:"headImg"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
