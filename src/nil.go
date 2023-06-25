package main

import "fmt"

func main() {
	var a int
	var b bool
	var s string
	var slice []int
	var arr [1]int
	fmt.Println(a == 0)
	fmt.Println(b == false)
	fmt.Println(s == "")
	fmt.Println(slice == nil)
	fmt.Println(len(arr) == 1)
	fmt.Println(arr[0] == 0)
}
