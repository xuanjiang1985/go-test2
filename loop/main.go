package main

import (
	"fmt"
	"strings"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)
	ChangeArr(slice1)
	fmt.Println(slice1)

	map1 := make(map[int]int, 10)
	map1[0] = 0
	map1[1] = 1
	map1[2] = 2
	map1[3] = 3
	map1[4] = 4

	fmt.Println(&map1)
	changeMap(map1)
	fmt.Println(map1)

	s5 := "hello world"
	isIn := strings.Contains(s5, "a")
	fmt.Println(isIn)
}

// ChangeArr is a func to change array
func ChangeArr(slice1 []int) {
	slice1[0] = 100
}

func changeMap(m map[int]int) {
	m[0] = 100
}
