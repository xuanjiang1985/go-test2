package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	read2("config.yaml")
}

//使用bufio
func read2(filename string) int {
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	buf := make([]byte, 4096)
	var nbytes int
	rd := bufio.NewReader(fi)
	for {
		n, err := rd.Read(buf)

		//s, err := rd.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		fmt.Printf("%v\n", string(buf))

		nbytes += n
	}
	return nbytes
}
