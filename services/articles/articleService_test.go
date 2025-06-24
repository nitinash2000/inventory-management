package articles

import (
	"errors"
	"inventory-management/constants"
	"inventory-management/dtos"
	"inventory-management/models"
	"inventory-management/repository/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type articleServiceTestSuite struct {
	suite.Suite
	mockCtrl        *gomock.Controller
	mockArticleRepo *mocks.MockArticleRepo
	articleService  ArticleService
}

func TestArticleTestSuite(t *testing.T) {
	suite.Run(t, new(articleServiceTestSuite))
}

func (suite *articleServiceTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())

	suite.mockArticleRepo = mocks.NewMockArticleRepo(suite.mockCtrl)

	suite.articleService = NewArticleService(suite.mockArticleRepo)
}

func (suite *articleServiceTestSuite) TestCreateArticle() {
	req := &dtos.Article{
		ArticleId:   "123",
		ArticleName: "Test Article",
		Price:       100,
		Stock:       50,
	}

	model := &models.Article{
		ArticleId:   "123",
		ArticleName: "Test Article",
		Price:       100,
		Stock:       50,
	}

	suite.mockArticleRepo.EXPECT().Create(model).Return(nil).Times(1)

	err := suite.articleService.CreateArticle(req)
	assert.NoError(suite.T(), err)
}

func (suite *articleServiceTestSuite) TestCreateArticleRepoError() {
	req := &dtos.Article{
		ArticleId:   "123",
		ArticleName: "Test Article",
		Price:       100,
		Stock:       50,
	}

	model := &models.Article{
		ArticleId:   "123",
		ArticleName: "Test Article",
		Price:       100,
		Stock:       50,
	}

	suite.mockArticleRepo.EXPECT().Create(model).Return(errors.New("repo error")).Times(1)

	err := suite.articleService.CreateArticle(req)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "repo error", err.Error())
}

func (suite *articleServiceTestSuite) TestGetArticle() {
	expected := &dtos.Article{
		ArticleId:   "123",
		ArticleName: "Test Article",
		Price:       100,
		Stock:       50,
	}

	mockArticle := &models.Article{
		ArticleId:   "123",
		ArticleName: "Test Article",
		Price:       100,
		Stock:       50,
	}

	suite.mockArticleRepo.EXPECT().Get("123").Return(mockArticle, nil).Times(1)

	result, err := suite.articleService.GetArticle("123")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, result)
}

func (suite *articleServiceTestSuite) TestGetArticleError() {
	expectedError := constants.ErrorNotFound

	suite.mockArticleRepo.EXPECT().Get("123").Return(nil, expectedError).Times(1)

	result, err := suite.articleService.GetArticle("123")

	assert.Nil(suite.T(), result)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
}

func (suite *articleServiceTestSuite) TestUpdateArticle() {
	req := &dtos.Article{
		ArticleId:   "123",
		ArticleName: "Test Article",
		Price:       100,
		Stock:       50,
	}

	model := &models.Article{
		ArticleId:   "123",
		ArticleName: "Test Article",
		Price:       100,
		Stock:       50,
	}

	suite.mockArticleRepo.EXPECT().Update("123", model).Return(nil).Times(1)

	err := suite.articleService.UpdateArticle("123", req)
	assert.NoError(suite.T(), err)
}

func (suite *articleServiceTestSuite) TestUpdateArticleError() {
	req := &dtos.Article{
		ArticleId:   "123",
		ArticleName: "Test Article",
		Price:       100,
		Stock:       50,
	}

	model := &models.Article{
		ArticleId:   "123",
		ArticleName: "Test Article",
		Price:       100,
		Stock:       50,
	}

	suite.mockArticleRepo.EXPECT().Update("123", model).Return(constants.ErrorNotFound).Times(1)

	err := suite.articleService.UpdateArticle("123", req)
	assert.Error(suite.T(), err)
}

func (suite *articleServiceTestSuite) TestDeleteArticle() {
	suite.mockArticleRepo.EXPECT().Delete("123").Return(nil).Times(1)

	err := suite.articleService.DeleteArticle("123")
	assert.NoError(suite.T(), err)
}

func (suite *articleServiceTestSuite) TestDeleteArticleError() {
	suite.mockArticleRepo.EXPECT().Delete("123").Return(constants.ErrorNotFound).Times(1)

	err := suite.articleService.DeleteArticle("123")
	assert.Error(suite.T(), err)
}

func (suite *articleServiceTestSuite) TestListArticle() {
	expectedDto := []*dtos.Article{
		{
			ArticleId:   "123",
			ArticleName: "Test 1",
			Price:       100,
			Stock:       50,
		},
		{
			ArticleId:   "321",
			ArticleName: "Test 2",
			Price:       200,
			Stock:       20,
		},
	}

	mockModel := []*models.Article{
		{
			ArticleId:   "123",
			ArticleName: "Test 1",
			Price:       100,
			Stock:       50,
		},
		{
			ArticleId:   "321",
			ArticleName: "Test 2",
			Price:       200,
			Stock:       20,
		},
	}

	suite.mockArticleRepo.EXPECT().GetAll().Return(mockModel, nil).Times(1)

	result, err := suite.articleService.ListArticle()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedDto, result)
}

func (suite *articleServiceTestSuite) TestListArticleError() {
	suite.mockArticleRepo.EXPECT().GetAll().Return(nil, errors.New("repo error")).Times(1)

	result, err := suite.articleService.ListArticle()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
}

func (suite *articleServiceTestSuite) TestUpdateArticleStock() {
	req := &dtos.UpdateStock{
		NewStock: 50,
	}

	suite.mockArticleRepo.EXPECT().UpdateArticleStock("123", int64(50)).Return(nil).Times(1)

	err := suite.articleService.UpdateArticleStock("123", req)
	assert.NoError(suite.T(), err)
}

func (suite *articleServiceTestSuite) TestUpdateArticleStockError() {
	req := &dtos.UpdateStock{
		NewStock: 50,
	}

	suite.mockArticleRepo.EXPECT().UpdateArticleStock("123", int64(50)).Return(constants.ErrorNotFound).Times(1)

	err := suite.articleService.UpdateArticleStock("123", req)
	assert.Error(suite.T(), err)
}

func (suite *articleServiceTestSuite) TestArticleDtosToModel() {
	req := &dtos.Article{
		ArticleId:   "123",
		ArticleName: "Test Article",
		Price:       100,
		Stock:       50,
	}

	expected := &models.Article{
		ArticleId:   "123",
		ArticleName: "Test Article",
		Price:       100,
		Stock:       50,
	}

	result := ArticleDtosToModel(req)
	assert.Equal(suite.T(), expected, result)
}

func (suite *articleServiceTestSuite) TestArticleModelToDtos() {
	model := &models.Article{
		ArticleId:   "123",
		ArticleName: "Test Article",
		Price:       100,
		Stock:       50,
	}

	expected := &dtos.Article{
		ArticleId:   "123",
		ArticleName: "Test Article",
		Price:       100,
		Stock:       50,
	}

	result := ArticleModelToDtos(model)
	assert.Equal(suite.T(), []*dtos.Article{expected}, result)
}
