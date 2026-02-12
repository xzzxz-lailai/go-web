package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web/utils"
	"net/http"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从header获取token
		token := c.GetHeader("authorization")

		// 判断token是否为空
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "缺少 Authorization Token"})
			c.Abort()
			return
		}

		// 解析token
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// 存到 context，后续 Handler 可直接用
		c.Set("UserID", claims.UserID)

		c.Next()
	}
}
