package repository

import (
	"errors"
	"inventory-management/models"

	"gorm.io/gorm"
)

type AddressRepo interface {
	Upsert(address *models.Address) error
	Update(addressId string, address *models.Address) error
	Get(addressId string) (*models.Address, error)
	Delete(addressId string) error
}

type addressRepo struct {
	db *gorm.DB
}

func NewAddressRepo(db *gorm.DB) AddressRepo {
	return &addressRepo{
		db: db,
	}
}

func (o *addressRepo) getTable() string {
	return "addresses"
}

func (o *addressRepo) Upsert(address *models.Address) error {
	err := o.db.Table(o.getTable()).Save(address).Error
	if err != nil {
		return err
	}

	return nil
}

func (o *addressRepo) Update(addressId string, address *models.Address) error {
	tx := o.db.Table(o.getTable()).Where("address_id = ?", addressId).UpdateColumns(address)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return errors.New("error updating address")
	}

	return nil
}

func (o *addressRepo) Get(addressId string) (*models.Address, error) {
	var result *models.Address

	err := o.db.Table(o.getTable()).Where("address_id = ?", addressId).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (o *addressRepo) Delete(addressId string) error {
	tx := o.db.Table(o.getTable()).Where("address_id = ?", addressId).Delete(&models.Address{})
	if tx.Error != nil || tx.RowsAffected == 0 {
		return errors.New("error deleting address")
	}

	return nil
}
