package main

import "fmt"

func main() {
	dict := map[string]int {
		"foo": 1,
		"bar": 2,
	}

	fmt.Println(dict)
	v, ok := dict["foo"]
	if ok {
		fmt.Println("foo exists in map. value", v)
	}

	_, ok = dict["taccisum"]
	if !ok {
		fmt.Println("taccisum does not exist in map.")
	}
}
