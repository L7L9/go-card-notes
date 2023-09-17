package router

type RouterGroup struct {
	BaseRouter
	UserRouter
	NoteRouter
}

var RouterGroupOuter = new(RouterGroup)
