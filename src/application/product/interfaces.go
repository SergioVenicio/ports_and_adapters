package product

type IProduct interface {
	IsValid() error
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

type ProductReader interface {
	Get(id string) (IProduct, error)
	List() ([]IProduct, error)
}

type ProductWriter interface {
	Save(product IProduct) error
}

type IProductPersistence interface {
	ProductReader
	ProductWriter
}

type IProductService interface {
	Get(id string) (IProduct, error)
	List() ([]IProduct, error)
	Create(name string, price float64) (IProduct, error)
	Enable(product IProduct) (IProduct, error)
	Disable(product IProduct) (IProduct, error)
}
