package main

import (
	"fmt"
	"test/logger"
	"time"

	"github.com/pkg/errors"
)

var logger1 = logger.Intance()
func main(){
	go say()
	go done()

	time.Sleep(10 * time.Second)
}

func say() {
	for i := 0; i < 10; i++ {
		fmt.Printf("hello\n")
		time.Sleep(time.Second)
	}
}

func done() {
	defer func() {
		if err := recover(); err != nil {
			logger1.Errorf("%+v", errors.WithStack(err.(error)))
		}
	}()

	time.Sleep(3*time.Second)
	var intMap map[int]string
	intMap[0] = "hi"
}