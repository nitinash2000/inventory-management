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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type orderHandlerTestSuite struct {
	suite.Suite
	mockCtrl         *gomock.Controller
	mockOrderService *mocks.MockOrderService
	orderHandler     *orderHandler
}

func TestOrderHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(orderHandlerTestSuite))
}

func (suite *orderHandlerTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())

	suite.mockOrderService = mocks.NewMockOrderService(suite.mockCtrl)

	suite.orderHandler = NewOrderHandler(suite.mockOrderService)
}

func (suite *orderHandlerTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *orderHandlerTestSuite) TestGetOrder() {
	now := time.Now()

	expected := &dtos.Order{
		OrderId:     "123",
		CustomerId:  "234",
		OrderedAt:   now,
		TotalAmount: 200,
		NoOfItems:   2,
		Items: []*dtos.OrderItems{
			{
				OrderItemId: "",
				OrderId:     "123",
				ArticleId:   "1",
				Quantity:    1,
			},
			{
				OrderItemId: "",
				OrderId:     "123",
				ArticleId:   "2",
				Quantity:    1,
			},
		},
	}

	suite.mockOrderService.EXPECT().GetOrder("123").Return(expected, nil).Times(1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodGet, "/orders/123", nil)

	suite.orderHandler.GetOrder(c)

	var result *dtos.Order

	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(suite.T(), err)

	assert.EqualValues(suite.T(), expected.CustomerId, result.CustomerId)
	assert.Equal(suite.T(), expected.TotalAmount, result.TotalAmount)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *orderHandlerTestSuite) TestGetOrderError() {
	suite.mockOrderService.EXPECT().GetOrder("123").Return(nil, constants.ErrorNotFound).Times(1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodGet, "/orders/123", nil)

	suite.orderHandler.GetOrder(c)

	var result *dtos.Order

	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), result)
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

func (suite *orderHandlerTestSuite) TestCreateOrder() {
	now := time.Now()

	req := &dtos.Order{
		OrderId:     "123",
		CustomerId:  "234",
		OrderedAt:   now,
		TotalAmount: 200,
		NoOfItems:   2,
		Items: []*dtos.OrderItems{
			{
				OrderItemId: "",
				OrderId:     "123",
				ArticleId:   "1",
				Quantity:    1,
			},
			{
				OrderItemId: "",
				OrderId:     "123",
				ArticleId:   "2",
				Quantity:    1,
			},
		},
	}

	body, _ := json.Marshal(req)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.mockOrderService.EXPECT().CreateOrder(gomock.AssignableToTypeOf(&dtos.Order{})).Return(nil).Times(1)

	suite.orderHandler.CreateOrder(c)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *orderHandlerTestSuite) TestCreateOrderError() {
	now := time.Now()

	req := &dtos.Order{
		OrderId:     "123",
		CustomerId:  "234",
		OrderedAt:   now,
		TotalAmount: 200,
		NoOfItems:   2,
		Items: []*dtos.OrderItems{
			{
				OrderItemId: "",
				OrderId:     "123",
				ArticleId:   "1",
				Quantity:    1,
			},
			{
				OrderItemId: "",
				OrderId:     "123",
				ArticleId:   "2",
				Quantity:    1,
			},
		},
	}

	body, _ := json.Marshal(req)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.mockOrderService.EXPECT().CreateOrder(gomock.AssignableToTypeOf(&dtos.Order{})).Return(constants.ErrorRecordExists).Times(1)

	suite.orderHandler.CreateOrder(c)
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

func (suite *orderHandlerTestSuite) TestCreateOrder_BadRequest() {
	invalidJSON := `{"order_id": 123, "orderName": "Test Order", "price": "not_a_number", "stock": "50"}`

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/orders", bytes.NewReader([]byte(invalidJSON)))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.orderHandler.CreateOrder(c)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *orderHandlerTestSuite) TestDeleteOrder() {
	suite.mockOrderService.EXPECT().DeleteOrder("123").Return(nil).Times(1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodDelete, "/orders/123", nil)

	suite.orderHandler.DeleteOrder(c)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *orderHandlerTestSuite) TestDeleteOrderError() {
	suite.mockOrderService.EXPECT().DeleteOrder("123").Return(constants.ErrorNotFound).Times(1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodDelete, "/orders/123", nil)

	suite.orderHandler.DeleteOrder(c)
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

func (suite *orderHandlerTestSuite) TestUpdateOrder() {
	now := time.Now()

	req := &dtos.Order{
		OrderId:     "123",
		CustomerId:  "234",
		OrderedAt:   now,
		TotalAmount: 200,
		NoOfItems:   2,
		Items: []*dtos.OrderItems{
			{
				OrderItemId: "",
				OrderId:     "123",
				ArticleId:   "1",
				Quantity:    1,
			},
			{
				OrderItemId: "",
				OrderId:     "123",
				ArticleId:   "2",
				Quantity:    1,
			},
		},
	}

	body, _ := json.Marshal(req)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodPut, "/orders/123", bytes.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.mockOrderService.EXPECT().UpdateOrder("123", gomock.Any()).Return(nil).Times(1)

	suite.orderHandler.UpdateOrder(c)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *orderHandlerTestSuite) TestUpdateOrderError() {
	now := time.Now()

	req := &dtos.Order{
		OrderId:     "123",
		CustomerId:  "234",
		OrderedAt:   now,
		TotalAmount: 200,
		NoOfItems:   2,
		Items: []*dtos.OrderItems{
			{
				OrderItemId: "",
				OrderId:     "123",
				ArticleId:   "1",
				Quantity:    1,
			},
			{
				OrderItemId: "",
				OrderId:     "123",
				ArticleId:   "2",
				Quantity:    1,
			},
		},
	}

	body, _ := json.Marshal(req)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodPut, "/orders/123", bytes.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.mockOrderService.EXPECT().UpdateOrder("123", gomock.Any()).Return(constants.ErrorNotFound).Times(1)

	suite.orderHandler.UpdateOrder(c)
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

func (suite *orderHandlerTestSuite) TestUpdateOrderBadRequest() {
	invalidJSON := `{"order_id": 123, "orderName": "Test Order", "price": "not_a_number", "stock": "50"}`

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodPut, "/orders/123", bytes.NewReader([]byte(invalidJSON)))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.orderHandler.UpdateOrder(c)
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}
