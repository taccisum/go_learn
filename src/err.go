package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero!")
	}
	return a / b, nil
}

func main() {
	fmt.Println(divide(10, 2))
	fmt.Println(divide(10, 0))
}
