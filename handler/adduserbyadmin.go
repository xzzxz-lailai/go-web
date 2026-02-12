package handler

import (
	"github.com/gin-gonic/gin"
	"go-web/service"
	"net/http"
)

func AddUserByAdminHandler(c *gin.Context) {
	var req service.AddUserRequest
	// 参数绑定
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 获取当前用户ID
	adminID := c.GetUint("UserID")
	// 调用注册业务逻辑
	err = service.AddUserByAdmin(adminID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 返回成功响应
	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:  "管理员添加普通用户成功",
		Data: nil,
	})
}
