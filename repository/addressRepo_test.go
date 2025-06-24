package repository

import (
	"inventory-management/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type AddressRepoTestSuite struct {
	suite.Suite
	db          *gorm.DB
	addressRepo AddressRepo
}

func TestAddressRepoTestSuite(t *testing.T) {
	suite.Run(t, new(AddressRepoTestSuite))
}

func (suite *AddressRepoTestSuite) SetupTest() {
	var err error
	suite.db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		suite.T().Fatal("failed to connect to database")
	}

	err = suite.db.AutoMigrate(&models.Address{})
	if err != nil {
		suite.T().Fatal("failed to migrate database")
	}

	suite.addressRepo = NewAddressRepo(suite.db)
}

func (suite *AddressRepoTestSuite) TearDownTest() {
	sqlDB, _ := suite.db.DB()
	sqlDB.Close()
}

func (suite *AddressRepoTestSuite) TestUpsertAddress() {
	address := &models.Address{
		AddressId: "123",
		Line1:     "123 Main St",
		Line2:     "Apt 4B",
		City:      "Springfield",
		State:     "IL",
		Country:   "USA",
		ZipCode:   "62704",
	}

	err := suite.addressRepo.Upsert(address)

	assert.NoError(suite.T(), err)

	var savedAddress models.Address
	suite.db.Table("addresses").Where("address_id = ?", address.AddressId).First(&savedAddress)
	assert.Equal(suite.T(), address.AddressId, savedAddress.AddressId)
	assert.Equal(suite.T(), address.Line1, savedAddress.Line1)
}

func (suite *AddressRepoTestSuite) TestUpsertAddressError() {
	address := &models.Address{
		AddressId: "123",
		Line1:     "123 Main St",
		Line2:     "Apt 4B",
		City:      "Springfield",
		State:     "IL",
		ZipCode:   "62704",
	}

	err := suite.addressRepo.Upsert(address)
	assert.Error(suite.T(), err)
}

func (suite *AddressRepoTestSuite) TestGetAddress() {
	address := &models.Address{
		AddressId: "123",
		Line1:     "123 Main St",
		Line2:     "Apt 4B",
		City:      "Springfield",
		State:     "IL",
		Country:   "USA",
		ZipCode:   "62704",
	}
	err := suite.addressRepo.Upsert(address)
	assert.NoError(suite.T(), err)

	result, err := suite.addressRepo.Get("123")

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), address.AddressId, result.AddressId)
	assert.Equal(suite.T(), address.Line1, result.Line1)
}

func (suite *AddressRepoTestSuite) TestGetAddressError() {
	result, err := suite.addressRepo.Get("123")

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Equal(suite.T(), gorm.ErrRecordNotFound, err)
}

func (suite *AddressRepoTestSuite) TestUpdateAddress() {
	address := &models.Address{
		AddressId: "123",
		Line1:     "123 Main St",
		Line2:     "Apt 4B",
		City:      "Springfield",
		State:     "IL",
		Country:   "USA",
		ZipCode:   "62704",
	}
	err := suite.addressRepo.Upsert(address)
	assert.NoError(suite.T(), err)

	address.Line1 = "456 Elm St"

	err = suite.addressRepo.Update(address.AddressId, address)
	assert.NoError(suite.T(), err)

	var updatedAddress models.Address
	suite.db.Table("addresses").Where("address_id = ?", address.AddressId).First(&updatedAddress)
	assert.Equal(suite.T(), "456 Elm St", updatedAddress.Line1)
}

func (suite *AddressRepoTestSuite) TestUpdateAddressError() {
	address := &models.Address{
		AddressId: "123",
		Line1:     "123 Main St",
		Line2:     "Apt 4B",
		City:      "Springfield",
		State:     "IL",
		Country:   "USA",
		ZipCode:   "62704",
	}

	err := suite.addressRepo.Update(address.AddressId, address)
	assert.Error(suite.T(), err)
}

func (suite *AddressRepoTestSuite) TestDeleteAddress() {
	address := &models.Address{
		AddressId: "123",
		Line1:     "123 Main St",
		Line2:     "Apt 4B",
		City:      "Springfield",
		State:     "IL",
		Country:   "USA",
		ZipCode:   "62704",
	}
	err := suite.addressRepo.Upsert(address)
	assert.NoError(suite.T(), err)

	err = suite.addressRepo.Delete(address.AddressId)

	assert.NoError(suite.T(), err)

	var deletedAddress models.Address
	err = suite.db.Table("addresses").Where("address_id = ?", address.AddressId).First(&deletedAddress).Error
	assert.Error(suite.T(), err)
}

func (suite *AddressRepoTestSuite) TestDeleteAddressError() {
	err := suite.addressRepo.Delete("123")

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "error deleting address", err.Error())
}
