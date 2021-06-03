package ztest

import (
	"strconv"
	"testing"
)

func TestMapLen(t *testing.T) {
	m := make(map[int]string, 8)

	for i := 0; i < 19; i++ {
		m[i] = strconv.Itoa(i)
		t.Logf("%p - %d", m, len(m))
	}
}

func TestSlice(t *testing.T) {
	s := make([]int, 0, 12)
	for i := 0; i < 13; i++ {
		s = append(s, i)
		t.Logf("%p - %d", s, cap(s))
	}
}
