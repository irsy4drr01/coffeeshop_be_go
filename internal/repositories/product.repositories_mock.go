package repositories

// import (
// 	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
// 	"github.com/stretchr/testify/mock"
// )

// type ProductRepoMock struct {
// 	mock.Mock
// }

// func (m *ProductRepoMock) CreateProduct(data *models.Product) (*models.Product, error) {
// 	args := m.Called(data)
// 	return args.Get(0).(*models.Product), args.Error(1)
// }

// func (m *ProductRepoMock) GetAllProducts(searchProductName string, minPrice int, maxPrice int, category string, sort string, page int, limit int) (*models.Products, error) {
// 	args := m.Called(searchProductName, minPrice, maxPrice, category, sort, page, limit)
// 	return args.Get(0).(*models.Products), args.Error(1)
// }

// func (m *ProductRepoMock) GetOneProduct(uuid string) (*models.Product, error) {
// 	args := m.Called(uuid)
// 	return args.Get(0).(*models.Product), args.Error(1)
// }

// func (m *ProductRepoMock) UpdateProduct(uuid string, body map[string]any) (*models.Product, error) {
// 	args := m.Called(uuid, body)
// 	return args.Get(0).(*models.Product), args.Error(1)
// }

// func (m *ProductRepoMock) DeleteProduct(uuid string) (*models.DeleteProduct, error) {
// 	args := m.Called(uuid)
// 	return args.Get(0).(*models.DeleteProduct), args.Error(1)
// }
