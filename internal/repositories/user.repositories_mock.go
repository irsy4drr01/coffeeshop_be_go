package repositories

// import (
// 	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
// 	"github.com/stretchr/testify/mock"
// )

// type UserRepoMock struct {
// 	mock.Mock
// }

// func (m *UserRepoMock) GetAllUser(searchUserName string, sort string, page int, limit int) (*models.Users, error) {
// 	args := m.Called(searchUserName, sort, page, limit)
// 	if args.Get(0) != nil {
// 		return args.Get(0).(*models.Users), args.Error(1)
// 	}
// 	return nil, args.Error(1)
// }

// func (m *UserRepoMock) GetOneUser(uuid string) (*models.User, error) {
// 	args := m.Called(uuid)
// 	if args.Get(0) != nil {
// 		return args.Get(0).(*models.User), args.Error(1)
// 	}
// 	return nil, args.Error(1)
// }

// func (m *UserRepoMock) UpdateUser(uuid string, body map[string]any) (*models.User, error) {
// 	args := m.Called(uuid, body)
// 	if args.Get(0) != nil {
// 		return args.Get(0).(*models.User), args.Error(1)
// 	}
// 	return nil, args.Error(1)
// }

// func (m *UserRepoMock) DeleteUser(uuid string) (*models.User, error) {
// 	args := m.Called(uuid)
// 	if args.Get(0) != nil {
// 		return args.Get(0).(*models.User), args.Error(1)
// 	}
// 	return nil, args.Error(1)
// }
