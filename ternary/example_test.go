package ternary

import "fmt"

const (
	x = -9
)

func ExampleGive() {
	absoluteValue := Give(x).If(x > 0).Else(-x)
	fmt.Println(absoluteValue)
	// Output: 9
}
