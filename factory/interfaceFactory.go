package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hello %s", p.name)
}

func NewPerson(name string, age int) Person {
	return &person{name, age}
}

func main() {
	p := NewPerson("Yap", 25)
	p.SayHello()
}
