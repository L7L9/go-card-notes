package service

type ServiceGroup struct {
	CasbinService
	BaseService
	UserService
}

var ServiceOuter = new(ServiceGroup)
