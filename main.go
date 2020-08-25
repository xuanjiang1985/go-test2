package main

import (
	"fmt"
	"os"
	"test/message"
	"test/monster"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	// 日志
	log.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to log to file, using default stderr")
	}
	defer file.Close()
	log.Out = file

	message1 := message.New(1, 123, "message1", "cache1")
	fmt.Printf("%+v\n", message1.Id)
	message2 := message.NewByOption(message.WithID(123))
	fmt.Println((message2.Id))

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1
	fmt.Println(s1, s2)
	s2[0] = 100

	fmt.Println(s1, s2)

	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)
	fmt.Printf("%p\n", &s1)
	fmt.Printf("%p\n", &s2)

	fmt.Println("-------------------")
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := arr1
	fmt.Println(arr1, arr2)
	arr1[0] = 100
	fmt.Println(arr1, arr2)

	monster := monster.Monster{
		Name:  "zhougang",
		Age:   18,
		Skill: "cs",
	}
	err = monster.Store()
	if err != nil {
		log.Errorf("%+v\n", err)
	}

	err = monster.Restore()

	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	fmt.Println("done")
}
