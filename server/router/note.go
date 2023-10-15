package router

import (
	"github.com/gin-gonic/gin"
	v1 "lqlzzz/go-card-notes/api/v1"
)

type NoteRouter struct {
}

// Initialize //
// 初始化note组路由
func (router *NoteRouter) Initialize(fatherGroup *gin.RouterGroup) {
	noteApi := v1.ApiOuter.NoteApi

	noteRouter := fatherGroup.Group("/note")
	{
		noteRouter.POST("/createNote/", noteApi.CreateNote)
		noteRouter.POST("/createNoteGroup/", noteApi.CreateNoteGroup)
	}
}
