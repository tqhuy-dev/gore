package example

import (
	"github.com/tqhuy-dev/gore/db"
)

type ProductsEntity struct {
	ID          int
	Name        string
	Price       float64
	Stock       int
	Image       string
	Status      int8
	Sku         string
	Description *string
}

type ProductsRepository[T any, Id any] interface {
	ExampleFunc()
	BaseRepo() db.BaseRepo[T, Id]
}

type productRepo[T any, Id any] struct {
	baseRepo db.BaseRepo[T, Id]
}

func (p *productRepo[T, Id]) BaseRepo() db.BaseRepo[T, Id] {
	return p.baseRepo
}

func (p *productRepo[T, Id]) ExampleFunc() {
}

func NewProductRepository[T any, Id any]() ProductsRepository[T, Id] {
	return &productRepo[T, Id]{}
}
