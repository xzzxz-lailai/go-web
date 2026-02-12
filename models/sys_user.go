package models

import (
	"gorm.io/gorm"
	"time"
)

type SysUser struct {
	ID        uint           `gorm:"column:id"`
	Username  string         `gorm:"column:username"`
	Password  string         `gorm:"column:password"`
	Role      string         `gorm:"column:role"`
	Phone     string         `gorm:"column:phone"`
	Email     string         `gorm:"column:email"`
	Avatar    string         `gorm:"column:avatar"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
