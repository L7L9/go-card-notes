package service

type ServiceGroup struct {
	CasbinService
	BaseService
	UserService
	NoteService
}

var ServiceOuter = new(ServiceGroup)
