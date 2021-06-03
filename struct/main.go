package main

import (
	"fmt"
	"os"
)

// Person define person struct
type Person struct {
	name string
}

func main() {
	p, err := newP()

	fmt.Println(p, err)
}

func newP() (*Person, error) {
	f, err := os.Open("aa.1")
	defer f.Close()

	if err != nil {
		return nil, err
	}

	p := &Person{"zhougang"}

	return p, nil
}
