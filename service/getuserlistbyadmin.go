package service

import (
	"go-web/global"
	"time"
)

// 接收前端传来的查询参数
type GetuserListRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"size"`
}

// 给前端返回用的管理员列表结构体
type GetuserListResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

// 管理员获取普通用户列表
func GetUserListByAdmin(adminID uint, req GetuserListRequest) ([]GetuserListResponse, int64, error) {
	var (
		list  []GetuserListResponse // 用户列表
		total int64                 // 总数的个数
	)
	// 防止 非法分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	offset := (req.Page - 1) * req.PageSize

	// 构造查询(确认选择查询的表 和 字段)
	tx := global.DB.Table("sys_user AS u").
		Select("u.id, u.username, u.role, u.created_at").
		Joins("join sys_user_relation r on u.id = r.user_id").
		Where("r.admin_id = ?", adminID)

	// 查总数的个数
	err := tx.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 查当前页的数据
	err = tx.Offset(offset).Limit(req.PageSize).Scan(&list).Error
	if err != nil {
		return nil, 0, err
	}

	// 返回当前页数据 和 数据总数的个数
	return list, total, nil
}
