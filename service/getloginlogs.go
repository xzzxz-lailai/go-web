package service

import (
	"go-web/global"
	"go-web/models"
	"time"
)

type GetLoginLogsRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"size"`
}

type LoginLogResponse struct {
	Username   string
	IP         string
	Status     string
	FailReason string
	UserAgent  string
	LoginAt    time.Time
}

func GetLoginLogs(req GetLoginLogsRequest) ([]LoginLogResponse, int64, error) {
	var (
		list  []LoginLogResponse // 返回记录登陆日志参数
		total int64              // 总数的个数
	)
	// 防止 非法分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 统计总数
	err := global.DB.Model(&models.SysLoginLog{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (req.Page - 1) * req.PageSize
	// 分页查询   // 按登录时间倒序
	err = global.DB.Model(&models.SysLoginLog{}).Order("login_at desc").Offset(offset).Limit(req.PageSize).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	// 返回响应
	return list, total, nil
}
