package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {

	src, err := os.OpenFile("./logrus.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("create file log failed: %v", err)
	}
	logger.Out = src
	logger.SetFormatter(&logrus.JSONFormatter{})
}

func main() {

	logger.Info("info msg")

	ctx, cancel := context.WithCancel(context.Background())
	ticker := time.NewTicker(30 * time.Second)
	go webServer(ctx, ticker)

	// 程序终止信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGHUP)

	// 退出信号
	s := <-signalChan
	cancel()
	logger.Info("所有进程退出:", s)
	time.Sleep(1 * time.Second)
}

func webServer(ctx context.Context, ticker *time.Ticker) {
	// ticker
	go func() {
		for {
			select {
			case <-ctx.Done():
				logger.Info("web 服务器退出")
				return
			case <-ticker.C:
				logger.Info(("运行中"))
			}
		}
	}()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	fmt.Println("请访问网站: http://127.0.0.1:8888")
	r.Run(":8888") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
