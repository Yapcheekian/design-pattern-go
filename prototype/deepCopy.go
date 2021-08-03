package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address2 struct {
	StreetAdress, City, Country string
}

func (a *Address2) DeepCopy2() *Address2 {
	return &Address2{
		a.StreetAdress,
		a.City,
		a.City,
	}
}

type Person struct {
	Name    string
	Address *Address2
	Friends []string
}

// deep copy by serialization and deserialization
func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)

	return &result
}

// manual deep copy
func (p *Person) DeepCopy2() *Person {
	q := *p
	q.Address = p.Address.DeepCopy2()
	copy(q.Friends, p.Friends)

	return &q
}

func _main() {
	john := Person{"John", &Address2{"123 Road", "Puchong", "Malaysia"}, []string{}}
	// jane := john

	// jane.Name = "Jane"
	// jane.Address.StreetAdress = "456 Road"

	// jane.Name = "Jane"
	// jane.Address = &Address{"456 Road", "Puchong", "Malaysia"}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAdress = "456 Road"

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
