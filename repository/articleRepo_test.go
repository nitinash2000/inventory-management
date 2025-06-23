package repository

import (
	"inventory-management/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ArticleRepoTestSuite struct {
	suite.Suite
	db          *gorm.DB
	articleRepo ArticleRepo
}

func TestArticleRepoTestSuite(t *testing.T) {
	suite.Run(t, new(ArticleRepoTestSuite))
}

func (suite *ArticleRepoTestSuite) SetupTest() {
	var err error
	suite.db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		suite.T().Fatal("failed to connect to database")
	}

	err = suite.db.AutoMigrate(&models.Article{})
	if err != nil {
		suite.T().Fatal("failed to migrate database")
	}

	suite.articleRepo = NewArticleRepo(suite.db)
}

func (suite *ArticleRepoTestSuite) TearDownTest() {
	sqlDB, _ := suite.db.DB()
	sqlDB.Close()
}

func (suite *ArticleRepoTestSuite) TestCreateArticle() {
	article := &models.Article{
		ArticleId:   "123",
		ArticleName: "article1",
		Price:       100.5,
		Stock:       6,
	}

	err := suite.articleRepo.Create(article)

	assert.NoError(suite.T(), err)

	var savedArticle models.Article
	suite.db.Table("articles").Where("article_id = ?", article.ArticleId).First(&savedArticle)
	assert.Equal(suite.T(), article.ArticleId, savedArticle.ArticleId)
	assert.Equal(suite.T(), article.ArticleName, savedArticle.ArticleName)
}

func (suite *ArticleRepoTestSuite) TestGetArticle() {
	article := &models.Article{
		ArticleId:   "123",
		ArticleName: "article1",
		Price:       100.5,
		Stock:       6,
	}
	err := suite.articleRepo.Create(article)
	assert.NoError(suite.T(), err)

	result, err := suite.articleRepo.Get("123")

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), article.ArticleId, result.ArticleId)
	assert.Equal(suite.T(), article.ArticleName, result.ArticleName)
}

func (suite *ArticleRepoTestSuite) TestUpdateArticle() {
	article := &models.Article{
		ArticleId:   "123",
		ArticleName: "article1",
		Price:       100.5,
		Stock:       6,
	}
	err := suite.articleRepo.Create(article)
	assert.NoError(suite.T(), err)

	article.ArticleName = "article2"

	err = suite.articleRepo.Update(article.ArticleId, article)
	assert.NoError(suite.T(), err)

	var updatedArticle models.Article
	suite.db.Table("articles").Where("article_id = ?", article.ArticleId).First(&updatedArticle)
	assert.Equal(suite.T(), "article2", updatedArticle.ArticleName)
}

func (suite *ArticleRepoTestSuite) TestDeleteArticle() {
	article := &models.Article{
		ArticleId:   "123",
		ArticleName: "article1",
		Price:       100.5,
		Stock:       6,
	}
	err := suite.articleRepo.Create(article)
	assert.NoError(suite.T(), err)

	err = suite.articleRepo.Delete(article.ArticleId)

	assert.NoError(suite.T(), err)

	var deletedArticle models.Article
	err = suite.db.Table("articles").Where("article_id = ?", article.ArticleId).First(&deletedArticle).Error
	assert.Error(suite.T(), err)
}
