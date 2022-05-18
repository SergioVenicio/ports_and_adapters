package product_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/stretchr/testify/require"

	"github.com/sergio/go-hexagonal/application/product"
)

func TestNewProduct(t *testing.T) {
	got := product.NewProduct("Test", 1.99, product.ENABLED)

	if got.GetID() == "" {
		t.Errorf("invalid id got %s", got.ID)
	}

	if got.GetName() != "Test" {
		t.Errorf("invalid name expect Test got %s", got.Name)
	}

	if got.GetStatus() != product.ENABLED {
		t.Errorf("invalid status expect %s got %s", product.ENABLED, got.Status)
	}

	if got.GetPrice() != 1.99 {
		t.Errorf("invalid price expect 1.99 got %f", got.Price)
	}
}

func TestProductEnable(t *testing.T) {
	id := uuid.NewV4().String()
	got := product.Product{
		ID:     id,
		Name:   "Test",
		Status: product.DISABLED,
		Price:  1.99,
	}

	err := got.Enable()
	require.Nil(t, err)

	if got.Status != product.ENABLED {
		t.Errorf("cant enable product expect %s got %s", product.ENABLED, got.Status)
	}
}

func TestProductDisable(t *testing.T) {
	id := uuid.NewV4().String()
	got := product.Product{
		ID:     id,
		Name:   "Test",
		Status: product.ENABLED,
		Price:  1.99,
	}

	err := got.Disable()
	require.Nil(t, err)
	if got.Status != product.DISABLED {
		t.Errorf("cant disable product expect %s got %s", product.DISABLED, got.Status)
	}
}

func TestProductCantEnableWithInvalidPrice(t *testing.T) {
	id := uuid.NewV4().String()
	got := product.Product{
		ID:     id,
		Name:   "Test",
		Status: product.DISABLED,
		Price:  0,
	}

	err := got.Enable()
	require.Equal(t, "Price: Missing required field", err.Error())
}

func TestProductIsValid(t *testing.T) {
	id := uuid.NewV4().String()
	got := product.Product{
		ID:     id,
		Name:   "Test",
		Status: product.ENABLED,
		Price:  1.99,
	}

	err := got.IsValid()
	require.Nil(t, err)

	got.Price = 0
	err = got.Enable()
	require.Equal(t, "Price: Missing required field", err.Error())
}
