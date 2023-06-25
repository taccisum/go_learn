package main

import "fmt"

func foo() {
	fmt.Println("hello function.")
}

func add(a int, b int) int {
	return a + b
}

func curringAdd(a int) func(int) int {
	return func(b int) int {
		return add(a, b)
	}
}

func main() {
	foo()
	add3 := curringAdd(3)
	fmt.Println(add3(5))
}
