package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int64
	b := &a
	fmt.Println(b)
	var f float64
	s := "hello world zhou zhoug"
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("类型%T, 大小 %dbyte\n", a, unsafe.Sizeof(a))
	fmt.Printf("类型%T, 大小 %dbyte\n", b, unsafe.Sizeof(b))
	fmt.Printf("类型%T, 大小 %dbyte\n", f, unsafe.Sizeof(f))
	fmt.Printf("类型%T, 大小 %dbyte\n", s, unsafe.Sizeof(s))
	fmt.Printf("类型%T, 大小 %dbyte\n", slice, unsafe.Sizeof(slice))
	ma := make(map[int]string, 5)
	fmt.Printf("类型%T, 大小 %dbyte\n", ma, unsafe.Sizeof(ma))

	// err1 := errors.New("这是一个错误")
	// fmt.Println(err1.Error())

	err2 := fmt.Errorf("cuowu: %d", 100)
	fmt.Println(err2)
	fmt.Println([]byte("周刚"))
	sb := []byte{229, 145, 168, 98}
	fmt.Println(string(sb))
}
