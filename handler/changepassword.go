package handler

import (
	"github.com/gin-gonic/gin"
	"go-web/service"
	"net/http"
)

type ChangePasswordReq struct {
	OldPassword     string `json:"old_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

// ChangepasswordHandler 更改密码
func ChangepasswordHandler(c *gin.Context) {
	var req ChangePasswordReq
	// 参数绑定
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 获取当前用户ID
	userID := c.GetUint("UserID")
	// 调用注册更改密码逻辑
	err = service.ChangePassword(userID, req.OldPassword, req.NewPassword, req.ConfirmPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 返回成功响应
	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:  "修改密码成功",
		Data: nil,
	})
}
