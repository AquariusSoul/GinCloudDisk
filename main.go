package main

import (
	"GinCloudDisk/conf"
	"GinCloudDisk/log"
	"GinCloudDisk/router"
	"GinCloudDisk/utils"
	g "GinCloudDisk/utils/global"
	"net/http"
)

func main() {
	// 读取配置文件
	conf := conf.ReadConfig("../conf.yaml")

	// 初始化基础资源
	sLog := log.InitLogger(conf)
	defer sLog.Sync()
	db := g.InitDatabase(conf)
	rdb := g.InitRedis(conf)

	// 初始化路由
	r := router.InitRouter(db, rdb)

	// TODO:根据gin的开发模式，调整 默认 or 自定义恢复中间件/日志中间件

	srv := &http.Server{
		Addr:    ":" + conf.Server.Port,
		Handler: r,
	}

	// 启动一个协程，监听端口
	go func() {
		sLog.Info("Gin-cloud-disk server start")
		// connect serve
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			sLog.Fatalf("listen: %s\n", err)
		}
	}()

	// 关闭服务程序
	utils.CloseServer(srv)
}
