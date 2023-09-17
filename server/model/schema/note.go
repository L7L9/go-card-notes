package schema

import "gorm.io/gorm"

// Note //
// gcn_note table
type Note struct {
	gorm.Model
	Title        string `json:"title" gorm:"NOT NULL;comment:笔记标题"`
	AuthorName   string `json:"authorName" gorm:"NOT NULL;comment:作者昵称"`
	AuthorID     uint   `json:"authorID" gorm:"NOT NULL;comment:作者ID"`
	Uri          string `json:"uri" gorm:"NOT NULL;comment:笔记标题"`
	CollectCount int    `json:"collectCount" gorm:"default:0;comment:笔记标题"`
	Tags         []Tag  `json:"tags" gorm:"many2many:'gcn_note_tag'"`
	Status       int    `json:"status" gorm:"comment:笔记状态(0=>公开;1=>私藏)"`
}

func (Note) TableName() string {
	return "gcn_note"
}
