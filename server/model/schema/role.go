package schema

import "gorm.io/gorm"

// Role //
// gcn_role table
type Role struct {
	gorm.Model
	Name string
}

func (Role) TableName() string {
	return "gcn_role"
}
