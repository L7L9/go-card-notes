package request

import "mime/multipart"

// CreateNoteRequest //
// 创建笔记请求
type CreateNoteRequest struct {
	Title    string                `form:"title" json:"title"`
	TextFile *multipart.FileHeader `form:"textFile" json:"textFile"`
	Tags     []string              `form:"tags" json:"tags" `
	Status   int                   `form:"status" json:"status"`
}
