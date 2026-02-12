package service

import (
	"errors"
	"go-web/global"
	"go-web/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ChangePassword(UserID uint, OldPassword, NewPassword, ConfirmPassword string) error {
	var user models.SysUser
	err := global.DB.Where("id = ?", UserID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("用户不存在")
		}
		return err
	}

	// 校验旧密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(OldPassword))
	if err != nil {
		return errors.New("旧密码错误")
	}

	// 校验新密码和确认密码一致性
	if NewPassword != ConfirmPassword {
		return errors.New("两次密码不一致")
	}

	// 加密新密码
	hashpass, err := bcrypt.GenerateFromPassword([]byte(NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 修改密码
	err = global.DB.Model(&models.SysUser{}).Where("id = ?", user.ID).Update("password", hashpass).Error
	if err != nil {
		return errors.New("密码更新失败")
	}

	return nil
}
