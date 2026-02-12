package service

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-web/global"
	"go-web/models"
	"go-web/pkg"
	"math/rand"
	"strconv"
	"time"
)

// 生成6位验证码
func GenerateVerifyCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000 // 范围是 100000 ~ 999999
	return strconv.Itoa(code)
}

// redis key 设置redis key
func EmailRedisKey(UserID uint) string {
	return fmt.Sprintf("email_change:%d", UserID)
}

// SendChangeEmailCode 发送更改邮箱验证码
func SendChangeEmailCode(UserID uint, newEmail string) error {
	ctx := context.Background()
	// 生成验证码
	code := GenerateVerifyCode()
	// 设置redis key
	key := EmailRedisKey(UserID)
	// 写入redis 验证码过期时间
	err := global.RDB.Set(ctx, key, code, time.Minute*5).Err()
	if err != nil {
		return err
	}

	// 发送邮件
	return pkg.SendVerifyCode(newEmail, code)
}

// ConfirmChangeEmail 确认绑定电子邮件
func ConfirmChangeEmail(UserID uint, newEmail, code string) error {
	ctx := context.Background()
	// 设置redis key
	key := EmailRedisKey(UserID)
	// 从 Redis 取验证码
	cacheCode, err := global.RDB.Get(ctx, key).Result()
	if err == redis.Nil {
		return fmt.Errorf("验证码已过期")
	}
	if err != nil {
		return err
	}

	// 校验验证码
	if cacheCode != code {
		return fmt.Errorf("验证码错误")
	}

	// 更改邮箱
	global.DB.Model(&models.SysUser{}).Where("id = ?", UserID).Update("email", newEmail)

	// 删除验证码（只用一次, 防止重复使用）
	global.RDB.Del(ctx, key)

	return nil

}
