package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 100
	test01(&num)

	fmt.Printf("num=%d\n", num)
}

func test01(n interface{}) {
	rValue := reflect.ValueOf(n)
	rValue.Elem().SetInt(200)
}
