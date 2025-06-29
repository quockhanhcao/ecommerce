package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"column:uuid;type:varchar(255);not null;index:idx_uuid;unique"`
	UserName string    `gorm:"column:user_name"`
	IsActive bool      `gorm:"column:is_active;default:true;type:boolean"`
	Roles    []Role    `gorm:"many2many:go_users_roles"`
}

func (u *User) TableName() string {
	return "go_db_users"
}
