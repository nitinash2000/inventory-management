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

type articleService struct {
	articleRepo repository.ArticleRepo
}

func NewArticleService(articleRepo repository.ArticleRepo) IArticleService {
	return &articleService{
		articleRepo: articleRepo,
	}
}

func (a *articleService) CreateArticle(req *dtos.Article) error {
	model := ArticleDtosToModel(req)

	err := a.articleRepo.Create(model)
	if err != nil {
		return err
	}

	return nil
}

func (a *articleService) UpdateArticle(id string, req *dtos.Article) error {
	model := ArticleDtosToModel(req)

	err := a.articleRepo.Update(id, model)
	if err != nil {
		return err
	}

	return nil
}

func (a *articleService) GetArticle(articleId string) (*dtos.Article, error) {
	article, err := a.articleRepo.Get(articleId)
	if err != nil {
		return nil, err
	}

	result := ArticleModelToDtos(article)

	return result[0], nil
}

func (a *articleService) ListArticle() ([]*dtos.Article, error) {
	articles, err := a.articleRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return ArticleModelToDtos(articles...), nil
}

func (a *articleService) DeleleArticle(articleId string) error {
	err := a.articleRepo.Delele(articleId)
	if err != nil {
		return err
	}

	return nil
}

func (a *articleService) UpdateArticleStock(articleId string, req *dtos.UpdateStock) error {
	err := a.articleRepo.UpdateArticleStock(articleId, req.NewStock)
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
