package articles

import (
	"inventory-management/dtos"
	"inventory-management/models"
	"inventory-management/repository"
)

type IArticleService interface {
	CreateArticle(req *dtos.Article) error
	UpdateArticle(id string, req *dtos.Article) error
	GetArticle(articleId string) (*dtos.Article, error)
	ListArticle() ([]*dtos.Article, error)
	DeleleArticle(articleId string) error
	UpdateArticleStock(articleId string, req *dtos.UpdateStock) error
}

type ArticleService struct {
	ArticleRepo repository.IArticle
}

func NewArticleService() IArticleService {
	return &ArticleService{
		ArticleRepo: repository.NewArticle(),
	}
}

func (a *ArticleService) CreateArticle(req *dtos.Article) error {
	model := ArticleDtosToModel(req)

	err := a.ArticleRepo.Create(model)
	if err != nil {
		return err
	}

	return nil
}

func (a *ArticleService) UpdateArticle(id string, req *dtos.Article) error {
	model := ArticleDtosToModel(req)

	err := a.ArticleRepo.Update(id, model)
	if err != nil {
		return err
	}

	return nil
}

func (a *ArticleService) GetArticle(articleId string) (*dtos.Article, error) {
	article, err := a.ArticleRepo.Get(articleId)
	if err != nil {
		return nil, err
	}

	result := ArticleModelToDtos(article)

	return result[0], nil
}

func (a *ArticleService) ListArticle() ([]*dtos.Article, error) {
	articles, err := a.ArticleRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return ArticleModelToDtos(articles...), nil
}

func (a *ArticleService) DeleleArticle(articleId string) error {
	err := a.ArticleRepo.Delele(articleId)
	if err != nil {
		return err
	}

	return nil
}

func (a *ArticleService) UpdateArticleStock(articleId string, req *dtos.UpdateStock) error {
	err := a.ArticleRepo.UpdateArticleStock(articleId, req.NewStock)
	if err != nil {
		return err
	}

	return nil
}

func ArticleModelToDtos(m ...*models.Article) []*dtos.Article {
	var a []*dtos.Article

	for _, v := range m {
		a = append(a, &dtos.Article{
			ArticleId:   v.ArticleId,
			ArticleName: v.ArticleName,
			Price:       v.Price,
			Stock:       v.Stock,
		})
	}

	return a
}

func ArticleDtosToModel(m *dtos.Article) *models.Article {
	return &models.Article{
		ArticleId:   m.ArticleId,
		ArticleName: m.ArticleName,
		Price:       m.Price,
		Stock:       m.Stock,
	}
}
