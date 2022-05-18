package product_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"

	mocks "github.com/sergio/go-hexagonal/application/mocks/application"
	"github.com/sergio/go-hexagonal/application/product"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedProduct := mocks.NewMockIProduct(ctrl)
	mockedPersistence := mocks.NewMockIProductPersistence(ctrl)

	mockedProduct.EXPECT().GetID().Return(uuid.NewV4().String()).AnyTimes()
	mockedPersistence.EXPECT().Get(gomock.Any()).Return(mockedProduct, nil).AnyTimes()

	service := product.ProductService{
		Persistece: mockedPersistence,
	}

	goted, err := service.Get(mockedProduct.GetID())
	require.Nil(t, err)
	require.Equal(t, mockedProduct, goted)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedProduct := mocks.NewMockIProduct(ctrl)
	mockedPersistence := mocks.NewMockIProductPersistence(ctrl)

	mockedProduct.EXPECT().GetID().Return(uuid.NewV4().String()).AnyTimes()
	mockedProduct.EXPECT().GetName().Return("test").AnyTimes()
	mockedPersistence.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()

	service := product.ProductService{
		Persistece: mockedPersistence,
	}

	product, err := service.Create("test", 1.99)
	require.Nil(t, err)
	require.Equal(t, mockedProduct.GetName(), "test")
	require.NotEmpty(t, product.GetID())
}

func TestProductService_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedProduct := mocks.NewMockIProduct(ctrl)
	mockedPersistence := mocks.NewMockIProductPersistence(ctrl)

	mockedProduct.EXPECT().Enable().Return(nil).AnyTimes()
	mockedProduct.EXPECT().GetStatus().Return(product.ENABLED).AnyTimes()
	mockedPersistence.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()

	service := product.ProductService{
		Persistece: mockedPersistence,
	}

	goted, err := service.Enable(mockedProduct)

	require.Nil(t, err)
	require.Equal(t, mockedProduct, goted)
	require.Equal(t, mockedProduct.GetStatus(), product.ENABLED)
}

func TestProductService_Disable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedProduct := mocks.NewMockIProduct(ctrl)
	persistence := mocks.NewMockIProductPersistence(ctrl)

	mockedProduct.EXPECT().Disable().Return(nil).AnyTimes()
	mockedProduct.EXPECT().GetStatus().Return(product.DISABLED).AnyTimes()
	persistence.EXPECT().Save(gomock.Any()).Return(nil).AnyTimes()

	service := product.ProductService{
		Persistece: persistence,
	}

	goted, err := service.Disable(mockedProduct)

	require.Nil(t, err)
	require.Equal(t, mockedProduct, goted)
	require.Equal(t, mockedProduct.GetStatus(), product.DISABLED)
}
