package unittest

import (
	"testing"
)

func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 45 {
		t.Fatalf("addUpper(10) err 实际值=%v，期望值=%v\n", res, 55)
	}

	t.Logf("addUpper(10) ---- success")
}

func TestGetSub(t *testing.T) {
	res := getSub(10, 3)
	if res != 7 {
		t.Fatalf("getSub(10, 3) err 实际值=%v，期望值=%v\n", res, 7)
	}

	t.Logf("getSub(10, 3) ---- success")
}
