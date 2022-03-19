package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mrpsousa/go-hexagonal/application"
	mock_application "github.com/mrpsousa/go-hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
)

func TesteProductService_Get(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	// persistence.EXPECT().Get(gomock.Any()).Return(product, nil).Anytimes()
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil)

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}
