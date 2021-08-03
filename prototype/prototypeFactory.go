package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Suite                int
	StreetAddreess, City string
}

type Employee struct {
	Name   string
	Office Address
}

func (em *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(em)

	d := gob.NewDecoder(&b)
	result := Employee{}
	_ = d.Decode(&result)

	return &result
}

var mainOffice = Employee{
	"",
	Address{0, "123 East", "Taipei"},
}

var auxOffice = Employee{
	"",
	Address{0, "666 West", "Taipei"},
}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite

	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}

func main() {
	john := NewMainOfficeEmployee("John", 100)
	jane := NewMainOfficeEmployee("Jane", 200)

	fmt.Println(john)
	fmt.Println(jane)
}
