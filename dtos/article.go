package dtos

type Article struct {
	ArticleId   string  `json:"article_id"`
	ArticleName string  `json:"article_name"`
	Price       float64 `json:"price"`
	Stock       int64   `json:"stock"`
}

type UpdateStock struct {
	NewStock int64 `json:"new_stock"`
}
