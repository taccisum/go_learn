package main

import "fmt"

var (
	a = 23
	b = 123
)

const (
	Greeting = "vim-go!"
)

func main() {
	c := 10
	fmt.Println(Greeting)
	fmt.Println(a*b*c)

	for i, v := range "123" {
		fmt.Println(i, string(v))
	}
}
