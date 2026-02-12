package service

import (
	"errors"
	"go-web/global"
	"go-web/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(req RegisterRequest) (RegisterResponse, error) {
	var user models.SysUser
	// 查询用户名是否存在
	err := global.DB.Where("username = ?", req.Username).First(&user).Error
	if err == nil {
		return RegisterResponse{}, errors.New("用户名已存在")
	}
	// 数据库异常(只要不是「没查到数据」，那就一定是数据库出问题了)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return RegisterResponse{}, err
	}
	// 密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return RegisterResponse{}, err
	}
	// 创建用户
	newUser := models.SysUser{
		Username: req.Username,
		Password: string(hashPassword),
	}
	err = global.DB.Create(&newUser).Error
	if err != nil {
		return RegisterResponse{}, err
	}
	// 返回响应
	return RegisterResponse{
		UserID: newUser.ID,
	}, nil
}
