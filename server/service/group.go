package service

type ServiceGroup struct {
	CasbinService
	BaseService
}

var ServiceOuter = new(ServiceGroup)
