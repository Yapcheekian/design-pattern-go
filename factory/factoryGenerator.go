package main

import "fmt"

type Employee2 struct {
	Name, Position string
	AnnualIncome   int
}

// functional
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee2 {
	return func(name string) *Employee2 {
		return &Employee2{name, position, annualIncome}
	}
}

// structural
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func (f *EmployeeFactory) Create(name string) *Employee2 {
	return &Employee2{name, f.Position, f.AnnualIncome}
}

func ___main() {
	developerFactory := NewEmployeeFactory("developer", 1000)
	managerFactory := NewEmployeeFactory("manager", 2000)

	developer := developerFactory("yap")
	manager := managerFactory("yap")

	fmt.Println(developer)
	fmt.Println(manager)

	salesFactory2 := NewEmployeeFactory2("sales", 1000)
	sales := salesFactory2.Create("yap")

	fmt.Println(sales)
}
