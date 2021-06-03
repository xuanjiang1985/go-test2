package main

import "fmt"

func create() func() int {
	c := 2
	return func() int {
		return c + 1
	}
}

func main() {
	f1 := create()
	f2 := create()
	fmt.Println(f1(), f2())
}
