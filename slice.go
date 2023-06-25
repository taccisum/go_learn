package main

import "fmt"

func main() {
	ls := []int {1,2,3}
	fmt.Println("slice", ls)
	ls = append(ls, 4)
	fmt.Println("append 4", ls)
	fmt.Println("slice[1:]", ls[1:])
	fmt.Println("slice[:3]", ls[:3])
	fmt.Println("slice len", len(ls), "cap", cap(ls))
}
