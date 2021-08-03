package main

import "fmt"

type Person2 struct {
	name, position string
}

type person2Mod func(*Person2)

type Person2Builder struct {
	actions []person2Mod
}

func (b *Person2Builder) Called(name string) *Person2Builder {
	b.actions = append(b.actions, func(p *Person2) {
		p.name = name
	})

	return b
}

func (b *Person2Builder) WorkAsA(position string) *Person2Builder {
	b.actions = append(b.actions, func(p *Person2) {
		p.position = position
	})

	return b
}

func (b *Person2Builder) Build() *Person2 {
	p := Person2{}

	for _, action := range b.actions {
		action(&p)
	}

	return &p
}

func main() {
	b := Person2Builder{}
	p := b.Called("yap").WorkAsA("enginner").Build()
	fmt.Println(*&p.name)
}
