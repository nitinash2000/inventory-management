package repository

import (
	"inventory-management/models"

	"gorm.io/gorm"
)

type OrderItemRepo interface {
	Create(orderItem ...*models.OrderItem) error
	Update(orderItemId string, orderItem *models.OrderItem) error
	Get(orderItemId string) (*models.OrderItem, error)
	Delete(orderItemId string) error
	GetByOrder(orderId string) ([]*models.OrderItem, error)
}

type orderItemRepo struct {
	db *gorm.DB
}

func NewOrderItemRepo(db *gorm.DB) OrderItemRepo {
	return &orderItemRepo{
		db: db,
	}
}

func (o *orderItemRepo) getTable() string {
	return "orderItems"
}

func (o *orderItemRepo) Create(orderItem ...*models.OrderItem) error {
	err := o.db.Table(o.getTable()).Save(orderItem).Error
	if err != nil {
		return err
	}

	return nil
}

func (o *orderItemRepo) Update(orderItemId string, orderItem *models.OrderItem) error {
	err := o.db.Table(o.getTable()).Where("order_item_id = ?", orderItemId).UpdateColumns(orderItem).Error
	if err != nil {
		return err
	}

	return nil
}

func (o *orderItemRepo) Get(orderItemId string) (*models.OrderItem, error) {
	var result *models.OrderItem

	err := o.db.Table(o.getTable()).Where("order_item_id = ?", orderItemId).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (o *orderItemRepo) Delete(orderItemId string) error {
	err := o.db.Table(o.getTable()).Delete("order_item_id = ?", orderItemId).Error
	if err != nil {
		return err
	}

	return nil
}

func (o *orderItemRepo) GetByOrder(orderId string) ([]*models.OrderItem, error) {
	var result []*models.OrderItem

	err := o.db.Table(o.getTable()).Where("order_id = ?", orderId).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
