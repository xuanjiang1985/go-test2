package ztest

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var s string
var number = 1000

func BenchmarkStringPlus(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < number; j++ {
			s += strconv.Itoa(j)
		}

	}
}

func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	var builder strings.Builder
	for i := 0; i < b.N; i++ {
		for j := 0; j < number; j++ {
			builder.WriteString(strconv.Itoa(j))
		}

	}
	_ = builder.String()
}

func BenchmarkStringPrint(b *testing.B) {
	var builder strings.Builder
	for i := 0; i < b.N; i++ {
		fmt.Fprint(&builder, "?")
		_ = builder.String()
	}

}

// 最优
func BenchmarkStringWrite(b *testing.B) {
	var builder strings.Builder
	for i := 0; i < b.N; i++ {
		builder.WriteString("1")
		_ = builder.String()
	}

}

func BenchmarkStringByte(b *testing.B) {
	var builder strings.Builder
	for i := 0; i < b.N; i++ {
		builder.Write([]byte("1"))
		_ = builder.String()
	}

}
