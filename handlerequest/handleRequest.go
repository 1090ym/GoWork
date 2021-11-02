package main

import "fmt"

type patient struct {
	Name string
}

type Department interface {
	Execute(*patient)
	SetNext(Department)
}

// Reception
type Reception struct {
	NextHandle Department
}

func NewReception() Department {
	return &Reception{}
}

func (r *Reception) Execute(p *patient) {
	fmt.Println("Reception patient")
	r.NextHandle.Execute(p)
}

func (r *Reception) SetNext(handle Department) {
	r.NextHandle = handle
}

// Doctor
type Doctor struct {
	NextHandle Department
}

func NewDoctor() Department {
	return &Doctor{}
}

func (d *Doctor) Execute(p *patient) {
	fmt.Println("check patience")
	d.NextHandle.Execute(p)
}
func (d *Doctor) SetNext(handle Department) {
	d.NextHandle = handle
}

// submit
type Payment struct {
	NextHandle Department
}

func NewPayment() Department {
	return &Payment{}
}

func (pay *Payment) Execute(p *patient) {
	fmt.Println("pay money")
}

func (pay *Payment) SetNext(handle Department) {
	pay.NextHandle = handle
}

func main() {
	p := patient{
		Name: "Alice",
	}
	rep := NewReception()
	doc := NewDoctor()
	pay := NewPayment()

	// 责任链
	rep.SetNext(doc)
	doc.SetNext(pay)

	rep.Execute(&p)
}
