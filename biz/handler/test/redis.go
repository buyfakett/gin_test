package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Redis 测试redis
// @Tags 测试接口
// @Summary 测试redis
// @Description 测试redis
// @Accept application/json
// @Produce application/json
// @Router /api/test/redis [get]
func Redis(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "pong",
	})
}
