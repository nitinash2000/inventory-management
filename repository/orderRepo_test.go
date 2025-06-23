package repository

import (
	"inventory-management/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type OrderRepoTestSuite struct {
	suite.Suite
	db        *gorm.DB
	orderRepo OrderRepo
}

func TestOrderRepoTestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepoTestSuite))
}

func (suite *OrderRepoTestSuite) SetupTest() {
	var err error
	suite.db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		suite.T().Fatal("failed to connect to database")
	}

	err = suite.db.AutoMigrate(&models.Order{})
	if err != nil {
		suite.T().Fatal("failed to migrate database")
	}

	suite.orderRepo = NewOrderRepo(suite.db)
}

func (suite *OrderRepoTestSuite) TearDownTest() {
	sqlDB, _ := suite.db.DB()
	sqlDB.Close()
}

func (suite *OrderRepoTestSuite) TestCreateOrder() {
	order := &models.Order{
		OrderId:     "123",
		CustomerId:  "234",
		OrderedAt:   time.Now(),
		TotalAmount: 123.5,
		NoOfItems:   2,
	}

	err := suite.orderRepo.Create(order)

	assert.NoError(suite.T(), err)

	var savedOrder models.Order
	suite.db.Table("orders").Where("order_id = ?", order.OrderId).First(&savedOrder)
	assert.Equal(suite.T(), order.OrderId, savedOrder.OrderId)
	assert.Equal(suite.T(), order.TotalAmount, savedOrder.TotalAmount)
}

func (suite *OrderRepoTestSuite) TestGetOrder() {
	order := &models.Order{
		OrderId:     "123",
		CustomerId:  "234",
		OrderedAt:   time.Now(),
		TotalAmount: 123.5,
		NoOfItems:   2,
	}
	err := suite.orderRepo.Create(order)
	assert.NoError(suite.T(), err)

	result, err := suite.orderRepo.Get("123")

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), order.OrderId, result.OrderId)
	assert.Equal(suite.T(), order.TotalAmount, result.TotalAmount)
}

func (suite *OrderRepoTestSuite) TestUpdateOrder() {
	order := &models.Order{
		OrderId:     "123",
		CustomerId:  "234",
		OrderedAt:   time.Now(),
		TotalAmount: 123.5,
		NoOfItems:   2,
	}
	err := suite.orderRepo.Create(order)
	assert.NoError(suite.T(), err)

	order.TotalAmount = 230

	err = suite.orderRepo.Update(order.OrderId, order)
	assert.NoError(suite.T(), err)

	var updatedOrder models.Order
	suite.db.Table("orders").Where("order_id = ?", order.OrderId).First(&updatedOrder)
	assert.Equal(suite.T(), float64(230), updatedOrder.TotalAmount)
}

func (suite *OrderRepoTestSuite) TestDeleteOrder() {
	order := &models.Order{
		OrderId:     "123",
		CustomerId:  "234",
		OrderedAt:   time.Now(),
		TotalAmount: 123.5,
		NoOfItems:   2,
	}
	err := suite.orderRepo.Create(order)
	assert.NoError(suite.T(), err)

	err = suite.orderRepo.Delete(order.OrderId)

	assert.NoError(suite.T(), err)

	var deletedOrder models.Order
	err = suite.db.Table("orders").Where("order_id = ?", order.OrderId).First(&deletedOrder).Error
	assert.Error(suite.T(), err)
}
