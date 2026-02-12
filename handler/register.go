package handler

import (
	"github.com/gin-gonic/gin"
	"go-web/service"
	"net/http"
)

func RegisterHandler(c *gin.Context) {
	var req service.RegisterRequest
	// 参数绑定
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code: http.StatusBadRequest,
			Msg:  "参数错误",
		})
		return
	}
	// 调用注册业务逻辑
	resp, err := service.Register(req)
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
		Msg:  "注册成功",
		Data: resp,
	})
}
