package service

import (
	domain2 "example.com/m/v2/internal/domain"
	"example.com/m/v2/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBasket(t *testing.T) {
	testTable := []struct {
		name            string
		idUser          int
		returnResp      domain2.Basket
		returnError     error
		returnRespRepo  *domain2.Basket
		returnErrorRepo error
		isCheckError    bool
	}{
		{
			name:   "ok",
			idUser: 1,
			returnRespRepo: &domain2.Basket{
				ID: 1,
				Products: []domain2.BasketProduct{
					{
						ID:         1,
						BasketID:   1,
						ProductID:  1,
						Count:      2,
						TotalPrice: 200,
					},
				},
			},
			returnResp: domain2.Basket{
				ID:     1,
				UserID: 1,
				Products: []domain2.BasketProduct{
					{
						ID:         1,
						BasketID:   1,
						ProductID:  1,
						Count:      2,
						TotalPrice: 200,
					},
				},
				TotalPrice: 200,
			},
			returnError:     nil,
			returnErrorRepo: nil,
			isCheckError:    false,
		},
		{
			name:            "error",
			idUser:          1,
			returnRespRepo:  nil,
			returnResp:      domain2.Basket{},
			returnError:     testErr,
			returnErrorRepo: testErr,
			isCheckError:    true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockBasketsRepository := mock_repository.NewMockBasketsStorage(c)
			mockProductsRepository := mock_repository.NewMockProductsStorage(c)
			basketsService := NewBasketService(mockBasketsRepository, mockProductsRepository)

			mockBasketsRepository.
				EXPECT().
				GetBasket(testCase.idUser).
				Return(testCase.returnRespRepo, testCase.returnErrorRepo)

			respBasket, err := basketsService.GetBasket(testCase.idUser)
			if testCase.isCheckError {
				assert.ErrorIs(t, err, testCase.returnError)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, testCase.returnResp, respBasket)
			}
		})
	}
}

func TestAddProductToBasket(t *testing.T) {
	testTable := []struct {
		name         string
		inputProduct domain2.BasketProduct
		returnResp   domain2.BasketProduct
		returnError  error

		isCallRepositoryMethodGetBasketProduct bool
		returnRespRepoGetBasketProduct         *domain2.BasketProduct
		returnErrorRepoGetBasketProduct        error

		isCallRepositoryMethodEditBasketProduct bool
		firstArgumentEditBasketProduct          domain2.BasketProduct
		returnRespRepoEditBasketProduct         domain2.BasketProduct
		returnErrorRepoEditBasketProduct        error

		isCallRepositoryMethodAddBasketProduct bool
		returnRespRepoAddBasketProduct         domain2.BasketProduct
		returnErrorRepoAddBasketProduct        error

		isCallRepositoryMethodGetProduct bool
		returnRespRepoGetProduct         *domain2.Product
		returnErrorRepoGetProduct        error

		isCheckError bool
	}{
		{
			name: "Add product count",
			inputProduct: domain2.BasketProduct{
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnResp: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     3,
			},
			returnError: nil,

			isCallRepositoryMethodGetBasketProduct: true,
			returnRespRepoGetBasketProduct: &domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     2,
			},
			returnErrorRepoGetBasketProduct: nil,

			isCallRepositoryMethodEditBasketProduct: true,
			firstArgumentEditBasketProduct: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     3,
			},
			returnRespRepoEditBasketProduct: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     3,
			},
			returnErrorRepoEditBasketProduct: nil,

			isCallRepositoryMethodAddBasketProduct: false,
			returnRespRepoAddBasketProduct: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     2,
			},
			returnErrorRepoAddBasketProduct: nil,

			isCallRepositoryMethodGetProduct: false,
			returnRespRepoGetProduct: &domain2.Product{
				ID:          1,
				Name:        "car",
				Description: "car is a cool",
				Price:       100,
			},
			returnErrorRepoGetProduct: nil,

			isCheckError: false,
		},
		{
			name: "Add product",
			inputProduct: domain2.BasketProduct{
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnResp: domain2.BasketProduct{
				ID:         1,
				BasketID:   1,
				ProductID:  1,
				Count:      1,
				TotalPrice: 100,
			},
			returnError: nil,

			isCallRepositoryMethodGetBasketProduct: true,
			returnRespRepoGetBasketProduct:         nil,
			returnErrorRepoGetBasketProduct:        testErr,

			isCallRepositoryMethodEditBasketProduct: false,
			firstArgumentEditBasketProduct:          domain2.BasketProduct{},
			returnRespRepoEditBasketProduct:         domain2.BasketProduct{},
			returnErrorRepoEditBasketProduct:        nil,

			isCallRepositoryMethodAddBasketProduct: true,
			returnRespRepoAddBasketProduct: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnErrorRepoAddBasketProduct: nil,

			isCallRepositoryMethodGetProduct: true,
			returnRespRepoGetProduct: &domain2.Product{
				ID:          1,
				Name:        "car",
				Description: "car is a cool",
				Price:       100,
			},
			returnErrorRepoGetProduct: nil,

			isCheckError: false,
		},
		{
			name: "Error method EditBasketProduct",
			inputProduct: domain2.BasketProduct{
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnResp: domain2.BasketProduct{
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnError: domain2.ErrBasketProductNotFound,

			isCallRepositoryMethodGetBasketProduct: true,
			returnRespRepoGetBasketProduct: &domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     2,
			},
			returnErrorRepoGetBasketProduct: nil,

			isCallRepositoryMethodEditBasketProduct: true,
			firstArgumentEditBasketProduct: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     3,
			},
			returnRespRepoEditBasketProduct:  domain2.BasketProduct{},
			returnErrorRepoEditBasketProduct: testErr,

			isCallRepositoryMethodAddBasketProduct: false,
			returnRespRepoAddBasketProduct:         domain2.BasketProduct{},
			returnErrorRepoAddBasketProduct:        nil,

			isCallRepositoryMethodGetProduct: false,
			returnRespRepoGetProduct:         nil,
			returnErrorRepoGetProduct:        nil,

			isCheckError: true,
		},
		{
			name: "Error method AddBasketProduct",
			inputProduct: domain2.BasketProduct{
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnResp:  domain2.BasketProduct{},
			returnError: testErr,

			isCallRepositoryMethodGetBasketProduct: true,
			returnRespRepoGetBasketProduct:         nil,
			returnErrorRepoGetBasketProduct:        testErr,

			isCallRepositoryMethodEditBasketProduct: false,
			firstArgumentEditBasketProduct:          domain2.BasketProduct{},
			returnRespRepoEditBasketProduct:         domain2.BasketProduct{},
			returnErrorRepoEditBasketProduct:        nil,

			isCallRepositoryMethodAddBasketProduct: true,
			returnRespRepoAddBasketProduct:         domain2.BasketProduct{},
			returnErrorRepoAddBasketProduct:        testErr,

			isCallRepositoryMethodGetProduct: false,
			returnRespRepoGetProduct:         nil,
			returnErrorRepoGetProduct:        nil,

			isCheckError: true,
		},
		{
			name: "Error method GetProduct",
			inputProduct: domain2.BasketProduct{
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnResp:  domain2.BasketProduct{},
			returnError: testErr,

			isCallRepositoryMethodGetBasketProduct: true,
			returnRespRepoGetBasketProduct:         nil,
			returnErrorRepoGetBasketProduct:        testErr,

			isCallRepositoryMethodEditBasketProduct: false,
			firstArgumentEditBasketProduct:          domain2.BasketProduct{},
			returnRespRepoEditBasketProduct:         domain2.BasketProduct{},
			returnErrorRepoEditBasketProduct:        nil,

			isCallRepositoryMethodAddBasketProduct: true,
			returnRespRepoAddBasketProduct: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnErrorRepoAddBasketProduct: nil,

			isCallRepositoryMethodGetProduct: true,
			returnRespRepoGetProduct:         nil,
			returnErrorRepoGetProduct:        testErr,

			isCheckError: true,
		},
		{
			name: "No BasketID (validation)",
			inputProduct: domain2.BasketProduct{
				ProductID: 1,
				Count:     1,
			},
			returnResp:  domain2.BasketProduct{},
			returnError: domain2.ErrBasketProductNoBasketID,

			isCallRepositoryMethodGetBasketProduct: false,
			returnRespRepoGetBasketProduct:         nil,
			returnErrorRepoGetBasketProduct:        nil,

			isCallRepositoryMethodEditBasketProduct: false,
			firstArgumentEditBasketProduct:          domain2.BasketProduct{},
			returnRespRepoEditBasketProduct:         domain2.BasketProduct{},
			returnErrorRepoEditBasketProduct:        nil,

			isCallRepositoryMethodAddBasketProduct: false,
			returnRespRepoAddBasketProduct:         domain2.BasketProduct{},
			returnErrorRepoAddBasketProduct:        nil,

			isCallRepositoryMethodGetProduct: false,
			returnRespRepoGetProduct:         nil,
			returnErrorRepoGetProduct:        nil,

			isCheckError: true,
		},
		{
			name: "No ProductID (validation)",
			inputProduct: domain2.BasketProduct{
				BasketID: 1,
				Count:    1,
			},
			returnResp:  domain2.BasketProduct{},
			returnError: domain2.ErrBasketProductNoProductID,

			isCallRepositoryMethodGetBasketProduct: false,
			returnRespRepoGetBasketProduct:         nil,
			returnErrorRepoGetBasketProduct:        nil,

			isCallRepositoryMethodEditBasketProduct: false,
			firstArgumentEditBasketProduct:          domain2.BasketProduct{},
			returnRespRepoEditBasketProduct:         domain2.BasketProduct{},
			returnErrorRepoEditBasketProduct:        nil,

			isCallRepositoryMethodAddBasketProduct: false,
			returnRespRepoAddBasketProduct:         domain2.BasketProduct{},
			returnErrorRepoAddBasketProduct:        nil,

			isCallRepositoryMethodGetProduct: false,
			returnRespRepoGetProduct:         nil,
			returnErrorRepoGetProduct:        nil,

			isCheckError: true,
		},
		{
			name: "No Count (validation)",
			inputProduct: domain2.BasketProduct{
				ProductID: 1,
				BasketID:  1,
			},
			returnResp:  domain2.BasketProduct{},
			returnError: domain2.ErrBasketProductNoCount,

			isCallRepositoryMethodGetBasketProduct: false,
			returnRespRepoGetBasketProduct:         nil,
			returnErrorRepoGetBasketProduct:        nil,

			isCallRepositoryMethodEditBasketProduct: false,
			firstArgumentEditBasketProduct:          domain2.BasketProduct{},
			returnRespRepoEditBasketProduct:         domain2.BasketProduct{},
			returnErrorRepoEditBasketProduct:        nil,

			isCallRepositoryMethodAddBasketProduct: false,
			returnRespRepoAddBasketProduct:         domain2.BasketProduct{},
			returnErrorRepoAddBasketProduct:        nil,

			isCallRepositoryMethodGetProduct: false,
			returnRespRepoGetProduct:         nil,
			returnErrorRepoGetProduct:        nil,

			isCheckError: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockBasketsRepository := mock_repository.NewMockBasketsStorage(c)
			mockProductsRepository := mock_repository.NewMockProductsStorage(c)
			basketsService := NewBasketService(mockBasketsRepository, mockProductsRepository)

			if testCase.isCallRepositoryMethodGetBasketProduct {
				mockBasketsRepository.
					EXPECT().
					GetBasketProduct(testCase.inputProduct.BasketID, testCase.inputProduct.ProductID).
					Return(testCase.returnRespRepoGetBasketProduct, testCase.returnErrorRepoGetBasketProduct)
			}

			if testCase.isCallRepositoryMethodEditBasketProduct {
				mockBasketsRepository.
					EXPECT().
					EditBasketProduct(testCase.firstArgumentEditBasketProduct).
					Return(testCase.returnRespRepoEditBasketProduct, testCase.returnErrorRepoEditBasketProduct)
			}

			if testCase.isCallRepositoryMethodAddBasketProduct {
				mockBasketsRepository.
					EXPECT().
					AddBasketProduct(testCase.inputProduct).
					Return(testCase.returnRespRepoAddBasketProduct, testCase.returnErrorRepoAddBasketProduct)
			}

			if testCase.isCallRepositoryMethodGetProduct {
				mockProductsRepository.
					EXPECT().
					GetProduct(testCase.inputProduct.ProductID).
					Return(testCase.returnRespRepoGetProduct, testCase.returnErrorRepoGetProduct)
			}

			respBasketProduct, err := basketsService.AddProductToBasket(testCase.inputProduct)
			if testCase.isCheckError {
				assert.ErrorIs(t, err, testCase.returnError)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, testCase.returnResp, respBasketProduct)
			}
		})
	}
}

func TestDecreaseQuantityProductToBasket(t *testing.T) {
	testTable := []struct {
		name         string
		inputProduct domain2.BasketProduct
		returnResp   domain2.BasketProduct
		returnError  error

		isCallRepositoryMethodGetBasketProduct bool
		returnRespRepoGetBasketProduct         *domain2.BasketProduct
		returnErrorRepoGetBasketProduct        error

		isCallRepositoryMethodDeleteBasketProduct bool
		returnRespRepoDeleteBasketProduct         bool
		returnErrorRepoDeleteBasketProduct        error

		isCallRepositoryMethodEditBasketProduct bool
		firstArgumentEditBasketProduct          domain2.BasketProduct
		returnRespRepoEditBasketProduct         domain2.BasketProduct
		returnErrorRepoEditBasketProduct        error

		isCallRepositoryMethodGetProduct bool
		returnRespRepoGetProduct         *domain2.Product
		returnErrorRepoGetProduct        error

		isCheckError bool
	}{
		{
			name: "Delete product count",
			inputProduct: domain2.BasketProduct{
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnResp: domain2.BasketProduct{
				ID:         1,
				BasketID:   1,
				ProductID:  1,
				Count:      0,
				TotalPrice: 0,
			},
			returnError: nil,

			isCallRepositoryMethodGetBasketProduct: true,
			returnRespRepoGetBasketProduct: &domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnErrorRepoGetBasketProduct: nil,

			isCallRepositoryMethodDeleteBasketProduct: true,
			returnRespRepoDeleteBasketProduct:         true,
			returnErrorRepoDeleteBasketProduct:        nil,

			isCallRepositoryMethodEditBasketProduct: false,
			firstArgumentEditBasketProduct:          domain2.BasketProduct{},
			returnRespRepoEditBasketProduct:         domain2.BasketProduct{},
			returnErrorRepoEditBasketProduct:        nil,

			isCallRepositoryMethodGetProduct: false,
			returnRespRepoGetProduct:         nil,
			returnErrorRepoGetProduct:        nil,

			isCheckError: false,
		},
		{
			name: "Decrease product count",
			inputProduct: domain2.BasketProduct{
				BasketID:  1,
				ProductID: 1,
				Count:     2,
			},
			returnResp: domain2.BasketProduct{
				ID:         1,
				BasketID:   1,
				ProductID:  1,
				Count:      1,
				TotalPrice: 100,
			},
			returnError: nil,

			isCallRepositoryMethodGetBasketProduct: true,
			returnRespRepoGetBasketProduct: &domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     2,
			},
			returnErrorRepoGetBasketProduct: nil,

			isCallRepositoryMethodDeleteBasketProduct: false,
			returnRespRepoDeleteBasketProduct:         false,
			returnErrorRepoDeleteBasketProduct:        nil,

			isCallRepositoryMethodEditBasketProduct: true,
			firstArgumentEditBasketProduct: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnRespRepoEditBasketProduct: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnErrorRepoEditBasketProduct: nil,

			isCallRepositoryMethodGetProduct: true,
			returnRespRepoGetProduct: &domain2.Product{
				ID:          1,
				Name:        "car",
				Description: "car is a cool",
				Price:       100,
			},
			returnErrorRepoGetProduct: nil,

			isCheckError: false,
		},
		{
			name: "Error method GetBasketProduct",
			inputProduct: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnResp: domain2.BasketProduct{
				ID:         1,
				BasketID:   1,
				ProductID:  1,
				Count:      0,
				TotalPrice: 0,
			},
			returnError: nil,

			isCallRepositoryMethodGetBasketProduct: true,
			returnRespRepoGetBasketProduct:         nil,
			returnErrorRepoGetBasketProduct:        testErr,

			isCallRepositoryMethodDeleteBasketProduct: false,
			returnRespRepoDeleteBasketProduct:         false,
			returnErrorRepoDeleteBasketProduct:        nil,

			isCallRepositoryMethodEditBasketProduct: false,
			firstArgumentEditBasketProduct:          domain2.BasketProduct{},
			returnRespRepoEditBasketProduct:         domain2.BasketProduct{},
			returnErrorRepoEditBasketProduct:        nil,

			isCallRepositoryMethodGetProduct: false,
			returnRespRepoGetProduct:         nil,
			returnErrorRepoGetProduct:        nil,

			isCheckError: false,
		},
		{
			name: "Error method DeleteBasketProduct",
			inputProduct: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnResp: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnError: testErr,

			isCallRepositoryMethodGetBasketProduct: true,
			returnRespRepoGetBasketProduct: &domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnErrorRepoGetBasketProduct: nil,

			isCallRepositoryMethodDeleteBasketProduct: true,
			returnRespRepoDeleteBasketProduct:         false,
			returnErrorRepoDeleteBasketProduct:        testErr,

			isCallRepositoryMethodEditBasketProduct: false,
			firstArgumentEditBasketProduct:          domain2.BasketProduct{},
			returnRespRepoEditBasketProduct:         domain2.BasketProduct{},
			returnErrorRepoEditBasketProduct:        nil,

			isCallRepositoryMethodGetProduct: false,
			returnRespRepoGetProduct:         nil,
			returnErrorRepoGetProduct:        nil,

			isCheckError: true,
		},
		{
			name: "isDeleted = false, method DeleteBasketProduct",
			inputProduct: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnResp: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnError: domain2.ErrBasketNotDeleted,

			isCallRepositoryMethodGetBasketProduct: true,
			returnRespRepoGetBasketProduct: &domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnErrorRepoGetBasketProduct: nil,

			isCallRepositoryMethodDeleteBasketProduct: true,
			returnRespRepoDeleteBasketProduct:         false,
			returnErrorRepoDeleteBasketProduct:        nil,

			isCallRepositoryMethodEditBasketProduct: false,
			firstArgumentEditBasketProduct:          domain2.BasketProduct{},
			returnRespRepoEditBasketProduct:         domain2.BasketProduct{},
			returnErrorRepoEditBasketProduct:        nil,

			isCallRepositoryMethodGetProduct: false,
			returnRespRepoGetProduct:         nil,
			returnErrorRepoGetProduct:        nil,

			isCheckError: true,
		},
		{
			name: "Error method EditBasketProduct",
			inputProduct: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     2,
			},
			returnResp: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnError: domain2.ErrBasketProductNotFound,

			isCallRepositoryMethodGetBasketProduct: true,
			returnRespRepoGetBasketProduct: &domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     2,
			},
			returnErrorRepoGetBasketProduct: nil,

			isCallRepositoryMethodDeleteBasketProduct: false,
			returnRespRepoDeleteBasketProduct:         false,
			returnErrorRepoDeleteBasketProduct:        nil,

			isCallRepositoryMethodEditBasketProduct: true,
			firstArgumentEditBasketProduct: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnRespRepoEditBasketProduct:  domain2.BasketProduct{},
			returnErrorRepoEditBasketProduct: testErr,

			isCallRepositoryMethodGetProduct: false,
			returnRespRepoGetProduct:         nil,
			returnErrorRepoGetProduct:        nil,

			isCheckError: true,
		},
		{
			name: "Error method GetProduct",
			inputProduct: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     2,
			},
			returnResp: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnError: testErr,

			isCallRepositoryMethodGetBasketProduct: true,
			returnRespRepoGetBasketProduct: &domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     2,
			},
			returnErrorRepoGetBasketProduct: nil,

			isCallRepositoryMethodDeleteBasketProduct: false,
			returnRespRepoDeleteBasketProduct:         false,
			returnErrorRepoDeleteBasketProduct:        nil,

			isCallRepositoryMethodEditBasketProduct: true,
			firstArgumentEditBasketProduct: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnRespRepoEditBasketProduct: domain2.BasketProduct{
				ID:        1,
				BasketID:  1,
				ProductID: 1,
				Count:     1,
			},
			returnErrorRepoEditBasketProduct: nil,

			isCallRepositoryMethodGetProduct: true,
			returnRespRepoGetProduct:         nil,
			returnErrorRepoGetProduct:        testErr,

			isCheckError: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockBasketsRepository := mock_repository.NewMockBasketsStorage(c)
			mockProductsRepository := mock_repository.NewMockProductsStorage(c)
			basketsService := NewBasketService(mockBasketsRepository, mockProductsRepository)

			if testCase.isCallRepositoryMethodGetBasketProduct {
				mockBasketsRepository.
					EXPECT().
					GetBasketProduct(testCase.inputProduct.BasketID, testCase.inputProduct.ProductID).
					Return(testCase.returnRespRepoGetBasketProduct, testCase.returnErrorRepoGetBasketProduct)
			}

			if testCase.isCallRepositoryMethodDeleteBasketProduct {
				mockBasketsRepository.
					EXPECT().
					DeleteBasketProduct(1).
					Return(testCase.returnRespRepoDeleteBasketProduct, testCase.returnErrorRepoDeleteBasketProduct)
			}

			if testCase.isCallRepositoryMethodEditBasketProduct {
				mockBasketsRepository.
					EXPECT().
					EditBasketProduct(testCase.firstArgumentEditBasketProduct).
					Return(testCase.returnRespRepoEditBasketProduct, testCase.returnErrorRepoEditBasketProduct)
			}

			if testCase.isCallRepositoryMethodGetProduct {
				mockProductsRepository.
					EXPECT().
					GetProduct(testCase.inputProduct.ProductID).
					Return(testCase.returnRespRepoGetProduct, testCase.returnErrorRepoGetProduct)
			}

			respBasketProduct, err := basketsService.DecreaseQuantityProductToBasket(testCase.inputProduct)
			if testCase.isCheckError {
				assert.ErrorIs(t, err, testCase.returnError)
			} else {
				assert.NoError(t, err)
				assert.EqualValues(t, testCase.returnResp, respBasketProduct)
			}
		})
	}
}

func TestDecreaseDeleteProductToBasket(t *testing.T) {
	testTable := []struct {
		name                               string
		inputBasketID                      int
		returnError                        error
		returnRespRepoDeleteBasketProduct  bool
		returnErrorRepoDeleteBasketProduct error
	}{
		{
			name:                               "ok",
			inputBasketID:                      1,
			returnError:                        nil,
			returnRespRepoDeleteBasketProduct:  true,
			returnErrorRepoDeleteBasketProduct: nil,
		},
		{
			name:                               "error",
			inputBasketID:                      1,
			returnError:                        testErr,
			returnRespRepoDeleteBasketProduct:  false,
			returnErrorRepoDeleteBasketProduct: testErr,
		},
		{
			name:                               "isDeleted = false",
			inputBasketID:                      1,
			returnError:                        domain2.ErrBasketProductNotFound,
			returnRespRepoDeleteBasketProduct:  false,
			returnErrorRepoDeleteBasketProduct: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockBasketsRepository := mock_repository.NewMockBasketsStorage(c)
			mockProductsRepository := mock_repository.NewMockProductsStorage(c)
			basketsService := NewBasketService(mockBasketsRepository, mockProductsRepository)

			mockBasketsRepository.
				EXPECT().
				DeleteBasketProduct(1).
				Return(testCase.returnRespRepoDeleteBasketProduct, testCase.returnErrorRepoDeleteBasketProduct)

			err := basketsService.DeleteProductToBasket(testCase.inputBasketID)
			assert.ErrorIs(t, err, testCase.returnError)
		})
	}
}
