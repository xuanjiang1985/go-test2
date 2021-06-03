package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/pkg/errors"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")

	if err != nil {
		fmt.Printf("%+v\n", errors.WithStack(err))
	}

	defer conn.Close()
	for {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("%+v\n", errors.WithStack(err))
		}

		if line == "exit\n" {
			fmt.Println("exit")
			break
		}

		n, err := conn.Write([]byte(line))

		if err != nil {
			fmt.Printf("%+v\n", errors.WithStack(err))
		}

		fmt.Printf("client sent %d bytes\n", n)
	}

}
