package iface

import "fmt"

type A interface {
	Age() int
	B
}

type B interface {
	Name() string
}

type Person struct {
	age  int
	name string
}

func (p *Person) Age() int {
	return p.age
}

func (p *Person) Name() string {
	return p.name
}

func NewPerson(name string, age int) *Person {
	return &Person{
		age:  age,
		name: name,
	}
}

func RunPerson() {
	p := NewPerson("卫雄伟", 28)
	var a A
	a = p
	fmt.Println(a.Name(), a.Age())

	var b B
	b = p
	fmt.Println(b.Name())
}
