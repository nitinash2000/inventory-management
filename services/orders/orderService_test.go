package orders

import (
	"inventory-management/repository/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type orderServiceTestSuite struct {
	suite.Suite
	mockCtrl          *gomock.Controller
	mockOrderRepo     *mocks.MockOrderRepo
	mockOrderItemRepo *mocks.MockOrderItemRepo
	orderService      OrderService
}

func TestOrderTestSuite(t *testing.T) {
	suite.Run(t, new(orderServiceTestSuite))
}

func (suite *orderServiceTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())

	suite.mockOrderRepo = mocks.NewMockOrderRepo(suite.mockCtrl)

	suite.orderService = NewOrderService(suite.mockOrderRepo, suite.mockOrderItemRepo)
}

// func (suite *orderServiceTestSuite) TestCreateOrder() {
// 	req := &dtos.Order{
// 		OrderId:     "123",
// 		CustomerId:  "",
// 		OrderedAt:   time.Time{},
// 		TotalAmount: 0,
// 		NoOfItems:   0,
// 		Items:       []*dtos.OrderItems{},
// 	}

// 	model := &models.Order{
// 		OrderId:   "123",
// 		OrderName: "Test Order",
// 		Price:     100,
// 		Stock:     50,
// 	}

// 	suite.mockOrderRepo.EXPECT().Create(model).Return(nil).Times(1)

// 	err := suite.orderService.CreateOrder(req)
// 	assert.NoError(suite.T(), err)
// }

// func (suite *orderServiceTestSuite) TestGetOrder() {
// 	expected := &dtos.Order{
// 		OrderId:   "123",
// 		OrderName: "Test Order",
// 		Price:     100,
// 		Stock:     50,
// 	}

// 	mockOrder := &models.Order{
// 		OrderId:   "123",
// 		OrderName: "Test Order",
// 		Price:     100,
// 		Stock:     50,
// 	}

// 	suite.mockOrderRepo.EXPECT().Get("123").Return(mockOrder, nil).Times(1)

// 	result, err := suite.orderService.GetOrder("123")
// 	assert.NoError(suite.T(), err)
// 	assert.Equal(suite.T(), expected, result)
// }

// func (suite *orderServiceTestSuite) TestUpdateOrder() {
// 	req := &dtos.Order{
// 		OrderId:   "123",
// 		OrderName: "Test Order",
// 		Price:     100,
// 		Stock:     50,
// 	}

// 	model := &models.Order{
// 		OrderId:   "123",
// 		OrderName: "Test Order",
// 		Price:     100,
// 		Stock:     50,
// 	}

// 	suite.mockOrderRepo.EXPECT().Update("123", model).Return(nil).Times(1)

// 	err := suite.orderService.UpdateOrder("123", req)
// 	assert.NoError(suite.T(), err)
// }

// func (suite *orderServiceTestSuite) TestDeleteOrder() {
// 	suite.mockOrderRepo.EXPECT().Delele("123").Return(nil).Times(1)

// 	err := suite.orderService.DeleleOrder("123")
// 	assert.NoError(suite.T(), err)
// }

// func (suite *orderServiceTestSuite) TestListOrder() {
// 	expectedDto := []*dtos.Order{
// 		{
// 			OrderId:   "123",
// 			OrderName: "Test 1",
// 			Price:     100,
// 			Stock:     50,
// 		},
// 		{
// 			OrderId:   "321",
// 			OrderName: "Test 2",
// 			Price:     200,
// 			Stock:     20,
// 		},
// 	}

// 	mockModel := []*models.Order{
// 		{
// 			OrderId:   "123",
// 			OrderName: "Test 1",
// 			Price:     100,
// 			Stock:     50,
// 		},
// 		{
// 			OrderId:   "321",
// 			OrderName: "Test 2",
// 			Price:     200,
// 			Stock:     20,
// 		},
// 	}

// 	suite.mockOrderRepo.EXPECT().GetAll().Return(mockModel, nil).Times(1)

// 	result, err := suite.orderService.ListOrder()
// 	assert.NoError(suite.T(), err)
// 	assert.Equal(suite.T(), expectedDto, result)
// }

// func (suite *orderServiceTestSuite) TestUpdateOrderStock() {
// 	req := &dtos.UpdateStock{
// 		NewStock: 50,
// 	}

// 	suite.mockOrderRepo.EXPECT().UpdateOrderStock("123", int64(50)).Return(nil).Times(1)

// 	err := suite.orderService.UpdateOrderStock("123", req)
// 	assert.NoError(suite.T(), err)
// }
