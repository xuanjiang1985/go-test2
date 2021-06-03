package ztest

import (
	"testing"
)

func TestSliceValue(t *testing.T) {
	var s []int
	t.Log(s, len(s), cap(s))
	if s == nil {
		t.Log("nil!")
	}
}
