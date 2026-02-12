package handler

import (
	"github.com/gin-gonic/gin"
	"go-web/service"
	"net/http"
)

func UpdateAvatarHandler(c *gin.Context) {
	// 从 JWT 中获取用户 ID
	userID := c.GetUint("UserID")

	//  接收上传的头像文件
	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(400, gin.H{"msg": "头像文件不能为空"})
		return
	}

	// 调用 service 处理头像逻辑
	avatarURL, err := service.UpdateUserAvatar(userID, file)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	//  返回结果
	c.JSON(200, Response{
		Code: http.StatusOK,
		Msg:  "修改成功",
		Data: map[string]string{
			"avatar": avatarURL,
		},
	})
}
