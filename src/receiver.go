package main

import "fmt"

type MyStr string

type MyWrapper struct {
	Str string
}

func (s MyStr) print() {
	fmt.Println("print", s)
}

func (s *MyStr) ptrPrint() {
	fmt.Println("pointer print", *s)
}

func (s MyWrapper) wrap() MyWrapper {
	s.Str = "[" + s.Str + "]"
	return s
}

func (s *MyWrapper) wrap1() {
	s.Str = "[" + s.Str + "]"
}

func main() {
	s := MyStr("hello receiver.")
	s.print()
	(&s).ptrPrint()

	ws := MyWrapper{Str: "wrap"}
	fmt.Println(ws.wrap().Str)    // 传递 copy
	fmt.Println((&ws).wrap().Str) // 传递 copy
	fmt.Println(ws.Str)           // 原对象值未被修改
	(&ws).wrap1()                 // 传递指针
	fmt.Println(ws.Str)           // 原对象值被修改
}
