package main

import "fmt"

func main() {
	fmt.Println([3]int {1,2,3})
	for _, v := range [5]int {4,5,6} {
		fmt.Println(v)
	}
}
