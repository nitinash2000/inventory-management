package repository

import (
	"inventory-management/constants"
	"inventory-management/models"
)

type IArticle interface {
	Create(article *models.Article) error
	Update(articleId string, article *models.Article) error
	Get(articleId string) (*models.Article, error)
	GetAll() ([]*models.Article, error)
	Delele(articleId string) error
	UpdateArticleStock(articleId string, stock int64) error
}

var articleMap map[string]*models.Article

func init() {
	articleMap = map[string]*models.Article{
		"123": {
			ArticleId:   "123",
			ArticleName: "Book",
			Price:       200,
			Stock:       4,
		},
	}
}

type Article struct {
	// articleMap map[string]*models.Article
}

func NewArticle() IArticle {
	return &Article{
		// articleMap: articleMap
	}
}

func (a *Article) Create(article *models.Article) error {
	if _, exists := articleMap[article.ArticleId]; exists {
		return constants.ErrorRecordExists
	}

	articleMap[article.ArticleId] = article
	return nil
}

func (a *Article) Update(articleId string, article *models.Article) error {
	if _, exists := articleMap[article.ArticleId]; !exists {
		return constants.ErrorNotFound
	}

	articleMap[articleId] = article
	return nil
}

func (a *Article) Get(articleId string) (*models.Article, error) {
	article, exists := articleMap[articleId]

	if !exists {
		return nil, constants.ErrorNotFound
	}

	return article, nil
}

func (a *Article) GetAll() ([]*models.Article, error) {
	var result []*models.Article

	for _, v := range articleMap {
		result = append(result, v)
	}

	return result, nil
}

func (a *Article) Delele(articleId string) error {
	if _, exists := articleMap[articleId]; !exists {
		return constants.ErrorNotFound
	}

	delete(articleMap, articleId)

	return nil
}

func (a *Article) UpdateArticleStock(articleId string, stock int64) error {
	article, exists := articleMap[articleId]

	if !exists {
		return constants.ErrorNotFound
	}

	article.Stock = stock
	articleMap[articleId] = article

	return nil
}
