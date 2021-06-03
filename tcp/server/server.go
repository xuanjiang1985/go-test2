package main

import (
	"fmt"
	"io"
	"net"

	"github.com/pkg/errors"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8888")

	if err != nil {
		fmt.Printf("%+v\n", errors.WithStack(err))
	}

	defer listener.Close()

	fmt.Printf("server listening on 8888\n")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Printf("%+v\n", errors.WithStack(err))
		}

		fmt.Printf("a client connected %v, ip=%v\n", conn, conn.RemoteAddr().String())
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		b := make([]byte, 1024)
		n, err := conn.Read(b)

		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("%+v\n", errors.WithStack(err))
			return
		}

		fmt.Printf("%v", string(b[:n]))
	}

}
