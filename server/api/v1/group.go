package v1

import "lqlzzz/go-card-notes/service"

type ApiGroup struct {
	BaseApi
	UserApi
}

// ApiOuter //
// 接口的出口
var ApiOuter = new(ApiGroup)

// 接口内部需要用到的服务
var (
	baseService = service.ServiceOuter.BaseService
	userService = service.ServiceOuter.UserService
	noteService = service.ServiceOuter.NoteService
)
