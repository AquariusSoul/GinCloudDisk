package utils

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// CloseServer 优雅地关闭服务
func CloseServer(srv *http.Server) {
	// 优雅地通过os信号关闭程序
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Gin-cloud-disk Server Shutdown ...")

	// 设置超时时间，服务关闭后，最多5秒就关闭HTTP连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		log.Println("HTTP Server Shutdown ...")
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP Server Shutdown Failed:%+v", err)
	}
}
