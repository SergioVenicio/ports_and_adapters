package product

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float"`
	Status string  `valid:"required"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewProduct(name string, price float64, status string) *Product {
	return &Product{
		ID:     uuid.NewV4().String(),
		Name:   name,
		Price:  price,
		Status: status,
	}
}

func (p *Product) IsValid() error {
	_, err := govalidator.ValidateStruct(p)
	return err
}

func (p *Product) Enable() error {
	err := p.IsValid()
	if err != nil {
		return err
	}

	p.Status = ENABLED
	return nil
}

func (p *Product) Disable() error {
	err := p.IsValid()
	if err != nil {
		return err
	}

	p.Status = DISABLED
	return nil
}

func (p *Product) GetID() string {
	return string(p.ID)
}

func (p *Product) GetName() string {
	return string(p.Name)
}

func (p *Product) GetStatus() string {
	return string(p.Status)
}

func (p *Product) GetPrice() float64 {
	return float64(p.Price)
}
