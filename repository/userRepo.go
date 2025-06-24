package repository

import (
	"errors"
	"inventory-management/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	Upsert(user *models.User) error
	Update(userId string, user *models.User) error
	Get(userId string) (*models.User, error)
	Delete(userId string) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (o *userRepo) getTable() string {
	return "users"
}

func (o *userRepo) Upsert(user *models.User) error {
	err := o.db.Table(o.getTable()).Save(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (o *userRepo) Update(userId string, user *models.User) error {
	tx := o.db.Table(o.getTable()).Where("id = ?", userId).UpdateColumns(user)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return errors.New("error updating user")
	}

	return nil
}

func (o *userRepo) Get(userId string) (*models.User, error) {
	var result *models.User

	err := o.db.Table(o.getTable()).Where("id = ?", userId).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (o *userRepo) Delete(userId string) error {
	tx := o.db.Table(o.getTable()).Where("id = ?", userId).Delete(&models.User{})
	if tx.Error != nil || tx.RowsAffected == 0 {
		return errors.New("error deleting user")
	}

	return nil
}
