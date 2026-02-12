package handler

import (
	"github.com/gin-gonic/gin"
	"go-web/service"
	"net/http"
	"strconv"
)

func GetUserDetailByAdminHandler(c *gin.Context) {
	// 从 URL 路径中获取用户 ID 参数
	idStr := c.Param("id")
	// 将idStr是string类型转成uint64类型
	id64, _ := strconv.ParseUint(idStr, 10, 64)
	// 获取当前用户ID
	adminID := c.GetUint("UserID")
	// 调用获取普通用户详情业务逻辑
	resp, err := service.UserDetailByAdmin(adminID, uint(id64))
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
		Msg:  "获取普通用户详情成功",
		Data: resp,
	})
}
