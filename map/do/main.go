package main

import (
	"fmt"
	"test/logger"
)

func main() {
	map1 := make(map[int]int, 5)
	map1[2] = 2
	map1[10] = 10
	map2 := map1
	changeMap(map2)
	fmt.Printf("map1: %v\n", map1)
	var logger = logger.Intance()
	logger.Info("hello 后杠 is error")
}

func changeMap(map1 map[int]int) {
	map1[10] = 100
}
