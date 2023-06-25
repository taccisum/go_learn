package main

import "fmt"

type Foo struct {
	Int int
	Str string
	Bar // 类型嵌入
}

type Bar struct {
	Bool bool
}

func main() {
	foo := Foo{
		Int: 1,
		Str: "hello",
		Bar: Bar{
			Bool: true,
		},
	}
	fmt.Println("foo:", foo)
	fmt.Println("foo.Bar.Bool:", foo.Bar.Bool)
	fmt.Println("foo.Bool:", foo.Bool)
}
