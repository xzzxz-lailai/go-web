package utils

import (
	"go-web/global"
	"go-web/models"
	"go.uber.org/zap"
	"time"
)

// RecordLoginLog 登录日志记录函数
func RecordLoginLog(UserID uint, username string, ip string, userAgent string, status int, failReason string) {
	log := models.SysLoginLog{
		UserID:     UserID,
		Username:   username,
		IP:         ip,
		UserAgent:  userAgent,
		Status:     status,
		FailReason: failReason,
		// City:     "", // 先不解析地区
		LoginAt: time.Now(),
	}

	err := global.DB.Create(&log).Error
	if err != nil {
		zap.L().Warn("记录登录日志失败", zap.Error(err))
	}

}
