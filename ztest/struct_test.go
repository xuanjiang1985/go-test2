package ztest

import (
	"testing"
)

const (
	_       = iota
	testvar // testvar 将是 int 类型
)

func TestStruct(t *testing.T) {
	var a, b struct{}
	t.Log(&a, &b)
	t.Logf("%p %p", &a, &b)
	t.Log(&a == &b)
	t.Logf("%T", testvar)
}

func testDefer(t *testing.T) int {
	var i int
	defer func() {
		i++
		t.Log("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		t.Log("defer1:", i) // 打印结果为 defer: 1
	}()
	return i
}

func TestDefer2(t *testing.T) {
	t.Log(testDefer(t))
}
