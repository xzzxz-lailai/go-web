package handler

import (
	"github.com/gin-gonic/gin"
	"go-web/service"
	"net/http"
)

func LoginHandler(c *gin.Context) {
	var req service.LoginRequest

	// 记录登陆日志操作,需要的
	ip := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	// 参数绑定
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code: http.StatusBadRequest,
			Msg:  "参数错误",
		})
		return
	}
	// 调用登陆业务逻辑
	resp, err := service.Login(req, ip, userAgent)
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
		Msg:  "登陆成功",
		Data: resp,
	})
}
