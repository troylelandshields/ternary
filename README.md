# Ternary

> 
> Don't write 4 lines of code when 1 will do.
> 
> -- Not a Go proverb

This completely useless package uses Go's new generics to add a ternary operator style syntax to Go! You know there have been at least a handful of times where you've quit your working day early because you just could not write one more `if` statement. Well now you don't have to, feel free to replace all of your `if` statements with this the beautiful and not-confusing-at-all syntax provided by `ternary`.

_Because this requires Generics, you must be using go 1.18._

## Installation

`go get github.com/troylelandshields/ternary`

## Usage

### Simple Example
```go
var x = -9
absoluteValue := Give(x).If(x > 0).Else(-x)
fmt.Println(absoluteValue)
```

### Interface Example
```go
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

func main() {
	var d *Dog
	c := &Cat{
		name: "Roger",
	}

	fmt.Println(Give[Nameable](d).If(d != nil).Else(c).Name())
}
```
