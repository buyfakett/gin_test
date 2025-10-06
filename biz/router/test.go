package router

import (
	hTest "gin_template/biz/handler/test"

	"github.com/gin-gonic/gin"
)

func testRoutes(apiGroup *gin.RouterGroup) {
	testGroup := apiGroup.Group("/test")
	{
		testGroup.GET("/redis", hTest.Redis)
	}
}
