package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
}

func (p *Person) Name() string {
	return p.FirstName + " " + p.LastName
}

type NameAware interface {
	Name() string
}

func main() {
	var na NameAware
	na = &Person{
		FirstName: "tac",
		LastName:  "cisum",
	}

	fmt.Println(na.Name())
}
