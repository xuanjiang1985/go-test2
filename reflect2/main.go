package main

import (
	"fmt"
	"reflect"
)

func main() {
	student := Student{
		Name:    "zhougang",
		Age:     35,
		Address: "民治大道",
	}

	typ := reflect.TypeOf(student)
	val := reflect.ValueOf(student)
	fmt.Printf("typ=%v, kind=%v, val=%v\n", typ, val.Kind(), val)
	numField := val.NumField()
	for i := 0; i < numField; i++ {
		fmt.Printf("第%d个字段值为%v\n", i, val.Field(i))
		tag := typ.Field(i).Tag.Get("json")
		fmt.Printf("第%d个字段tag为%v\n", i, tag)
	}

	numMethod := val.NumMethod()

	fmt.Printf("拥有%d个方法\n", numMethod)
	val.Method(1).Call(nil)
	rs := make([]reflect.Value, 0, 2)
	rs = append(rs, reflect.ValueOf(10))
	rs = append(rs, reflect.ValueOf(30))
	sum := val.Method(0).Call(rs)
	fmt.Printf("count sum=%d\n", sum[0].Int())
}

// Student struct
type Student struct {
	Name    string `json:"name"`
	Age     int8   `json:"age"`
	Address string
}

// Print method of Student
func (s Student) Print() {
	fmt.Printf("Student content Name=%v, Aget=%v\n", s.Name, s.Age)
}

// Count method
func (s Student) Count(a, b int) int {
	return a + b
}
