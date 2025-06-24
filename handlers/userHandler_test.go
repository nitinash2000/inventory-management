package handlers

import (
	"bytes"
	"encoding/json"
	"inventory-management/constants"
	"inventory-management/dtos"
	"inventory-management/services/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type userHandlerTestSuite struct {
	suite.Suite
	mockCtrl        *gomock.Controller
	mockUserService *mocks.MockUserService
	userHandler     *userHandler
}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(userHandlerTestSuite))
}

func (suite *userHandlerTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())

	suite.mockUserService = mocks.NewMockUserService(suite.mockCtrl)

	suite.userHandler = NewUserHandler(suite.mockUserService)
}

func (suite *userHandlerTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *userHandlerTestSuite) TestGetUser() {
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

	suite.mockUserService.EXPECT().GetUser("123").Return(expected, nil).Times(1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodGet, "/users/123", nil)

	suite.userHandler.GetUser(c)

	var result *dtos.User

	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expected.Name, result.Name)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *userHandlerTestSuite) TestGetUserError() {
	suite.mockUserService.EXPECT().GetUser("123").Return(nil, constants.ErrorNotFound).Times(1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodGet, "/users/123", nil)

	suite.userHandler.GetUser(c)

	var result *dtos.User

	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), result)
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

func (suite *userHandlerTestSuite) TestCreateUser() {
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

	body, _ := json.Marshal(req)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.mockUserService.EXPECT().CreateUser(gomock.AssignableToTypeOf(&dtos.User{})).Return(nil).Times(1)

	suite.userHandler.CreateUser(c)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *userHandlerTestSuite) TestCreateUserError() {
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

	body, _ := json.Marshal(req)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.mockUserService.EXPECT().CreateUser(gomock.AssignableToTypeOf(&dtos.User{})).Return(constants.ErrorRecordExists).Times(1)

	suite.userHandler.CreateUser(c)
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

func (suite *userHandlerTestSuite) TestCreateUser_BadRequest() {
	invalidJSON := `{"id": 123, "userName": "Test User", "price": "not_a_number", "stock": "50"}`

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(invalidJSON)))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.userHandler.CreateUser(c)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *userHandlerTestSuite) TestDeleteUser() {
	suite.mockUserService.EXPECT().DeleteUser("123").Return(nil).Times(1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodDelete, "/users/123", nil)

	suite.userHandler.DeleteUser(c)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *userHandlerTestSuite) TestDeleteUserError() {
	suite.mockUserService.EXPECT().DeleteUser("123").Return(constants.ErrorNotFound).Times(1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodDelete, "/users/123", nil)

	suite.userHandler.DeleteUser(c)
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

func (suite *userHandlerTestSuite) TestUpdateUser() {
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

	body, _ := json.Marshal(req)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodPut, "/users/123", bytes.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.mockUserService.EXPECT().UpdateUser("123", gomock.Any()).Return(nil).Times(1)

	suite.userHandler.UpdateUser(c)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *userHandlerTestSuite) TestUpdateUserError() {
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

	body, _ := json.Marshal(req)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodPut, "/users/123", bytes.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.mockUserService.EXPECT().UpdateUser("123", gomock.Any()).Return(constants.ErrorNotFound).Times(1)

	suite.userHandler.UpdateUser(c)
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

func (suite *userHandlerTestSuite) TestUpdateUserBadRequest() {
	invalidJSON := `{"id": 123, "userName": "Test User", "price": "not_a_number", "stock": "50"}`

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodPut, "/users/123", bytes.NewReader([]byte(invalidJSON)))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.userHandler.UpdateUser(c)
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}
