package repository

import (
	"inventory-management/models"

	"gorm.io/gorm"
)

type ArticleRepo interface {
	Create(article *models.Article) error
	Update(articleId string, article *models.Article) error
	Get(articleId string) (*models.Article, error)
	GetAll() ([]*models.Article, error)
	Delete(articleId string) error
	UpdateArticleStock(articleId string, stock int64) error
}

type articleRepo struct {
	db *gorm.DB
}

func NewArticleRepo(db *gorm.DB) ArticleRepo {
	return &articleRepo{
		db: db,
	}
}

func (a *articleRepo) getTable() string {
	return "articles"
}

func (a *articleRepo) Create(article *models.Article) error {
	err := a.db.Table(a.getTable()).Save(article).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *articleRepo) Update(articleId string, article *models.Article) error {
	err := a.db.Table(a.getTable()).Where("article_id = ?", articleId).UpdateColumns(article).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *articleRepo) Get(articleId string) (*models.Article, error) {
	var result *models.Article

	err := a.db.Table(a.getTable()).Where("article_id = ?", articleId).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *articleRepo) GetAll() ([]*models.Article, error) {
	var result []*models.Article

	tx := a.db.Table(a.getTable()).Where("1=1").Find(&result)
	err := tx.Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *articleRepo) Delete(articleId string) error {
	err := a.db.Table(a.getTable()).Where("article_id = ?", articleId).Delete(&models.Article{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *articleRepo) UpdateArticleStock(articleId string, stock int64) error {
	err := a.db.Table(a.getTable()).Where("article_id = ?", articleId).Update("stock", stock).Error
	if err != nil {
		return err
	}

	return nil
}
