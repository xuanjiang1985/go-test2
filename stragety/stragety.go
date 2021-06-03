package stragety

import "fmt"

type Stragety interface {
	Excute()
}

type stragetyA struct {
}

func NewStragetyA() Stragety {
	return &stragetyA{}
}

func (s *stragetyA) Excute() {
	fmt.Println("A plan excuted")
}

type stragetyB struct {
}

func NewStragetyB() Stragety {
	return &stragetyB{}
}

func (s *stragetyB) Excute() {
	fmt.Println("B plan excuted")
}

type Content struct {
	stragety Stragety
}

func NewContent() *Content {
	return &Content{}
}

func (c *Content) SetStragety(stragety Stragety) {
	c.stragety = stragety
}

func (c *Content) Excute() {
	c.stragety.Excute()
}
