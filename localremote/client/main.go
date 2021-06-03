package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
)

var (
	signalChan        = make(chan os.Signal)
	serverCommandChan = make(chan string)
	httpHost          = "http://127.0.0.1:8080"
	localUUID         = "17edc387-8f1c-fa1b-cf40-169d6f57ca9b"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// worker
	go process(ctx)

	// waiting exit signal
	go exit(signalChan)
	<-signalChan
	cancel()
	time.Sleep(100 * time.Millisecond)
}

func process(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf(" a goroutine worker exited\n")
			return
		case <-time.After(time.Second):
			fmt.Printf("I am alive %d \n", time.Now().Unix())
			execCommand()
		}
	}
}

func execCommand() {
	resp, err := http.Get(httpHost + "/gin/report/" + localUUID)

	if err != nil {
		fmt.Printf("%s done\n", errors.WithStack(err))
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s done\n", errors.WithStack(err))
		return
	}

	fmt.Printf("get server data: %s \n", b)
}

func exit(signalChan chan os.Signal) {
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
}
