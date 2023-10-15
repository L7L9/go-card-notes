package request

// CreateNoteRequest //
// 创建笔记请求
type CreateNoteRequest struct {
	Title  string `form:"title" json:"title"`
	Group  string `form:"Group" json:"group"`
	Status int    `form:"status" json:"status"`
}

// CreateNoteGroupRequest //
// 创建笔记分组
type CreateNoteGroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
