package ctxgo

import (
	"context"
	"fmt"
	"time"
)

func main() {
	s := []int{23, 32, 78, 43, 76, 65, 345, 762, 1, 2, 3, 4915, 86, 10}
	signal := make(chan struct{})

	timer := time.NewTimer(time.Second * 3)
	ctx, cancel := context.WithCancel(context.Background())

	go search(ctx, s[:len(s)/2], 86, signal)
	go search(ctx, s[len(s)/2:], 86, signal)

	select {
	case <-timer.C:
		fmt.Println("timeout")
		cancel()
	case <-signal:
		cancel()
	}
	time.Sleep(time.Second * 4)
}

func search(ctx context.Context, s []int, a int, c chan struct{}) {
	for _, v := range s {

		select {
		case <-ctx.Done():
			fmt.Println("task canceled")
			return
		default:
		}
		time.Sleep(time.Millisecond * 100)
		if a == v {
			fmt.Println("found it")
			c <- struct{}{}
			return
		}
	}
}
