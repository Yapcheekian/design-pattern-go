package main

import "fmt"

type Person struct {
	// address
	StreetAddress, Postcode, City string

	// job
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBulder struct {
	person *Person
}

func (b *PersonBulder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBulder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func NewPersonBuilder() *PersonBulder {
	return &PersonBulder{&Person{}}
}

type PersonAddressBuilder struct {
	PersonBulder
}

type PersonJobBuilder struct {
	PersonBulder
}

func (b *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	b.person.StreetAddress = streetAddress
	return b
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}

func (b *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	b.person.Postcode = postcode
	return b
}

func (pjb *PersonJobBuilder) At(company string) *PersonJobBuilder {
	pjb.person.CompanyName = company
	return pjb
}

func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

func (pjb *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	pjb.person.AnnualIncome = annualIncome
	return pjb
}

func (b *PersonBulder) Build() *Person {
	return b.person
}

func BuilderFacet() {
	pb := NewPersonBuilder()

	pb.Lives().At("Jalan Impiana").In("Puchong").WithPostcode("47120")
	pb.Works().At("Cyberbiz").AsA("SRE").Earning(1000)

	fmt.Println(pb.Build())
}
