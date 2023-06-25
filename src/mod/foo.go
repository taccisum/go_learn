package main

import "fmt"

func init() {
	fmt.Println("foo.go init has been invoked.")
}

func Sayhello() {
	fmt.Println("hello go mod.")
}
