package ztest

import (
	"testing"
)

func test(m map[int]int) {
	for i := 0; i < 7; i++ {
		m[i] = i
	}
}

func BenchmarkMap(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[int]int) //不带容量的初始化
		test(m)
	}
}

func BenchmarkMapCap(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[int]int, 16) //带容量的初始化
		test(m)
	}
}
