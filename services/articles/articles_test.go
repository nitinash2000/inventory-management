package articles

import (
	"inventory-management/dtos"
	"inventory-management/models"
	"inventory-management/repository/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ArticlesTestSuite struct {
	suite.Suite
	mockCtrl        *gomock.Controller
	mockArticleRepo *mocks.MockArticleRepo
	articleService  IArticleService
}

func TestArticleTestSuite(t *testing.T) {
	suite.Run(t, new(ArticlesTestSuite))
}

func (suite *ArticlesTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())

	suite.mockArticleRepo = mocks.NewMockArticleRepo(suite.mockCtrl)

	suite.articleService = NewArticleService(suite.mockArticleRepo)
}

func (suite *ArticlesTestSuite) TestCreateArticle() {
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

func (suite *ArticlesTestSuite) TestGetArticle() {
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

func (suite *ArticlesTestSuite) TestUpdateArticle() {
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

func (suite *ArticlesTestSuite) TestDeleteArticle() {
	suite.mockArticleRepo.EXPECT().Delele("123").Return(nil).Times(1)

	err := suite.articleService.DeleleArticle("123")
	assert.NoError(suite.T(), err)
}

func (suite *ArticlesTestSuite) TestListArticle() {
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

func (suite *ArticlesTestSuite) TestUpdateArticleStock() {
	req := &dtos.UpdateStock{
		NewStock: 50,
	}

	suite.mockArticleRepo.EXPECT().UpdateArticleStock("123", int64(50)).Return(nil).Times(1)

	err := suite.articleService.UpdateArticleStock("123", req)
	assert.NoError(suite.T(), err)
}
