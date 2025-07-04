package repository

import (
	"errors"
	"inventory-management/models"

	"gorm.io/gorm"
)

type OrderRepo interface {
	Create(order *models.Order) error
	Update(orderId string, order *models.Order) error
	Get(orderId string) (*models.Order, error)
	Delete(orderId string) error
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return &orderRepo{
		db: db,
	}
}

func (o *orderRepo) getTable() string {
	return "orders"
}

func (o *orderRepo) Create(order *models.Order) error {
	err := o.db.Table(o.getTable()).Save(order).Error
	if err != nil {
		return err
	}

	return nil
}

func (o *orderRepo) Update(orderId string, order *models.Order) error {
	tx := o.db.Table(o.getTable()).Where("order_id = ?", orderId).UpdateColumns(order)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return errors.New("error updating order")
	}

	return nil
}

func (o *orderRepo) Get(orderId string) (*models.Order, error) {
	var result *models.Order

	err := o.db.Table(o.getTable()).Where("order_id = ?", orderId).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (o *orderRepo) Delete(orderId string) error {
	tx := o.db.Table(o.getTable()).Where("order_id = ?", orderId).Delete(&models.Order{})
	if tx.Error != nil || tx.RowsAffected == 0 {
		return errors.New("error deleting order")
	}

	return nil
}
