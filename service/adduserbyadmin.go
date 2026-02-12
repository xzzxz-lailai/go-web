package service

import (
	"errors"
	"go-web/global"
	"go-web/models"
	"go-web/utils"
	"gorm.io/gorm"
)

const (
	RoleUser = "user"
)

type AddUserRequest struct {
	Username string `json:"username" `
	Password string `json:"password" `
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func AddUserByAdmin(adminID uint, req AddUserRequest) error {
	var user models.SysUser
	// 查询用户名是否存在
	err := global.DB.Where("username = ?", req.Username).First(&user).Error
	if err == nil {
		return errors.New("添加失败,用户名已存在")
	}
	// 数据库异常(只要不是「没查到数据」，那就一定是数据库出问题了)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// 密码加密
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	// 创建普通用户
	newUser := models.SysUser{
		Role:     RoleUser, // 角色 = 普通用户
		Username: req.Username,
		Password: hashPassword,
		Phone:    req.Phone,
		Email:    req.Email,
	}
	err = global.DB.Create(&newUser).Error
	if err != nil {
		return err
	}

	// 建立管理员-用户关系
	relation := models.SysUserRelation{
		AdminID: adminID,
		UserID:  newUser.ID,
	}
	err = global.DB.Create(&relation).Error
	if err != nil {
		return err
	}

	return nil
}
