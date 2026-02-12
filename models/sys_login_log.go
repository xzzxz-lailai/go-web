package models

import "time"

type SysLoginLog struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"` // 日志ID
	UserID     uint      `gorm:"column:user_id;not null"`  // 用户ID
	Username   string    `gorm:"column:username;size:64;not null" `
	IP         string    `gorm:"column:ip;size:64;not null" ` // 登录IP
	UserAgent  string    `gorm:"column:user_agent;size:255" ` // 浏览器/设备信息
	Status     int       `gorm:"column:status;not null" `     // 登录状态：1成功 0失败
	FailReason string    `gorm:"column:fail_reason;size:64"`  // 失败原因（密码错误/用户不存在等）
	City       string    `gorm:"column:city;size:64" `        // 登录地区（省/国家）
	LoginAt    time.Time `gorm:"column:login_at;not null" `   // 登录时间
}

// TableName 指定表名
func (SysLoginLog) TableName() string {
	return "sys_login_log"
}
