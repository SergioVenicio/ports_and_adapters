package cli_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/sergio/go-hexagonal/adapters/cli"
	mockApp "github.com/sergio/go-hexagonal/application/mocks/application"
)

func TestRun_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "Product Test"
	price := 1.99
	status := "enabled"
	id := "abc"

	productMock := mockApp.NewMockIProduct(ctrl)
	productMock.EXPECT().GetID().Return(id).AnyTimes()
	productMock.EXPECT().GetName().Return(name).AnyTimes()
	productMock.EXPECT().GetStatus().Return(status).AnyTimes()
	productMock.EXPECT().GetPrice().Return(price).AnyTimes()

	productServiceMock := mockApp.NewMockIProductService(ctrl)
	productServiceMock.EXPECT().Create(name, price).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Get(id).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	expected := "Product ID #abc with the name Product Test with price 1.99 and status enabled has been created"
	result, err := cli.Run(productServiceMock, "create", "", name, price)

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestRun_Enable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "Product Test"
	price := 1.99
	status := "enabled"
	id := "abc"

	productMock := mockApp.NewMockIProduct(ctrl)
	productMock.EXPECT().GetID().Return(id).AnyTimes()
	productMock.EXPECT().GetName().Return(name).AnyTimes()
	productMock.EXPECT().GetStatus().Return(status).AnyTimes()
	productMock.EXPECT().GetPrice().Return(price).AnyTimes()

	productServiceMock := mockApp.NewMockIProductService(ctrl)
	productServiceMock.EXPECT().Get(id).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()

	expected := "Product ID #abc has been enabled"
	result, err := cli.Run(productServiceMock, "enable", id, name, price)

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestRun_Disable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "Product Test"
	price := 1.99
	status := "enabled"
	id := "abc"

	productMock := mockApp.NewMockIProduct(ctrl)
	productMock.EXPECT().GetID().Return(id).AnyTimes()
	productMock.EXPECT().GetName().Return(name).AnyTimes()
	productMock.EXPECT().GetStatus().Return(status).AnyTimes()
	productMock.EXPECT().GetPrice().Return(price).AnyTimes()

	productServiceMock := mockApp.NewMockIProductService(ctrl)
	productServiceMock.EXPECT().Get(id).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	expected := "Product ID #abc has been disabled"
	result, err := cli.Run(productServiceMock, "disable", id, name, price)

	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestRun_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "Product Test"
	price := 1.99
	status := "enabled"
	id := "abc"

	productMock := mockApp.NewMockIProduct(ctrl)
	productMock.EXPECT().GetID().Return(id).AnyTimes()
	productMock.EXPECT().GetName().Return(name).AnyTimes()
	productMock.EXPECT().GetStatus().Return(status).AnyTimes()
	productMock.EXPECT().GetPrice().Return(price).AnyTimes()

	productServiceMock := mockApp.NewMockIProductService(ctrl)
	productServiceMock.EXPECT().Get(id).Return(productMock, nil).AnyTimes()

	expected := "Product ID abc,\n\t\t\tProduct Name Product Test,\n\t\t\tProduct Status enabled,\n\t\t\tProduct Price 1.99\n\t\t\t"
	result, err := cli.Run(productServiceMock, "", id, "", 0)

	require.Nil(t, err)
	require.Equal(t, expected, result)
}
