package main

import "strings"

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (eb *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contains @")
	}
	eb.email.from = from
	return eb
}

func (eb *EmailBuilder) To(to string) *EmailBuilder {
	if !strings.Contains(to, "@") {
		panic("email should contains @")
	}
	eb.email.to = to
	return eb
}

func (eb *EmailBuilder) Subject(subject string) *EmailBuilder {
	eb.email.subject = subject
	return eb
}

func (eb *EmailBuilder) Body(body string) *EmailBuilder {
	eb.email.body = body
	return eb
}

type build func(*EmailBuilder)

func sendMailImpl(email *email) {}

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendMailImpl(&builder.email)
}

func main() {
	SendEmail(func(eb *EmailBuilder) {
		eb.From("a@a.com").
			To("b@b.com").
			Subject("test").
			Body("123")
	})
}
