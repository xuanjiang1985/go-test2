package ztest

import "testing"

func TestHello(t *testing.T) {
	t.Log("hello world")
}

func TestB(t *testing.T) {
	t.Log("B")
}
func TestC(t *testing.T) {
	t.Log("C")
}

func Benchmark_Add(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}
