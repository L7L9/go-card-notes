package response

import (
	"lqlzzz/go-card-notes/model"
)

// SignInResponse //
// 登录响应
type SignInResponse struct {
	Token string
	User  model.User
}
