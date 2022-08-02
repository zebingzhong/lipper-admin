package request

type ArticleListRequest struct {
	Name  string `json:"name" binding:"required,max=10"`
	State uint8  `json:"state" binding:"required,oneof=0 1"`
}
