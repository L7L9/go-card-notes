package v1

import (
	"github.com/gin-gonic/gin"
	"lqlzzz/go-card-notes/model"
	"lqlzzz/go-card-notes/model/common/request"
	"lqlzzz/go-card-notes/model/common/response"
	"lqlzzz/go-card-notes/utils"
)

type NoteApi struct{}

// CreateNote //
// 创建笔记接口
func (api *NoteApi) CreateNote(c *gin.Context) {

}

// CreateNoteGroup //
// 创建笔记分组
func (api *NoteApi) CreateNoteGroup(c *gin.Context) {
	var createNoteGroupReq request.CreateNoteGroupRequest
	err := c.ShouldBindJSON(createNoteGroupReq)
	if err != nil {
		response.FailedWithMsg(c, err.Error())
		return
	}

	userID := utils.GetUserID(c)
	group := &model.NoteGroup{
		CreatorID:   userID,
		Name:        createNoteGroupReq.Name,
		Description: createNoteGroupReq.Description,
	}
	if err = noteService.CreateNoteGroup(group); err != nil {
		response.FailedWithMsg(c, "创建笔记分组失败")
		return
	}
	response.SuccessWithMsg(c, "创建笔记分组成功")
}

// GetUserGroup //
// 获取用户的笔记分组
