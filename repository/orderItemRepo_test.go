package repository

import (
	"inventory-management/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type OrderItemRepoTestSuite struct {
	suite.Suite
	db            *gorm.DB
	orderItemRepo OrderItemRepo
}

func TestOrderItemRepoTestSuite(t *testing.T) {
	suite.Run(t, new(OrderItemRepoTestSuite))
}

func (suite *OrderItemRepoTestSuite) SetupTest() {
	var err error
	suite.db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		suite.T().Fatal("failed to connect to database")
	}

	err = suite.db.AutoMigrate(&models.OrderItem{})
	if err != nil {
		suite.T().Fatal("failed to migrate database")
	}

	suite.orderItemRepo = NewOrderItemRepo(suite.db)
}

func (suite *OrderItemRepoTestSuite) TearDownTest() {
	sqlDB, _ := suite.db.DB()
	sqlDB.Close()
}

func (suite *OrderItemRepoTestSuite) TestCreateOrderItem() {
	orderItem := &models.OrderItem{
		OrderItemId: "1234",
		OrderId:     "12",
		ArticleId:   "453",
		Quantity:    5,
	}

	err := suite.orderItemRepo.Create(orderItem)

	assert.NoError(suite.T(), err)

	var savedOrderItem models.OrderItem
	suite.db.Table("order_items").Where("order_item_id = ?", orderItem.OrderItemId).First(&savedOrderItem)
	assert.Equal(suite.T(), orderItem.OrderItemId, savedOrderItem.OrderItemId)
	assert.Equal(suite.T(), orderItem.OrderId, savedOrderItem.OrderId)
}

func (suite *OrderItemRepoTestSuite) TestCreaterderItemError() {
	orderItem := &models.OrderItem{
		OrderItemId: "1234",
		OrderId:     "12",
		Quantity:    5,
	}

	err := suite.orderItemRepo.Create(orderItem)
	assert.Error(suite.T(), err)
}

func (suite *OrderItemRepoTestSuite) TestGetOrderItem() {
	orderItem := &models.OrderItem{
		OrderItemId: "1234",
		OrderId:     "12",
		ArticleId:   "453",
		Quantity:    5,
	}
	err := suite.orderItemRepo.Create(orderItem)
	assert.NoError(suite.T(), err)

	result, err := suite.orderItemRepo.Get("1234")

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), orderItem.OrderItemId, result.OrderItemId)
}

func (suite *OrderItemRepoTestSuite) TestGetOrderItemError() {
	result, err := suite.orderItemRepo.Get("123")

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Equal(suite.T(), gorm.ErrRecordNotFound, err)
}

func (suite *OrderItemRepoTestSuite) TestUpdateOrderItem() {
	orderItem := &models.OrderItem{
		OrderItemId: "1234",
		OrderId:     "12",
		ArticleId:   "453",
		Quantity:    5,
	}
	err := suite.orderItemRepo.Create(orderItem)
	assert.NoError(suite.T(), err)

	orderItem.ArticleId = "523"

	err = suite.orderItemRepo.Update(orderItem.OrderItemId, orderItem)
	assert.NoError(suite.T(), err)

	var updatedOrderItem models.OrderItem
	suite.db.Table("orderItemes").Where("orderItem_id = ?", orderItem.OrderItemId).First(&updatedOrderItem)
}

func (suite *OrderItemRepoTestSuite) TestUpdateOrderItemError() {
	orderItem := &models.OrderItem{
		OrderItemId: "1234",
		OrderId:     "12",
		ArticleId:   "453",
		Quantity:    5,
	}

	err := suite.orderItemRepo.Update(orderItem.OrderItemId, orderItem)
	assert.Error(suite.T(), err)
}

func (suite *OrderItemRepoTestSuite) TestDeleteOrderItem() {
	orderItem := &models.OrderItem{
		OrderItemId: "1234",
		OrderId:     "12",
		ArticleId:   "453",
		Quantity:    5,
	}
	err := suite.orderItemRepo.Create(orderItem)
	assert.NoError(suite.T(), err)

	err = suite.orderItemRepo.Delete(orderItem.OrderItemId)

	assert.NoError(suite.T(), err)

	var deletedOrderItem models.OrderItem
	err = suite.db.Table("order_items").Where("order_item_id = ?", orderItem.OrderItemId).First(&deletedOrderItem).Error
	assert.Error(suite.T(), err)
}

func (suite *OrderItemRepoTestSuite) TestDeleteOrderItemError() {
	err := suite.orderItemRepo.Delete("123")

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "error deleting orderItem", err.Error())
}

func (suite *OrderItemRepoTestSuite) TestGetByOrder() {
	orderItem := &models.OrderItem{
		OrderItemId: "1234",
		OrderId:     "12",
		ArticleId:   "453",
		Quantity:    5,
	}
	err := suite.orderItemRepo.Create(orderItem)
	assert.NoError(suite.T(), err)

	result, err := suite.orderItemRepo.GetByOrder("12")
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), result)
}

func (suite *OrderItemRepoTestSuite) TestGetByOrderError() {
	result, err := suite.orderItemRepo.GetByOrder("12")

	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), result)
	assert.Equal(suite.T(), "error getting orderItems", err.Error())
}

func (suite *OrderItemRepoTestSuite) TestDeleteAllOrderItem() {
	orderItem := &models.OrderItem{
		OrderItemId: "1234",
		OrderId:     "12",
		ArticleId:   "453",
		Quantity:    5,
	}
	err := suite.orderItemRepo.Create(orderItem)
	assert.NoError(suite.T(), err)

	err = suite.orderItemRepo.DeleteAll([]string{orderItem.OrderItemId})

	assert.NoError(suite.T(), err)

	var deletedOrderItem models.OrderItem
	err = suite.db.Table("order_items").Where("order_item_id = ?", orderItem.OrderItemId).First(&deletedOrderItem).Error
	assert.Error(suite.T(), err)
}

func (suite *OrderItemRepoTestSuite) TestDeleteAllOrderItemError() {
	err := suite.orderItemRepo.DeleteAll([]string{"123"})

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "error deleting orderItem", err.Error())
}

func (suite *OrderItemRepoTestSuite) TestUpsertOrderItem() {
	orderItem := &models.OrderItem{
		OrderItemId: "upsert-123",
		OrderId:     "order-001",
		ArticleId:   "article-123",
		Quantity:    10,
	}

	err := suite.orderItemRepo.Upsert(orderItem)
	assert.NoError(suite.T(), err)

	var savedOrderItem models.OrderItem
	err = suite.db.Table("order_items").Where("order_item_id = ?", orderItem.OrderItemId).First(&savedOrderItem).Error
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), orderItem.OrderItemId, savedOrderItem.OrderItemId)
	assert.Equal(suite.T(), orderItem.ArticleId, savedOrderItem.ArticleId)
	assert.Equal(suite.T(), orderItem.Quantity, savedOrderItem.Quantity)
}

func (suite *OrderItemRepoTestSuite) TestUpsertOrderItemError() {
	orderItem := &models.OrderItem{
		OrderItemId: "upsert-err-001",
		OrderId:     "order-001",
		Quantity:    2,
	}

	err := suite.orderItemRepo.Upsert(orderItem)
	assert.Error(suite.T(), err)
}
