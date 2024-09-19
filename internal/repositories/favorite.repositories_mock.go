package repositories

import (
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/stretchr/testify/mock"
)

// AuthRepoMock adalah mock implementasi dari AuthRepoInterface
type FavoriteRepoMock struct {
	mock.Mock
}

func (m *FavoriteRepoMock) AddFavorite(userID, productID int) (*models.Favorite, error) {
	args := m.Called(userID, productID)

	if args.Get(0) != nil {
		return args.Get(0).(*models.Favorite), args.Error(1)
	}
	return nil, args.Error(1)
}

// func (m *FavoriteRepoMock) AddFavorite(userID, productID int) (*models.Favorite, error) {
// 	args := m.Called(userID, productID)
// 	return args.Get(0).(*models.Favorite), args.Error(1)
// }

func (m *FavoriteRepoMock) RemoveFavorite(userID, productID int) error {
	args := m.Called(userID, productID)
	return args.Error(0)
}

func (m *FavoriteRepoMock) GetFavorites(userID int) (*models.Favorites, error) {
	args := m.Called(userID)

	if args.Get(0) != nil {
		return args.Get(0).(*models.Favorites), args.Error(1)
	}
	return nil, args.Error(1)
}

// func (m *FavoriteRepoMock) GetFavorites(userID int) (*models.Favorites, error) {
// 	args := m.Called(userID)
// 	return args.Get(0).(*models.Favorites), args.Error(1)
// }
