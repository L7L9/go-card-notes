package router

type RouterGroup struct {
	BaseRouter
	UserRouter
}

var RouterGroupOuter = new(RouterGroup)
