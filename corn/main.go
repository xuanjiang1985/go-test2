package main

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {

	log.Println("Starting...")

	c := cron.New(cron.WithSeconds(), cron.WithLocation(time.Local)) // 新建一个定时任务对象
	c.AddFunc("* * * * * *", func() {
		log.Println("hello world")
	}) // 给对象增加定时任务
	c.Start()
	select {}
}
