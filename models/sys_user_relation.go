package models

import (
	"gorm.io/gorm"
	"time"
)

type SysUserRelation struct {
	ID        uint           `gorm:"column:id"`
	AdminID   uint           `gorm:"column:admin_id"`
	UserID    uint           `gorm:"column:user_id"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (SysUserRelation) TableName() string {
	return "sys_user_relation"
}
