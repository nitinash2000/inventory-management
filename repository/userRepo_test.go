package repository

import (
	"inventory-management/constants"
	"inventory-management/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserRepoTestSuite struct {
	suite.Suite
	db       *gorm.DB
	userRepo UserRepo
}

func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}

func (suite *UserRepoTestSuite) SetupTest() {
	var err error
	suite.db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		suite.T().Fatal("failed to connect to database")
	}

	err = suite.db.AutoMigrate(&models.User{})
	if err != nil {
		suite.T().Fatal("failed to migrate database")
	}

	suite.userRepo = NewUserRepo(suite.db)
}

func (suite *UserRepoTestSuite) TearDownTest() {
	sqlDB, _ := suite.db.DB()
	sqlDB.Close()
}

func (suite *UserRepoTestSuite) TestCreateUser() {
	user := &models.User{
		Id:        "250",
		Name:      "John",
		Email:     "john@abc.com",
		Mobile:    "12345",
		AddressId: "25",
		Role:      constants.RoleCustomer,
	}

	err := suite.userRepo.Upsert(user)

	assert.NoError(suite.T(), err)

	var savedUser models.User
	suite.db.Table("users").Where("id = ?", user.Id).First(&savedUser)
	assert.Equal(suite.T(), user.Name, savedUser.Name)
	assert.Equal(suite.T(), user.Mobile, savedUser.Mobile)
}

func (suite *UserRepoTestSuite) TestCreaterderItemError() {
	user := &models.User{
		Id:        "250",
		Name:      "",
		Email:     "john@abc.com",
		Mobile:    "12345",
		AddressId: "25",
		Role:      constants.RoleCustomer,
	}

	err := suite.userRepo.Upsert(user)
	assert.Error(suite.T(), err)
}

func (suite *UserRepoTestSuite) TestGetUser() {
	user := &models.User{
		Id:        "250",
		Name:      "John",
		Email:     "john@abc.com",
		Mobile:    "12345",
		AddressId: "25",
		Role:      constants.RoleCustomer,
	}
	err := suite.userRepo.Upsert(user)
	assert.NoError(suite.T(), err)

	result, err := suite.userRepo.Get("250")

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.Email, result.Email)
}

func (suite *UserRepoTestSuite) TestGetUserError() {
	result, err := suite.userRepo.Get("123")

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Equal(suite.T(), gorm.ErrRecordNotFound, err)
}

func (suite *UserRepoTestSuite) TestUpdateUser() {
	user := &models.User{
		Id:        "250",
		Name:      "John",
		Email:     "john@abc.com",
		Mobile:    "12345",
		AddressId: "25",
		Role:      constants.RoleCustomer,
	}
	err := suite.userRepo.Upsert(user)
	assert.NoError(suite.T(), err)

	user.Name = "Joe"

	err = suite.userRepo.Update(user.Id, user)
	assert.NoError(suite.T(), err)

	var updatedUser models.User
	suite.db.Table("users").Where("id = ?", user.Id).First(&updatedUser)
}

func (suite *UserRepoTestSuite) TestUpdateUserError() {
	user := &models.User{
		Id:        "250",
		Name:      "John",
		Email:     "john@abc.com",
		Mobile:    "12345",
		AddressId: "25",
		Role:      constants.RoleCustomer,
	}

	err := suite.userRepo.Update(user.Id, user)
	assert.Error(suite.T(), err)
}

func (suite *UserRepoTestSuite) TestDeleteUser() {
	user := &models.User{
		Id:        "250",
		Name:      "John",
		Email:     "john@abc.com",
		Mobile:    "12345",
		AddressId: "25",
		Role:      constants.RoleCustomer,
	}
	err := suite.userRepo.Upsert(user)
	assert.NoError(suite.T(), err)

	err = suite.userRepo.Delete(user.Id)

	assert.NoError(suite.T(), err)

	var deletedUser models.User
	err = suite.db.Table("users").Where("id = ?", user.Id).First(&deletedUser).Error
	assert.Error(suite.T(), err)
}

func (suite *UserRepoTestSuite) TestDeleteUserError() {
	err := suite.userRepo.Delete("123")

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "error deleting user", err.Error())
}
