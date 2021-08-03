package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "develoer", 1000}
	case Manager:
		return &Employee{"", "manager", 2000}
	default:
		panic("Unsupported roles")
	}
}

func main() {
	m := NewEmployee(Manager)
	m.Name = "yap"
	fmt.Println(m)
}
