package router

import (
	hTest "gin_template/biz/handler/test"

	"github.com/gin-gonic/gin"
)

func testRoutes(r *gin.Engine) {
	userGroup := r.Group("/api")
	{
		userGroup.GET("/test/redis", hTest.Redis)
	}
}
