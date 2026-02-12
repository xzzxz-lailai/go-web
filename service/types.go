package service

// LoginRequest 登陆请求结构体
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登陆响应结构体
type LoginResponse struct {
	Token string `json:"token"`
}

// RegisterRequest 注册请求结构体
type RegisterRequest struct {
	Username string `json:"username" binding:"required"` // 用户名必填
	Password string `json:"password" binding:"required"` // 密码必填
	Phone    string `json:"phone"`                       // 手机号可选
	Avatar   string `json:"avatar"`                      // 头像 URL 可选
}

// RegisterResponse 注册响应结构体
type RegisterResponse struct {
	UserID uint `json:"user_id"`
}

// AddUserRequest 接受添加管理员结构体

// GetUserDetail获取管理员信息结构体
type GetUserDetailRequest struct {
	UserID uint `json:"user_id"`
}
