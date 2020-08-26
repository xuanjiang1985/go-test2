package main

import (
	"test/logger"
)

var (
	intCh    = make(chan int, 10)
	signalCh = make(chan struct{})
)

func main() {
	go write(intCh)
	go read(intCh, signalCh)
	<-signalCh
	//time.Sleep(time.Second * 3)
}

func write(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch)
}

func read(ch <-chan int, signal chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		logger.Intance().Errorf("get data from channel: %d", v)
	}

	signal <- struct{}{}
}
