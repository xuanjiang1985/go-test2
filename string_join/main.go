package main

import (
	"bytes"
	"fmt"
)

var sBuffer bytes.Buffer

func main() {
	sBuffer.WriteString("hello")
	sBuffer.WriteString(", ")
	sBuffer.WriteString("world")

	fmt.Printf("%s\n", sBuffer.String())
}
