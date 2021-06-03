package main

import "fmt"

type user struct {
	ID   int
	Name string
}

func main() {
	s := "ab"
	fmt.Printf("%s", s)
	user := newUser()
	fmt.Println(user)
}

func newUser() user {
	return user{
		ID:   1,
		Name: "zhougang",
	}
}
