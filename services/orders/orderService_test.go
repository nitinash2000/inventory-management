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

func (suite *orderServiceTestSuite) TestOrderDtosToModel_IdMissing() {
	now := time.Now()

	input := &dtos.Order{
		OrderId:     "",
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

	orderModel, _ := OrderDtosToModel(input)

	assert.NotEmpty(suite.T(), orderModel.OrderId)
}

func (suite *orderServiceTestSuite) TestOrderDtosToModel_TimedMissing() {
	input := &dtos.Order{
		OrderId:     "123",
		CustomerId:  "234",
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

	orderModel, _ := OrderDtosToModel(input)

	assert.False(suite.T(), orderModel.OrderedAt.IsZero())
}

func (suite *orderServiceTestSuite) TestCreateOrder_OrderRepoError() {
	req := &dtos.Order{
		OrderId:    "123",
		CustomerId: "234",
		Items:      []*dtos.OrderItems{},
	}

	orderModel, _ := OrderDtosToModel(req)

	suite.mockOrderRepo.EXPECT().Create(orderModel).Return(errors.New("create failed")).Times(1)

	err := suite.orderService.CreateOrder(req)

	assert.EqualError(suite.T(), err, "create failed")
}

func (suite *orderServiceTestSuite) TestCreateOrder_ItemRepoError() {
	req := &dtos.Order{
		OrderId:    "123",
		CustomerId: "234",
		Items:      []*dtos.OrderItems{},
	}

	orderModel, _ := OrderDtosToModel(req)

	suite.mockOrderRepo.EXPECT().Create(orderModel).Return(nil).Times(1)
	suite.mockOrderItemRepo.EXPECT().Create(gomock.Any()).Return(errors.New("item create error")).Times(1)

	err := suite.orderService.CreateOrder(req)

	assert.EqualError(suite.T(), err, "item create error")
}

func (suite *orderServiceTestSuite) TestUpdateOrder_OrderIdEmpty() {
	req := &dtos.Order{
		OrderId: "",
	}

	err := suite.orderService.UpdateOrder("some-id", req)

	assert.Equal(suite.T(), constants.ErrorOrderIdEmpty, err)
}

func (suite *orderServiceTestSuite) TestUpdateOrder_OrderRepoUpdateError() {
	req := &dtos.Order{
		OrderId:    "123",
		CustomerId: "234",
		Items:      []*dtos.OrderItems{},
	}

	orderModel, _ := OrderDtosToModel(req)

	suite.mockOrderRepo.EXPECT().Update("123", orderModel).Return(errors.New("update failed")).Times(1)

	err := suite.orderService.UpdateOrder("123", req)
	assert.EqualError(suite.T(), err, "update failed")
}

func (suite *orderServiceTestSuite) TestUpdateOrder_ItemRepoUpdateError() {
	req := &dtos.Order{
		OrderId:    "123",
		CustomerId: "234",
		Items:      []*dtos.OrderItems{},
	}

	orderModel, _ := OrderDtosToModel(req)

	suite.mockOrderRepo.EXPECT().Update("123", orderModel).Return(nil).Times(1)
	suite.mockOrderItemRepo.EXPECT().GetByOrder("123").Return(nil, errors.New("item update failed")).Times(1)

	err := suite.orderService.UpdateOrder("123", req)

	assert.EqualError(suite.T(), err, "item update failed")
}

func (suite *orderServiceTestSuite) TestUpdateOrder_DeleteAllError() {
	req := &dtos.Order{
		OrderId: "123",
		Items: []*dtos.OrderItems{
			{
				OrderItemId: "new-item",
				OrderId:     "123",
				ArticleId:   "1",
				Quantity:    1,
			},
		},
	}

	orderModel, _ := OrderDtosToModel(req)
	suite.mockOrderRepo.EXPECT().Update("123", orderModel).Return(nil).Times(1)

	existing := []*models.OrderItem{
		{OrderItemId: "old-item"},
	}
	suite.mockOrderItemRepo.EXPECT().GetByOrder("123").Return(existing, nil).Times(1)
	suite.mockOrderItemRepo.EXPECT().DeleteAll([]string{"old-item"}).Return(errors.New("delete failed")).Times(1)

	err := suite.orderService.UpdateOrder("123", req)
	assert.EqualError(suite.T(), err, "delete failed")
}

func (suite *orderServiceTestSuite) TestUpdateOrder_UpsertError() {
	req := &dtos.Order{
		OrderId: "123",
		Items: []*dtos.OrderItems{
			{
				OrderItemId: "item-1",
				OrderId:     "123",
				ArticleId:   "1",
				Quantity:    1,
			},
		},
	}

	orderModel, _ := OrderDtosToModel(req)

	suite.mockOrderRepo.EXPECT().Update("123", orderModel).Return(nil).Times(1)
	suite.mockOrderItemRepo.EXPECT().GetByOrder("123").Return([]*models.OrderItem{}, nil).Times(1)
	suite.mockOrderItemRepo.EXPECT().Upsert(gomock.Any()).Return(errors.New("upsert failed")).Times(1)

	err := suite.orderService.UpdateOrder("123", req)

	assert.EqualError(suite.T(), err, "upsert failed")
}

func (suite *orderServiceTestSuite) TestGetOrder_OrderItemRepoError() {
	suite.mockOrderRepo.EXPECT().Get("123").Return(&models.Order{}, nil).Times(1)
	suite.mockOrderItemRepo.EXPECT().GetByOrder("123").Return(nil, errors.New("items error")).Times(1)

	result, err := suite.orderService.GetOrder("123")

	assert.Nil(suite.T(), result)
	assert.EqualError(suite.T(), err, "items error")
}
