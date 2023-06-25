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

func create(i int) NameAware {
	switch i {
	case 1:
		return &Person{
			FirstName: "tac",
			LastName:  "cisum",
		}
	case 2:
		return &Cat{
			Master:   "taccisum",
			NickName: "jasmine",
		}
	}
	var def *Person = nil // 将 nil 的 Person 转换为接口 NameAware 会自动触发 boxing 操作，自动创建一个 iface 类型实例
	return def
}

func main() {
	var na NameAware

	na = create(1)
	fmt.Println("Person.Name():", na.Name())

	na = create(2)
	fmt.Println("Cat.Name():", na.Name())

	na = create(3)
	fmt.Println("Boxing example:", na, na == nil) // 由于触发了装箱，这里拿到的结果与 nil 直接比较会返回 false
}
