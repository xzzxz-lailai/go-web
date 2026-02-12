package handler

import (
	"github.com/gin-gonic/gin"
	"go-web/service"
	"net/http"
)

func GetUserListByAdminHandler(c *gin.Context) {
	var req service.GetuserListRequest
	// 参数绑定
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code: http.StatusBadRequest,
			Msg:  "参数错误",
		})
		return
	}
	// 获取当前用户ID
	adminID := c.GetUint("UserID")
	// 调用 Service
	list, total, err := service.GetUserListByAdmin(adminID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:  "加载数据成功",
		Data: gin.H{
			"list":  list,
			"total": total,
		},
	})
}
