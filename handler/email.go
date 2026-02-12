package handler

import (
	"github.com/gin-gonic/gin"
	"go-web/service"
	"net/http"
)

// 发送邮件验证码
func SendEmailCodeHandler(c *gin.Context) {
	userID := c.GetUint("UserID")

	var req struct {
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
		return
	}

	if err := service.SendChangeEmailCode(userID, req.Email); err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(200, Response{
		Code: http.StatusOK,
		Msg:  "验证码已发送",
		Data: nil,
	})
}

// 验证邮箱验证码
func ConfirmEmailHandler(c *gin.Context) {
	userID := c.GetUint("UserID")

	var req struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"msg": "参数错误"})
		return
	}

	if err := service.ConfirmChangeEmail(userID, req.Email, req.Code); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(200, Response{
		Code: http.StatusOK,
		Msg:  "邮箱更换成功",
		Data: nil,
	})
}
