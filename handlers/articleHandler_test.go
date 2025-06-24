package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"inventory-management/constants"
	"inventory-management/dtos"
	"inventory-management/services/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type articleHandlerTestSuite struct {
	suite.Suite
	mockCtrl           *gomock.Controller
	mockArticleService *mocks.MockArticleService
	mockOrderService   *mocks.MockOrderService
	articleHandler     *articleHandler
}

func TestArticleHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(articleHandlerTestSuite))
}

func (suite *articleHandlerTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())

	suite.mockArticleService = mocks.NewMockArticleService(suite.mockCtrl)

	suite.articleHandler = NewArticleHandler(suite.mockArticleService)
}

func (suite *articleHandlerTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *articleHandlerTestSuite) TestGetArticle() {
	expected := &dtos.Article{
		ArticleId:   "123",
		ArticleName: "Test Article",
		Price:       100,
		Stock:       50,
	}

	suite.mockArticleService.EXPECT().GetArticle("123").Return(expected, nil).Times(1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodGet, "/articles/123", nil)

	suite.articleHandler.GetArticle(c)

	var result *dtos.Article

	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(suite.T(), err)

	assert.Equal(suite.T(), expected, result)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *articleHandlerTestSuite) TestGetArticleError() {
	suite.mockArticleService.EXPECT().GetArticle("123").Return(nil, constants.ErrorNotFound).Times(1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodGet, "/articles/123", nil)

	suite.articleHandler.GetArticle(c)

	var result *dtos.Article

	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), result)
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

func (suite *articleHandlerTestSuite) TestCreateArticle() {
	req := &dtos.Article{
		ArticleId:   "123",
		ArticleName: "Test Article",
		Price:       100,
		Stock:       50,
	}

	body, _ := json.Marshal(req)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/articles", bytes.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.mockArticleService.EXPECT().CreateArticle(req).Return(nil).Times(1)

	suite.articleHandler.CreateArticle(c)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *articleHandlerTestSuite) TestCreateArticleError() {
	req := &dtos.Article{
		ArticleId:   "123",
		ArticleName: "Test Article",
		Price:       100,
		Stock:       50,
	}

	body, _ := json.Marshal(req)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/articles", bytes.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.mockArticleService.EXPECT().CreateArticle(req).Return(constants.ErrorRecordExists).Times(1)

	suite.articleHandler.CreateArticle(c)
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

func (suite *articleHandlerTestSuite) TestCreateArticle_BadRequest() {
	invalidJSON := `{"articleId": "123", "articleName": "Test Article", "price": "not_a_number", "stock": "50"}`

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/articles", bytes.NewReader([]byte(invalidJSON)))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.articleHandler.CreateArticle(c)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *articleHandlerTestSuite) TestDeleteArticle() {
	suite.mockArticleService.EXPECT().DeleteArticle("123").Return(nil).Times(1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodDelete, "/articles/123", nil)

	suite.articleHandler.DeleteArticle(c)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *articleHandlerTestSuite) TestDeleteArticleError() {
	suite.mockArticleService.EXPECT().DeleteArticle("123").Return(constants.ErrorNotFound).Times(1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodDelete, "/articles/123", nil)

	suite.articleHandler.DeleteArticle(c)
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

func (suite *articleHandlerTestSuite) TestUpdateArticle() {
	req := &dtos.Article{
		ArticleId:   "123",
		ArticleName: "Test Article",
		Price:       200,
		Stock:       50,
	}

	body, _ := json.Marshal(req)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodPut, "/articles/123", bytes.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.mockArticleService.EXPECT().UpdateArticle("123", req).Return(nil).Times(1)

	suite.articleHandler.UpdateArticle(c)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *articleHandlerTestSuite) TestUpdateArticleError() {
	req := &dtos.Article{
		ArticleId:   "123",
		ArticleName: "Test Article",
		Price:       200,
		Stock:       50,
	}

	body, _ := json.Marshal(req)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodPut, "/articles/123", bytes.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.mockArticleService.EXPECT().UpdateArticle("123", req).Return(constants.ErrorNotFound).Times(1)

	suite.articleHandler.UpdateArticle(c)
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

func (suite *articleHandlerTestSuite) TestUpdateArticleBadRequest() {
	invalidJSON := `{"articleId": "123", "articleName": "Test Article", "price": "not_a_number", "stock": "50"}`

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodPut, "/articles/123", bytes.NewReader([]byte(invalidJSON)))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.articleHandler.UpdateArticle(c)
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *articleHandlerTestSuite) TestListArticles() {
	expected := []*dtos.Article{
		{
			ArticleId:   "123",
			ArticleName: "Test Article",
			Price:       100,
			Stock:       50,
		},
		{
			ArticleId:   "321",
			ArticleName: "Test 2",
			Price:       200,
			Stock:       10,
		},
	}

	suite.mockArticleService.EXPECT().ListArticle().Return(expected, nil).Times(1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/articles-list", nil)

	suite.articleHandler.ListArticles(c)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *articleHandlerTestSuite) TestListArticlesError() {
	suite.mockArticleService.EXPECT().ListArticle().Return(nil, errors.New("repo error")).Times(1)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/articles-list", nil)

	suite.articleHandler.ListArticles(c)
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

func (suite *articleHandlerTestSuite) TestUpdateArticleStock() {
	req := &dtos.UpdateStock{
		NewStock: 200,
	}

	body, _ := json.Marshal(req)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodPatch, "/articles/123", bytes.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.mockArticleService.EXPECT().UpdateArticleStock("123", req).Return(nil).Times(1)

	suite.articleHandler.UpdateArticleStock(c)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *articleHandlerTestSuite) TestUpdateArticleStockError() {
	req := &dtos.UpdateStock{
		NewStock: 200,
	}

	body, _ := json.Marshal(req)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodPatch, "/articles/123", bytes.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.mockArticleService.EXPECT().UpdateArticleStock("123", req).Return(constants.ErrorNotFound).Times(1)

	suite.articleHandler.UpdateArticleStock(c)
	assert.Equal(suite.T(), http.StatusInternalServerError, w.Code)
}

func (suite *articleHandlerTestSuite) TestUpdateArticleStockBadRequest() {
	invalidJSON := `{"new_id": "123", "articleName": "Test Article"`

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{Key: "id", Value: "123"},
	}
	c.Request = httptest.NewRequest(http.MethodPatch, "/articles/123", bytes.NewReader([]byte(invalidJSON)))
	c.Request.Header.Set("Content-Type", "application/json")

	suite.articleHandler.UpdateArticleStock(c)
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}
