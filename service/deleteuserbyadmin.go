package service

import (
	"errors"
	"go-web/global"
	"go-web/models"
	"gorm.io/gorm"
)

func DeleteUserByAdmin(adminID, userID uint) error {
	var relation models.SysUserRelation
	// 校验权限：检查关系表中是否存在 admin 创建的该用户
	err := global.DB.Model(&models.SysUserRelation{}).Where("admin_id = ? AND user_id = ?", adminID, userID).First(&relation).Error
	// 没查到关系 → 说明无权操作
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("无权删除该用户")
	}
	// 数据库错误
	if err != nil {
		return err
	}
	// 删除管理员-用户关系
	err = global.DB.Delete(&relation).Error
	if err != nil {
		return err
	}
	return nil
}
