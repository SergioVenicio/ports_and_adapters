package dto

import appProduct "github.com/sergio/go-hexagonal/application/product"

type Product struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func NewProductDTO() *Product {
	return &Product{}
}

func (p *Product) Bind(product *appProduct.Product) (*appProduct.Product, error) {
	if p.ID != "" {
		product.ID = p.ID
	}

	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status

	err := product.IsValid()
	if err != nil {
		return nil, err
	}

	return product, nil
}
