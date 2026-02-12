package service

import (
	"errors"
	"go-web/global"
	"go-web/models"
	"go-web/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(req LoginRequest, ip, userAgent string) (LoginResponse, error) {
	var user models.SysUser
	// 查询用户名是否存在
	err := global.DB.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		// 判断用户是否存在（业务错误）
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return LoginResponse{}, errors.New("用户不存在")
		}
		// 判断数据库的err
		return LoginResponse{}, err
	}
	// 校验检查密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		// 记录登陆日志操作信息 (密码错误)
		utils.RecordLoginLog(user.ID, user.Username, ip, userAgent, 0, "密码错误")

		return LoginResponse{}, errors.New("密码错误")
	}

	// 生成 Token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return LoginResponse{}, err
	}

	// 记录登陆日志操作信息 (登陆成功)
	utils.RecordLoginLog(user.ID, user.Username, ip, userAgent, 1, "登陆成功")
	// 返回 Token
	return LoginResponse{
		Token: token,
	}, nil

}
