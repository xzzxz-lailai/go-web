package service

import (
	"errors"
	"go-web/global"
	"time"
)

// UserDetailResponse 返回管理员详细信息结构体
type UserDetailResponse struct {
	ID        uint      `json:"id"`
	Role      string    `json:"role"`
	Username  string    `json:"username"`
	Phone     string    `json:"phone"`
	Sex       string    `json:"sex"`
	Email     string    `json:"email"`
	Remarks   string    `json:"remarks"`
	CreatedAt time.Time `json:"created_at"`
}

// UserDetail 获取普通用户详细信息
func UserDetailByAdmin(adminID, userID uint) (UserDetailResponse, error) {
	var user UserDetailResponse
	err := global.DB.
		Table("sys_user AS u").
		Select("u.id, u.username,u.role,u.phone, u.email, u.created_at").
		Joins("join sys_user_relation r on u.id = r.user_id").
		Where("u.id = ? AND r.admin_id = ?", userID, adminID).
		First(&user).Error
	if err != nil {
		return user, errors.New("无权查看该用户或用户不存在")
	}
	return user, nil
}
