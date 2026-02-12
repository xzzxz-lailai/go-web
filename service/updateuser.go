package service

import (
	"errors"
	"go-web/global"
	"go-web/models"
)

type UpDateUserRequest struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	Remarks  string `json:"remarks"`
}

func UpdateUserByAdmin(adminID, userID uint, req UpDateUserRequest) error {
	// 校验"普通用户是否属于当前管理员"
	var relation models.SysUserRelation
	err := global.DB.
		Where("admin_id = ? AND user_id = ?", adminID, userID).
		First(&relation).Error
	if err != nil {
		return errors.New("无权修改该用户")
	}
	// 如果username不等于nil就执行查询
	if req.Username != "" {
		var count int64
		err := global.DB.
			Model(&models.SysUser{}).
			// username = ? 查有没有人用了这个用户名,    AND id != ? 排除自己（不然你自己肯定查到一条）
			Where("username = ? AND id != ?", req.Username, userID).
			Count(&count).Error
		if err != nil {
			return err
		}

		// 除了当前用户之外，已经有多少个用户使用了这个 username
		if count > 0 {
			return errors.New("用户名已存在")
		}
	}
	// 更新
	updates := map[string]interface{}{
		"username": req.Username,
		"phone":    req.Phone,
	}

	// 执行更新
	result := global.DB.Model(&models.SysUser{}).Where("id = ?", userID).Updates(updates)
	// 判断错误
	if result.Error != nil {
		return result.Error
	}
	// 受影响的行数
	if result.RowsAffected == 0 {
		return errors.New("用户不存在或数据未发生变化")
	}

	return nil
}
