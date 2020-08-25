package monster

import (
	"testing"
)

func TestStore(t *testing.T) {
	monster := &Monster{
		Name:  "wang",
		Age:   19,
		Skill: "da da",
	}

	err := monster.Store()

	if err != nil {
		t.Fatalf("monster.Store test err %+v", err)
	}

	t.Logf("monster.Store() test ok")
}

func TestRestore(t *testing.T) {
	monster := &Monster{}

	err := monster.Restore()

	if err != nil {
		t.Fatalf("monster.Restore() test err %+v", err)
	}

	if monster.Name != "wang" {
		t.Fatalf("monster.Restore() err %v, expect %v", monster.Name, "wang")
	}
	t.Logf("monster.Restore() test ok")
}
