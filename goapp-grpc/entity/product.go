package entity

import uuid "github.com/satori/go.uuid"

type Product struct {
	ID string
	Name string
}

type Products struct {
	Elements []Product
}

func NewProduct(name string) *Product {
	return &Product{ID: uuid.NewV4().String(), Name: name}
}

func (ps *Products) Add (p *Product) {
	ps.Elements = append(ps.Elements, *p)
}

func NewProducts() *Products {
	return &Products{}
}
