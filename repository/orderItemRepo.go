package repository

import (
	"errors"
	"inventory-management/models"

	"gorm.io/gorm"
)

type OrderItemRepo interface {
	Create(orderItem ...*models.OrderItem) error
	Update(orderItemId string, orderItem *models.OrderItem) error
	Get(orderItemId string) (*models.OrderItem, error)
	Delete(orderItemId string) error
	GetByOrder(orderId string) ([]*models.OrderItem, error)
	Upsert(orderItems ...*models.OrderItem) error
	DeleteAll(orderItemIds []string) error
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
	return "order_items"
}

func (o *orderItemRepo) Create(orderItem ...*models.OrderItem) error {
	err := o.db.Table(o.getTable()).Save(orderItem).Error
	if err != nil {
		return err
	}

	return nil
}

func (o *orderItemRepo) Update(orderItemId string, orderItem *models.OrderItem) error {
	tx := o.db.Table(o.getTable()).Where("order_item_id = ?", orderItemId).UpdateColumns(orderItem)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return errors.New("error updating orderItem")
	}

	return nil
}

func (o *orderItemRepo) Upsert(orderItems ...*models.OrderItem) error {
	err := o.db.Table(o.getTable()).Save(&orderItems).Error
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
	tx := o.db.Table(o.getTable()).Where("order_item_id = ?", orderItemId).Delete(&models.OrderItem{})
	if tx.Error != nil || tx.RowsAffected == 0 {
		return errors.New("error deleting orderItem")
	}

	return nil
}

func (o *orderItemRepo) GetByOrder(orderId string) ([]*models.OrderItem, error) {
	var result []*models.OrderItem

	err := o.db.Table(o.getTable()).Where("order_id = ?", orderId).Find(&result).Error
	if err != nil || len(result) == 0 {
		return nil, errors.New("error getting orderItems")
	}

	return result, nil
}

func (o *orderItemRepo) DeleteAll(orderItemIds []string) error {
	tx := o.db.Table(o.getTable()).Where(`order_item_id IN (?)`, orderItemIds).Delete(&models.OrderItem{})
	if tx.Error != nil || tx.RowsAffected == 0 {
		return errors.New("error deleting orderItem")
	}

	return nil
}
