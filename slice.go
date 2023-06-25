package main

import "fmt"

func main() {
	ls := []int {1,2,3}
	fmt.Println(ls)
	ls = append(ls, 4)
	fmt.Println(ls)
	fmt.Println(ls[1:])
	fmt.Println(ls[:3])
}
