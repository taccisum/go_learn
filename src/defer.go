package main

import "fmt"

func deferTest() {
	fmt.Println("hello defer")

	fmt.Println("lock resource 1")
	defer fmt.Println("defer unclok resource 1")

	fmt.Println("lock resource 2")
	defer func() {
		fmt.Println("defer unclok resource 2")
	}()

	fmt.Println("lock resource 3")
	panic("oooops! panic")
	defer fmt.Println("defer unclok resource 3")
}

func main() {
	deferTest()
}
