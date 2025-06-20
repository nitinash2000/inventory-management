package models

type Article struct {
	ArticleId   string  `json:"article_id" gorm:"primaryKey"`
	ArticleName string  `json:"article_name"`
	Price       float64 `json:"price"`
	Stock       int64   `json:"stock"`
}
