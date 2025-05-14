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

type ProductsRepository interface {
	ExampleFunc()
	BaseRepo() db.BaseRepo[ProductsEntity, int]
}

type productRepo struct {
	baseRepo db.BaseRepo[ProductsEntity, int]
}

func (p *productRepo) BaseRepo() db.BaseRepo[ProductsEntity, int] {
	return p.baseRepo
}

func (p *productRepo) ExampleFunc() {
}

func NewProductRepository() ProductsRepository {
	return &productRepo{}
}
