package handler

import (
	"github.com/gin-gonic/gin"
	"go-web/service"
	"net/http"
)

// GetWeatherHandler 查询天气
func GetWeatherHandler(c *gin.Context) {
	city := c.Query("city")

	weather, err := service.GetWeather(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code: http.StatusOK,
			Msg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:  "success",
		Data: weather,
	})
}
