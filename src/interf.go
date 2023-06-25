package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
}

type Cat struct {
	Master   string
	NickName string
}

func (p *Person) Name() string {
	return p.FirstName + " " + p.LastName
}

func (c *Cat) Name() string {
	return fmt.Sprintf("%s's %s", c.Master, c.NickName)
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
	na = &Cat{
		Master:   "taccisum",
		NickName: "jasmine",
	}
	fmt.Println(na.Name())
}
