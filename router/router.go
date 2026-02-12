package router

import (
	"github.com/gin-gonic/gin"
	"go-web/handler"
	"go-web/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors()) // 全局跨域

	// 静态资源路由
	r.Static("/static", "./uploads")

	// 无需鉴权的接口
	user := r.Group("/api")
	{
		user.POST("/register", handler.RegisterHandler) // 注册
		user.POST("/login", handler.LoginHandler)       // 登陆
	}

	// 鉴权接口
	auth := r.Group("/api")
	auth.Use(middleware.Authorization())
	{
		auth.POST("/user/list", handler.GetUserListByAdminHandler)   //  管理员获取普通用户列表(分页查询)
		auth.GET("/user/:id", handler.GetUserDetailByAdminHandler)   // 管理员获取普通详情信息
		auth.PUT("/user/:id", handler.UpDateUserByAdminHandler)      // 管理员修改普通用户信息
		auth.POST("/user/add", handler.AddUserByAdminHandler)        // 管理员添加普通用户
		auth.DELETE("/user/:id", handler.DeleteUserByAdminHandler)   // 管理员删除普通用户
		auth.PUT("/user/avatar", handler.UpdateAvatarHandler)        // 上传/修改头像
		auth.PUT("/user/password", handler.ChangepasswordHandler)    // 修改密码
		auth.POST("/email/code", handler.SendEmailCodeHandler)       // 发送邮件验证码(用于绑定邮箱,不是更换邮箱)
		auth.POST("/email/confirm", handler.ConfirmEmailHandler)     // 验证邮箱验证码(用于绑定邮箱,不是更换邮箱)
		auth.POST("/user/loginlog/list", handler.GetLoginLogHandler) //  管理员获取登录日志记录列表(分页查询)
	}
	return r
}
