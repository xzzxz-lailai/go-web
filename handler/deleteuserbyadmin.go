package handler

import (
	"github.com/gin-gonic/gin"
	"go-web/service"
	"net/http"
	"strconv"
)

func DeleteUserByAdminHandler(c *gin.Context) {
	// 从 URL 路径中获取用户 ID 参数
	idStr := c.Param("id")
	// 将idStr是string类型转成uint64类型
	id64, _ := strconv.ParseUint(idStr, 10, 64)
	// 获取当前用户ID
	adminID := c.GetUint("UserID")
	// 调用注册业务逻辑
	err := service.DeleteUserByAdmin(adminID, uint(id64))
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
		Msg:  "管理员删除普通用户成功",
		Data: nil,
	})
}
