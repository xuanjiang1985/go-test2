package ztest

import (
	"testing"
)

type User struct {
	name string
	age  int
}

func test2(m map[int]User) {
	for i := 0; i < 10000; i++ {
		user := User{
			name: "小明",
			age:  i,
		}
		m[i] = user
	}
}

func testPointer(m map[int]*User) {
	for i := 0; i < 10000; i++ {
		user := User{
			name: "小明",
			age:  i,
		}
		m[i] = &user
	}
}

func BenchmarkStruct(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[int]User, 1000)
		test2(m)
	}
}

func BenchmarkStructPointer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[int]*User, 10000)
		testPointer(m)
	}
}
