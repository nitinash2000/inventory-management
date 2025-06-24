package orders

import (
	"errors"
	"inventory-management/constants"
	"inventory-management/dtos"
	"inventory-management/models"
	"inventory-management/repository/mocks"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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
	suite.mockOrderItemRepo = mocks.NewMockOrderItemRepo(suite.mockCtrl)

	suite.orderService = NewOrderService(suite.mockOrderRepo, suite.mockOrderItemRepo)
}

func (suite *orderServiceTestSuite) TestCreateOrder() {
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

	orderModel := &models.Order{
		OrderId:     "123",
		CustomerId:  "234",
		OrderedAt:   now,
		TotalAmount: 200,
		NoOfItems:   2,
	}

	// itemsModel := []*models.OrderItem{
	// 	{
	// 		OrderItemId: "",
	// 		OrderId:     "123",
	// 		ArticleId:   "1",
	// 		Quantity:    1,
	// 	},
	// 	{
	// 		OrderItemId: "",
	// 		OrderId:     "123",
	// 		ArticleId:   "2",
	// 		Quantity:    1,
	// 	},
	// }

	suite.mockOrderRepo.EXPECT().Create(orderModel).Return(nil).Times(1)
	suite.mockOrderItemRepo.EXPECT().Create(gomock.Any()).Return(nil).Times(1)

	err := suite.orderService.CreateOrder(req)
	assert.NoError(suite.T(), err)
}

func (suite *orderServiceTestSuite) TestCreateOrderRepoError() {
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

	model := &models.Order{
		OrderId:     "123",
		CustomerId:  "234",
		OrderedAt:   now,
		TotalAmount: 200,
		NoOfItems:   2,
	}

	suite.mockOrderRepo.EXPECT().Create(model).Return(errors.New("repo error")).Times(1)

	err := suite.orderService.CreateOrder(req)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "repo error", err.Error())
}

func (suite *orderServiceTestSuite) TestGetOrder() {
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

	mockOrder := &models.Order{
		OrderId:     "123",
		CustomerId:  "234",
		OrderedAt:   now,
		TotalAmount: 200,
		NoOfItems:   2,
	}

	itemsMock := []*models.OrderItem{
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
	}

	suite.mockOrderRepo.EXPECT().Get("123").Return(mockOrder, nil).Times(1)
	suite.mockOrderItemRepo.EXPECT().GetByOrder("123").Return(itemsMock, nil).Times(1)

	result, err := suite.orderService.GetOrder("123")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, result)
}

func (suite *orderServiceTestSuite) TestGetOrderError() {
	expectedError := constants.ErrorNotFound

	suite.mockOrderRepo.EXPECT().Get("123").Return(nil, expectedError).Times(1)

	result, err := suite.orderService.GetOrder("123")

	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
}

func (suite *orderServiceTestSuite) TestUpdateOrder() {
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

	model := &models.Order{
		OrderId:     "123",
		CustomerId:  "234",
		OrderedAt:   now,
		TotalAmount: 200,
		NoOfItems:   2,
	}

	suite.mockOrderRepo.EXPECT().Update("123", model).Return(nil).Times(1)
	suite.mockOrderItemRepo.EXPECT().GetByOrder("123").Return(nil, nil).Times(1)
	suite.mockOrderItemRepo.EXPECT().Upsert(gomock.AssignableToTypeOf([]*models.OrderItem{})).Return(nil).Times(1)

	err := suite.orderService.UpdateOrder("123", req)
	assert.NoError(suite.T(), err)
}

func (suite *orderServiceTestSuite) TestUpdateOrderError() {
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

	model := &models.Order{
		OrderId:     "123",
		CustomerId:  "234",
		OrderedAt:   now,
		TotalAmount: 200,
		NoOfItems:   2,
	}

	suite.mockOrderRepo.EXPECT().Update("123", model).Return(constants.ErrorNotFound).Times(1)

	err := suite.orderService.UpdateOrder("123", req)
	assert.Error(suite.T(), err)
}

func (suite *orderServiceTestSuite) TestDeleteOrder() {
	suite.mockOrderRepo.EXPECT().Delete("123").Return(nil).Times(1)

	err := suite.orderService.DeleteOrder("123")
	assert.NoError(suite.T(), err)
}

func (suite *orderServiceTestSuite) TestDeleteOrderError() {
	suite.mockOrderRepo.EXPECT().Delete("123").Return(constants.ErrorNotFound).Times(1)

	err := suite.orderService.DeleteOrder("123")
	assert.Error(suite.T(), err)
}

func (suite *orderServiceTestSuite) TestOrderDtosToModel() {
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

	expected := &models.Order{
		OrderId:     "123",
		CustomerId:  "234",
		OrderedAt:   now,
		TotalAmount: 200,
		NoOfItems:   2,
	}

	result, _ := OrderDtosToModel(req)
	assert.Equal(suite.T(), expected, result)
}

func (suite *orderServiceTestSuite) TestOrderModelToDtos() {
	now := time.Now()

	orderModel := &models.Order{
		OrderId:     "123",
		CustomerId:  "234",
		OrderedAt:   now,
		TotalAmount: 200,
		NoOfItems:   2,
	}

	itemsModel := []*models.OrderItem{
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
	}

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

	result := OrderModelToDtos(orderModel, itemsModel)
	assert.Equal(suite.T(), expected, result)
}
