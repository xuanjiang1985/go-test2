package main

import "fmt"

func main() {
	f0()
	f1()
	f2()
	f3()
	f4()
}

func f0() {
	a := "abc"
	bs := []byte(a)
	fmt.Println(bs, len(bs), cap(bs))
	// 输出： [97 98 99] 3 8
}

func f1() {
	a := "abc"
	bs := []byte(a)
	fmt.Println(len(bs), cap(bs))
	// 输出: 3 32
}

func f2() {
	bs := []byte("abc")
	fmt.Println(len(bs), cap(bs))
	// 输出: 3 3
}

func f3() {
	a := ""
	bs := []byte(a)
	fmt.Println(bs, len(bs), cap(bs))
	// 输出: [] 0 0
}

func f4() {
	a := ""
	bs := []byte(a)
	fmt.Println(len(bs), cap(bs))
	// 输出: 0 32
}
