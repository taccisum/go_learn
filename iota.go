package main

import "fmt"

const (
	a = 1 << iota
	b
	c
)

func main() {
	fmt.Println(a, b, c)
}
