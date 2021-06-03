package stragety

import "testing"

func TestContext_Excute(t *testing.T) {
	stragetyA := NewStragetyA()
	content := NewContent()
	content.SetStragety(stragetyA)
	content.Excute()

	stragetyB := NewStragetyB()
	content.SetStragety(stragetyB)
	content.Excute()
}
