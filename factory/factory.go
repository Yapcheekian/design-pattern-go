package main

import "fmt"

type Person2 struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson2(name string, age int) *Person2 {
	return &Person2{name, age, 2}
}

func _main() {
	p := NewPerson2("Yap", 25)
	fmt.Println(p)
}
