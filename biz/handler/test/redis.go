package user

import (
	"context"
	"gin_template/biz/dal/redis"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SetRedis struct {
	Key   string `json:"key" binding:"required,min=1,max=255"`
	Value string `json:"value" binding:"required,min=1,max=255"`
}

// RedisSet 测试设置redis
// @Tags 测试接口
// @Summary 测试设置redis
// @Description 测试设置redis
// @Accept application/json
// @Produce application/json
// @Param req body SetRedis true "键值对信息"
// @Router /api/test/redis/set [PUT]
func RedisSet(c *gin.Context) {
	req := new(SetRedis)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := redis.RDB.Set(context.Background(), req.Key, req.Value, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "设置键值对失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "设置成功",
	})
}

type getRedis struct {
	Key string `uri:"key" binding:"required"`
}

// RedisGet 测试获取redis
// @Tags 测试接口
// @Summary 测试获取redis
// @Description 测试获取redis
// @Accept application/json
// @Produce application/json
// @Param key path string true "key"
// @Router /api/test/redis/get/:key [GET]
func RedisGet(c *gin.Context) {
	uriReq := new(getRedis)
	if err := c.ShouldBindUri(uriReq); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	val, err := redis.RDB.Get(context.Background(), uriReq.Key).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取键值对失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"key": uriReq.Key, "value": val})
	}
}
