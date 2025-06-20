package repository

import (
	"inventory-management/constants"
	"inventory-management/models"
)

type ArticleRepo interface {
	Create(article *models.Article) error
	Update(articleId string, article *models.Article) error
	Get(articleId string) (*models.Article, error)
	GetAll() ([]*models.Article, error)
	Delele(articleId string) error
	UpdateArticleStock(articleId string, stock int64) error
}

type articleRepo struct {
	articleMap map[string]*models.Article
}

func NewArticleRepo(articleMap map[string]*models.Article) ArticleRepo {
	return &articleRepo{
		articleMap: articleMap,
	}
}

func (a *articleRepo) Create(article *models.Article) error {
	if _, exists := a.articleMap[article.ArticleId]; exists {
		return constants.ErrorRecordExists
	}

	a.articleMap[article.ArticleId] = article
	return nil
}

func (a *articleRepo) Update(articleId string, article *models.Article) error {
	if _, exists := a.articleMap[article.ArticleId]; !exists {
		return constants.ErrorNotFound
	}

	a.articleMap[articleId] = article
	return nil
}

func (a *articleRepo) Get(articleId string) (*models.Article, error) {
	article, exists := a.articleMap[articleId]

	if !exists {
		return nil, constants.ErrorNotFound
	}

	return article, nil
}

func (a *articleRepo) GetAll() ([]*models.Article, error) {
	var result []*models.Article

	for _, v := range a.articleMap {
		result = append(result, v)
	}

	return result, nil
}

func (a *articleRepo) Delele(articleId string) error {
	if _, exists := a.articleMap[articleId]; !exists {
		return constants.ErrorNotFound
	}

	delete(a.articleMap, articleId)

	return nil
}

func (a *articleRepo) UpdateArticleStock(articleId string, stock int64) error {
	article, exists := a.articleMap[articleId]

	if !exists {
		return constants.ErrorNotFound
	}

	article.Stock = stock
	a.articleMap[articleId] = article

	return nil
}
