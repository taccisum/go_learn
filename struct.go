package main

import "fmt"

type Foo struct {
	Int int
	Str string
	Bar Bar
}

type Bar struct {
	Bool bool
}

func main() {
	foo := Foo {
		Int: 1,
		Str: "hello",
		Bar: Bar {
			Bool: true,
		},
	}
	fmt.Println(foo)
	fmt.Println(foo.Bar.Bool)
}
