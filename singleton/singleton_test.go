package singleton

import (
	"fmt"
	"testing"
)

func TestGetInstance(t *testing.T) {
	p := GetInstance()
	p.SetName("zhougang")
	fmt.Println(p.GetName())
}
