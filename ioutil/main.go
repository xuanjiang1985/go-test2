package main

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
)

func main() {
	byte1, err := ioutil.ReadFile("config.yaml1")
	if err != nil {
		fmt.Printf("%+v", errors.WithStack(err))
	}

	ioutil.WriteFile("aa.yaml", []byte("hello"), 0666)
	fmt.Printf("%s\n", byte1)
	fmt.Printf("%.2f\n", 1.23456789)
	fmt.Printf("%02d\n", 1)
}
