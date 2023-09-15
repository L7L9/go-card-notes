package response

import "lqlzzz/go-card-notes/model/schema"

// SignInResponse //
// 登录响应
type SignInResponse struct {
	Token string
	User  schema.User
}
