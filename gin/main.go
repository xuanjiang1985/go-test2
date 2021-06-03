package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {

	log.Out = os.Stdout

	// 为当前logrus实例设置消息输出格式为json格式。
	// 同样地，也可以单独为某个logrus实例设置日志级别和hook，这里不详细叙述。
	log.Formatter = &logrus.JSONFormatter{}
	// writer1 := &bytes.Buffer{}
	// writer2 := os.Stdout
	writer3, err := os.OpenFile("log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	log.SetOutput(io.MultiWriter(writer3))
	log.Info("info msg")

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		log.Info("ping")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	ctx, cancel := context.WithCancel(context.Background())

	signalChan := make(chan os.Signal, 1)

	s := make(chan string)
	s2 := make(chan string)
	s3 := make(chan string)

	go work(ctx, s, s2, s3)

	//r.Run(":8888") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		s3 <- "s3"
	}()

	<-s3
	cancel()
	log.Info("s3 done")
	time.Sleep(3 * time.Second)
}

func work(ctx context.Context, s, s2, s3 chan string) {
	log.Info("I am work")
	s <- "work in channel"
	go work2(ctx, s2, s3)
}

func work2(ctx context.Context, s2 chan string, s3 chan string) {

	for {
		select {
		case <-ctx.Done():
			log.Info("ctx done")
			s3 <- "done"
			return
		case a := <-s2:
			fmt.Println(a)
		}
	}
}

func exit(signalChan chan os.Signal) {
	<-signalChan
}
