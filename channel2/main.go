package main

import (
	"fmt"
	"strconv"
	"time"
)

var (
	intCh    = make(chan int, 10)
	stringCh = make(chan string, 10)
)

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			intCh <- i + 1
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			stringCh <- "hello " + strconv.Itoa(i+1)
		}

	}()

	for {
		select {
		case i := <-intCh:
			fmt.Printf("int channel value= %d\n", i)
		case s := <-stringCh:
			fmt.Printf("string channel value= %s\n", s)
		case <-time.After(100 * time.Millisecond):
			//fmt.Printf("done\n")
			//return
			// default:
			// 	fmt.Printf("done\n")
			// 	//return
			// 	time.Sleep(time.Millisecond * 1000)
		}
	}
}
