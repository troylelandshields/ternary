package ternary

import "fmt"

type Nameable interface {
	Name() string
}

type Cat struct {
	name string
}

func (c Cat) Name() string {
	return c.name
}

type Dog struct {
	name string
}

func (d Dog) Name() string {
	return d.name
}

func ExampleGive_second() {
	var d *Dog
	c := &Cat{
		name: "Roger",
	}

	fmt.Println(Give[Nameable](d).If(d != nil).Else(c).Name())
	// Output: Roger
}
