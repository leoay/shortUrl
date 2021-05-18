package router

import (
	"ShortUrl/internal/service"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	route := gin.Default()
	route.POST("/shorturl", service.Long2Short)
	route.Run("0.0.0.0:8087")
}
