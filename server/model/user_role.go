package model

// UserRole //
// gcn_user_role table
type UserRole struct {
	ID     uint `gorm:"primarykey"`
	UserID uint `gorm:"column:user_id"`
	RoleID uint `gorm:"column:role_id"`
}

func (UserRole) TableName() string {
	return "gcn_user_role"
}
