package router

import "github.com/gin-gonic/gin"

type NoteRouter struct {
}

// Initialize //
// 初始化note组路由
func (router *NoteRouter) Initialize(fatherGroup *gin.RouterGroup) {
	noteRouter := fatherGroup.Group("/note")
	{
		noteRouter.POST("/createNote/")
	}
}
