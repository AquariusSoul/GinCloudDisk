package global

import (
	"GinCloudDisk/conf"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

func InitDatabase(conf *conf.Config) *gorm.DB {
	var db *gorm.DB
	var level logger.LogLevel
	switch conf.Server.DbLogMode {
	case "silent":
		level = logger.Silent
	case "info":
		level = logger.Info
	case "warn":
		level = logger.Warn
	case "error":
		fallthrough
	default:
		level = logger.Error
	}
	config := &gorm.Config{
		Logger:                                   logger.Default.LogMode(level),
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
		SkipDefaultTransaction:                   true, //禁用默认事务（提高运行速度）
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //单数表名
		},
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		conf.Mysql.Username, conf.Mysql.Password, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.Dbname, conf.Mysql.Config,
	)
	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		log.Fatalf("数据库连接失败", err)
	}
	log.Println("数据库连接成功")

	// 迁移数据库
	if conf.Server.DbAutoMigrate {

		log.Println("数据库自动迁移成功")
	}
	return db
}

func InitRedis(conf *conf.Config) *redis.Client {
	var client *redis.Client
	client = redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("Redis连接失败：", err)
	}
	log.Println("Redis连接成功", conf.Redis.Addr, conf.Redis.DB, conf.Redis.Password)
	return client
}
