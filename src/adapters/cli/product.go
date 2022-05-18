package cli

import (
	"fmt"

	"github.com/sergio/go-hexagonal/application/product"
)

func Run(service product.IProductService, action string, id string, name string, price float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		newProduct, err := service.Create(name, price)
		if err != nil {
			return "", err
		}

		result = fmt.Sprintf(
			"Product ID #%s with the name %s with price %.2f and status %s has been created",
			newProduct.GetID(),
			newProduct.GetName(),
			newProduct.GetPrice(),
			newProduct.GetStatus(),
		)
	case "enable":
		dbProduct, err := service.Get(id)
		if err != nil {
			return "", err
		}

		service.Enable(dbProduct)
		result = fmt.Sprintf(
			"Product ID #%s has been enabled",
			dbProduct.GetID(),
		)
	case "disable":
		dbProduct, err := service.Get(id)
		if err != nil {
			return "", err
		}

		service.Disable(dbProduct)
		result = fmt.Sprintf(
			"Product ID #%s has been disabled",
			dbProduct.GetID(),
		)
	default:
		dbProduct, err := service.Get(id)
		if err != nil {
			return "", err
		}

		result = fmt.Sprintf(
			`Product ID %s,
			Product Name %s,
			Product Status %s,
			Product Price %.2f
			`,
			dbProduct.GetID(),
			dbProduct.GetName(),
			dbProduct.GetStatus(),
			dbProduct.GetPrice(),
		)
	}

	return result, nil
}
