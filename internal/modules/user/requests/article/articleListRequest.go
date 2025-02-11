package article

type ArticleListRequest struct {
	Page     int    `form:"page" binding:"required,gte=0"`
	Limit    int    `form:"limit" binding:"required,gte=1"`
	KeyWords string `form:"key_words"`
}
