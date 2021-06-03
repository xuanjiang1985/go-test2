package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string
	Age  int8
}

func main() {
	var num int = 100
	test01(num)

	student := student{
		Name: "zhougang",
		Age:  32,
	}
	test02(student)
}

func test01(n interface{}) {
	rType := reflect.TypeOf(n)
	fmt.Println(rType)
	fmt.Println(rType.String())
	rValue := reflect.ValueOf(n)
	fmt.Println(rValue.Kind())

	iV := rValue.Interface()
	num := iV.(int)
	fmt.Printf("%d\n", num)
}

func test02(a interface{}) {
	rType := reflect.TypeOf(a)
	fmt.Println(rType.Kind())
	rValue := reflect.ValueOf(a)

	iV := rValue.Interface()

	fmt.Printf("iV=%v, iV type=%T\n", iV, iV)
	fmt.Printf("rValue.kind=%v rType.kind=%v\n", rValue.Kind(), rType.Kind())
}
