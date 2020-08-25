package main

import (
	"fmt"
)

var (
	ch   = make(chan struct{})
	done = make(chan struct{})
)

func main() {
	go work(ch)
	fmt.Println("end")
	<-ch
}

func work(ch chan struct{}) {
	fmt.Println("working")
	ch <- struct{}{}
}
