package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func main() {
	file, err := os.OpenFile("config.yaml", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("%+v", errors.WithStack(err))
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		writer.WriteString("hello zhougang 1\n")
	}
	//io.Copy()
	writer.Flush()
}
