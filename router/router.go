package router

import (
	"GinCloudDisk/controller"
	"GinCloudDisk/middleware"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	UserAPI controller.User
)

// InitRouter 初始化路由
func InitRouter(db *gorm.DB, rdb *redis.Client) *gin.Engine {
	r := gin.Default()
	// 添加中间件
	r.Use(middleware.CORS())
	r.Use(middleware.WithGormDB(db))
	r.Use(middleware.WithRedis(rdb))

	// 测试连接
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 用户模块
	r.GET("/login", UserAPI.User_Login)
	r.POST("/registry", UserAPI.User_Register)
	user := r.Group("/user")
	{
		user.GET("/list") //获取共享文件列表
	}
	return r
}
