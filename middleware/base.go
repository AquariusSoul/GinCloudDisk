package middleware

import (
	g "GinCloudDisk/utils/global"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 24 * time.Hour,
	})
}

func WithRedis(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(g.CTX_RDB, rdb)
		c.Next()
	}
}

func WithGormDB(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(g.CTX_DB, db)
		c.Next()
	}
}
