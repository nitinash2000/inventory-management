package users

import (
	"inventory-management/repository/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type userServiceTestSuite struct {
	suite.Suite
	mockCtrl        *gomock.Controller
	mockUserRepo    *mocks.MockUserRepo
	mockAddressRepo *mocks.MockAddressRepo
	userService     UserService
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(userServiceTestSuite))
}

func (suite *userServiceTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())

	suite.mockUserRepo = mocks.NewMockUserRepo(suite.mockCtrl)

	suite.userService = NewUserService(suite.mockUserRepo, suite.mockAddressRepo)
}

// func (suite *userServiceTestSuite) TestCreateUser() {
// 	req := &dtos.User{
// 		UserId:   "123",
// 		UserName: "Test User",
// 		Price:    100,
// 		Stock:    50,
// 	}

// 	model := &models.User{
// 		UserId:   "123",
// 		UserName: "Test User",
// 		Price:    100,
// 		Stock:    50,
// 	}

// 	suite.mockUserRepo.EXPECT().Create(model).Return(nil).Times(1)

// 	err := suite.userService.CreateUser(req)
// 	assert.NoError(suite.T(), err)
// }

// func (suite *userServiceTestSuite) TestGetUser() {
// 	expected := &dtos.User{
// 		UserId:   "123",
// 		UserName: "Test User",
// 		Price:    100,
// 		Stock:    50,
// 	}

// 	mockUser := &models.User{
// 		UserId:   "123",
// 		UserName: "Test User",
// 		Price:    100,
// 		Stock:    50,
// 	}

// 	suite.mockUserRepo.EXPECT().Get("123").Return(mockUser, nil).Times(1)

// 	result, err := suite.userService.GetUser("123")
// 	assert.NoError(suite.T(), err)
// 	assert.Equal(suite.T(), expected, result)
// }

// func (suite *userServiceTestSuite) TestUpdateUser() {
// 	req := &dtos.User{
// 		UserId:   "123",
// 		UserName: "Test User",
// 		Price:    100,
// 		Stock:    50,
// 	}

// 	model := &models.User{
// 		UserId:   "123",
// 		UserName: "Test User",
// 		Price:    100,
// 		Stock:    50,
// 	}

// 	suite.mockUserRepo.EXPECT().Update("123", model).Return(nil).Times(1)

// 	err := suite.userService.UpdateUser("123", req)
// 	assert.NoError(suite.T(), err)
// }

// func (suite *userServiceTestSuite) TestDeleteUser() {
// 	suite.mockUserRepo.EXPECT().Delele("123").Return(nil).Times(1)

// 	err := suite.userService.DeleleUser("123")
// 	assert.NoError(suite.T(), err)
// }

// func (suite *userServiceTestSuite) TestListUser() {
// 	expectedDto := []*dtos.User{
// 		{
// 			UserId:   "123",
// 			UserName: "Test 1",
// 			Price:    100,
// 			Stock:    50,
// 		},
// 		{
// 			UserId:   "321",
// 			UserName: "Test 2",
// 			Price:    200,
// 			Stock:    20,
// 		},
// 	}

// 	mockModel := []*models.User{
// 		{
// 			UserId:   "123",
// 			UserName: "Test 1",
// 			Price:    100,
// 			Stock:    50,
// 		},
// 		{
// 			UserId:   "321",
// 			UserName: "Test 2",
// 			Price:    200,
// 			Stock:    20,
// 		},
// 	}

// 	suite.mockUserRepo.EXPECT().GetAll().Return(mockModel, nil).Times(1)

// 	result, err := suite.userService.ListUser()
// 	assert.NoError(suite.T(), err)
// 	assert.Equal(suite.T(), expectedDto, result)
// }

// func (suite *userServiceTestSuite) TestUpdateUserStock() {
// 	req := &dtos.UpdateStock{
// 		NewStock: 50,
// 	}

// 	suite.mockUserRepo.EXPECT().UpdateUserStock("123", int64(50)).Return(nil).Times(1)

// 	err := suite.userService.UpdateUserStock("123", req)
// 	assert.NoError(suite.T(), err)
// }
