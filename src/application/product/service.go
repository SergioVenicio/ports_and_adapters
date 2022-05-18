package product

type ProductService struct {
	Persistece IProductPersistence
}

func NewProductService(persistence IProductPersistence) *ProductService {
	return &ProductService{
		Persistece: persistence,
	}
}

func (s *ProductService) List() ([]IProduct, error) {
	products, err := s.Persistece.List()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductService) Get(id string) (IProduct, error) {
	product, err := s.Persistece.Get(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Create(name string, price float64) (IProduct, error) {
	newProduct := NewProduct(name, price, ENABLED)
	err := newProduct.IsValid()
	if err != nil {
		return nil, err
	}

	err = s.Persistece.Save(newProduct)
	if err != nil {
		return nil, err
	}

	return newProduct, nil
}

func (s *ProductService) Enable(product IProduct) (IProduct, error) {
	err := product.Enable()
	if err != nil {
		return nil, err
	}

	err = s.Persistece.Save(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Disable(product IProduct) (IProduct, error) {
	err := product.Disable()
	if err != nil {
		return nil, err
	}

	err = s.Persistece.Save(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
