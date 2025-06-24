package users

import (
	"errors"
	"inventory-management/constants"
	"inventory-management/dtos"
	"inventory-management/models"
	"inventory-management/repository/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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
	suite.mockAddressRepo = mocks.NewMockAddressRepo(suite.mockCtrl)

	suite.userService = NewUserService(suite.mockUserRepo, suite.mockAddressRepo)
}

func (suite *userServiceTestSuite) TestCreateUser() {
	req := &dtos.User{
		Id:     "123",
		Name:   "John",
		Email:  "john@abc.com",
		Mobile: "12345",
		Address: dtos.Address{
			AddressId: "5",
			Line1:     "12",
			Line2:     "park street",
			City:      "chennai",
			State:     "tn",
			Country:   "in",
			ZipCode:   "600001",
		},
		Role: constants.RoleCustomer,
	}

	userModel := &models.User{
		Id:        "123",
		Name:      "John",
		Email:     "john@abc.com",
		Mobile:    "12345",
		AddressId: "5",
		Role:      constants.RoleCustomer,
	}

	addressModel := &models.Address{
		AddressId: "5",
		Line1:     "12",
		Line2:     "park street",
		City:      "chennai",
		State:     "tn",
		Country:   "in",
		ZipCode:   "600001",
	}

	suite.mockUserRepo.EXPECT().Upsert(userModel).Return(nil).Times(1)
	suite.mockAddressRepo.EXPECT().Upsert(addressModel).Return(nil).Times(1)

	err := suite.userService.CreateUser(req)
	assert.NoError(suite.T(), err)
}

func (suite *userServiceTestSuite) TestCreateUserRepoError() {
	req := &dtos.User{
		Id:     "123",
		Name:   "John",
		Email:  "john@abc.com",
		Mobile: "12345",
		Address: dtos.Address{
			AddressId: "5",
			Line1:     "12",
			Line2:     "park street",
			City:      "chennai",
			State:     "tn",
			Country:   "in",
			ZipCode:   "600001",
		},
		Role: constants.RoleCustomer,
	}

	userModel := &models.User{
		Id:        "123",
		Name:      "John",
		Email:     "john@abc.com",
		Mobile:    "12345",
		AddressId: "5",
		Role:      constants.RoleCustomer,
	}

	// addressModel := []*models.Address{
	// 	{
	// 		AddressId: "5",
	// 		Line1:     "12",
	// 		Line2:     "park street",
	// 		City:      "chennai",
	// 		State:     "tn",
	// 		Country:   "in",
	// 		ZipCode:   "600001",
	// 	},
	// }

	suite.mockUserRepo.EXPECT().Upsert(userModel).Return(errors.New("repo error")).Times(1)

	err := suite.userService.CreateUser(req)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "repo error", err.Error())
}

func (suite *userServiceTestSuite) TestGetUser() {
	expected := &dtos.User{
		Id:     "123",
		Name:   "John",
		Email:  "john@abc.com",
		Mobile: "12345",
		Address: dtos.Address{
			AddressId: "5",
			Line1:     "12",
			Line2:     "park street",
			City:      "chennai",
			State:     "tn",
			Country:   "in",
			ZipCode:   "600001",
		},
		Role: constants.RoleCustomer,
	}

	userModel := &models.User{
		Id:        "123",
		Name:      "John",
		Email:     "john@abc.com",
		Mobile:    "12345",
		AddressId: "5",
		Role:      constants.RoleCustomer,
	}

	addressModel := &models.Address{
		AddressId: "5",
		Line1:     "12",
		Line2:     "park street",
		City:      "chennai",
		State:     "tn",
		Country:   "in",
		ZipCode:   "600001",
	}

	suite.mockUserRepo.EXPECT().Get("123").Return(userModel, nil).Times(1)
	suite.mockAddressRepo.EXPECT().Get("5").Return(addressModel, nil).Times(1)

	result, err := suite.userService.GetUser("123")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, result)
}

func (suite *userServiceTestSuite) TestGetUserError() {
	expectedError := constants.ErrorNotFound

	suite.mockUserRepo.EXPECT().Get("123").Return(nil, expectedError).Times(1)

	result, err := suite.userService.GetUser("123")

	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
}

func (suite *userServiceTestSuite) TestUpdateUser() {
	req := &dtos.User{
		Id:     "123",
		Name:   "John",
		Email:  "john@abc.com",
		Mobile: "12345",
		Address: dtos.Address{
			AddressId: "5",
			Line1:     "12",
			Line2:     "park street",
			City:      "chennai",
			State:     "tn",
			Country:   "in",
			ZipCode:   "600001",
		},
		Role: constants.RoleCustomer,
	}

	userModel := &models.User{
		Id:        "123",
		Name:      "John",
		Email:     "john@abc.com",
		Mobile:    "12345",
		AddressId: "5",
		Role:      constants.RoleCustomer,
	}

	addressModel := &models.Address{
		AddressId: "5",
		Line1:     "12",
		Line2:     "park street",
		City:      "chennai",
		State:     "tn",
		Country:   "in",
		ZipCode:   "600001",
	}

	suite.mockUserRepo.EXPECT().Upsert(userModel).Return(nil).Times(1)
	suite.mockAddressRepo.EXPECT().Upsert(addressModel).Return(nil).Times(1)

	err := suite.userService.UpdateUser("123", req)
	assert.NoError(suite.T(), err)
}

func (suite *userServiceTestSuite) TestUpdateUserError() {
	req := &dtos.User{
		Id:     "123",
		Name:   "John",
		Email:  "john@abc.com",
		Mobile: "12345",
		Address: dtos.Address{
			AddressId: "5",
			Line1:     "12",
			Line2:     "park street",
			City:      "chennai",
			State:     "tn",
			Country:   "in",
			ZipCode:   "600001",
		},
		Role: constants.RoleCustomer,
	}

	userModel := &models.User{
		Id:        "123",
		Name:      "John",
		Email:     "john@abc.com",
		Mobile:    "12345",
		AddressId: "5",
		Role:      constants.RoleCustomer,
	}

	suite.mockUserRepo.EXPECT().Upsert(userModel).Return(constants.ErrorNotFound).Times(1)

	err := suite.userService.UpdateUser("123", req)
	assert.Error(suite.T(), err)
}

func (suite *userServiceTestSuite) TestDeleteUser() {
	suite.mockUserRepo.EXPECT().Delete("123").Return(nil).Times(1)

	err := suite.userService.DeleteUser("123")
	assert.NoError(suite.T(), err)
}

func (suite *userServiceTestSuite) TestDeleteUserError() {
	suite.mockUserRepo.EXPECT().Delete("123").Return(constants.ErrorNotFound).Times(1)

	err := suite.userService.DeleteUser("123")
	assert.Error(suite.T(), err)
}

func (suite *userServiceTestSuite) TestUserDtosToModel() {
	req := &dtos.User{
		Id:     "123",
		Name:   "John",
		Email:  "john@abc.com",
		Mobile: "12345",
		Address: dtos.Address{
			AddressId: "5",
			Line1:     "12",
			Line2:     "park street",
			City:      "chennai",
			State:     "tn",
			Country:   "in",
			ZipCode:   "600001",
		},
		Role: constants.RoleCustomer,
	}

	expected := &models.User{
		Id:        "123",
		Name:      "John",
		Email:     "john@abc.com",
		Mobile:    "12345",
		AddressId: "5",
		Role:      constants.RoleCustomer,
	}

	result, _ := UserDtosToModel(req)
	assert.Equal(suite.T(), expected, result)
}

func (suite *userServiceTestSuite) TestUserModelToDtos() {
	userModel := &models.User{
		Id:        "123",
		Name:      "John",
		Email:     "john@abc.com",
		Mobile:    "12345",
		AddressId: "5",
		Role:      constants.RoleCustomer,
	}

	addressModel := &models.Address{
		AddressId: "5",
		Line1:     "12",
		Line2:     "park street",
		City:      "chennai",
		State:     "tn",
		Country:   "in",
		ZipCode:   "600001",
	}

	expected := &dtos.User{
		Id:     "123",
		Name:   "John",
		Email:  "john@abc.com",
		Mobile: "12345",
		Address: dtos.Address{
			AddressId: "5",
			Line1:     "12",
			Line2:     "park street",
			City:      "chennai",
			State:     "tn",
			Country:   "in",
			ZipCode:   "600001",
		},
		Role: constants.RoleCustomer,
	}

	result := UserModelToDtos(userModel, addressModel)
	assert.Equal(suite.T(), expected, result)
}
