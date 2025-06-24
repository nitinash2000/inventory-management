package users

import (
	"inventory-management/dtos"
	"inventory-management/models"
	"inventory-management/repository"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(req *dtos.User) error
	UpdateUser(id string, req *dtos.User) error
	GetUser(userId string) (*dtos.User, error)
	DeleteUser(userId string) error
}

type userService struct {
	userRepo    repository.UserRepo
	addressRepo repository.AddressRepo
}

func NewUserService(userRepo repository.UserRepo, addressRepo repository.AddressRepo) UserService {
	return &userService{
		userRepo:    userRepo,
		addressRepo: addressRepo,
	}
}

func (o *userService) CreateUser(req *dtos.User) error {
	userModel, addressModel := UserDtosToModel(req)

	err := o.userRepo.Upsert(userModel)
	if err != nil {
		return err
	}

	err = o.addressRepo.Upsert(addressModel)
	if err != nil {
		return err
	}

	return nil
}

func (o *userService) UpdateUser(id string, req *dtos.User) error {
	userModel, addressModel := UserDtosToModel(req)

	err := o.userRepo.Upsert(userModel)
	if err != nil {
		return err
	}

	err = o.addressRepo.Upsert(addressModel)
	if err != nil {
		return err
	}

	return nil
}

func (o *userService) GetUser(userId string) (*dtos.User, error) {
	user, err := o.userRepo.Get(userId)
	if err != nil {
		return nil, err
	}

	address, err := o.addressRepo.Get(user.AddressId)
	if err != nil {
		return nil, err
	}

	result := UserModelToDtos(user, address)

	return result, nil
}

func (o *userService) DeleteUser(userId string) error {
	err := o.userRepo.Delete(userId)
	if err != nil {
		return err
	}

	return nil
}

func UserModelToDtos(m *models.User, a *models.Address) *dtos.User {
	user := &dtos.User{
		Id:     m.Id,
		Name:   m.Name,
		Email:  m.Email,
		Mobile: m.Mobile,
		Address: dtos.Address{
			AddressId: a.AddressId,
			Line1:     a.Line1,
			Line2:     a.Line2,
			City:      a.City,
			State:     a.State,
			Country:   a.Country,
			ZipCode:   a.ZipCode,
		},
		Role: m.Role,
	}

	return user
}

func UserDtosToModel(m *dtos.User) (*models.User, *models.Address) {
	if m.Id == "" {
		m.Id = uuid.NewString()
	}

	addressId := m.Address.AddressId
	if addressId == "" {
		addressId = uuid.NewString()
	}

	userModel := &models.User{
		Id:        m.Id,
		Name:      m.Name,
		Email:     m.Email,
		Mobile:    m.Mobile,
		AddressId: addressId,
		Role:      m.Role,
	}

	addressModel := &models.Address{
		AddressId: addressId,
		Line1:     m.Address.Line1,
		Line2:     m.Address.Line2,
		City:      m.Address.City,
		State:     m.Address.State,
		Country:   m.Address.Country,
		ZipCode:   m.Address.ZipCode,
	}

	return userModel, addressModel
}
