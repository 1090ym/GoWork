package gonote

import (
	"fmt"
	"testing"
)

type Person interface {
	GetName() string
}

type Student struct {
	name string
}

type Teacher struct {
	name string
}

func (t *Teacher) GetName() string {
	return t.name
}

func (s *Student) GetName() string {
	return s.name
}

func TestInterface(t *testing.T) {
	var s Person = &Student{"Alice"}
	var tea Person = &Teacher{"Bob"}

	x, ok := s.(*Student)
	if ok {
		PrintName(x)
	}
	//PrintName(x)
	PrintName(s)
	PrintName(tea)
}

func PrintName(p Person) {
	fmt.Println(p.GetName())
}
