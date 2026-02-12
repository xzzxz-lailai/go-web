package handler

import (
	"github.com/gin-gonic/gin"
	"go-web/service"
	"net/http"
	"strconv"
)

func UpDateUserByAdminHandler(c *gin.Context) {
	//  获取路径参数 id
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(400, gin.H{"msg": "参数错误"})
		return
	}

	// id 转换
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"msg": "参数错误"})
		return
	}

	// 绑定请求体
	var req service.UpDateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"msg": "参数错误"})
		return
	}
	// 获取当前用户ID
	adminID := c.GetUint("UserID")
	// 调用 service
	if err := service.UpdateUserByAdmin(adminID, uint(id), req); err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:  "修改成功",
		Data: nil,
	})
}
