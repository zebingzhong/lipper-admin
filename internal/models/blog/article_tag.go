package blog

import "github.com/zebingzhong/lipper-admin/internal/models"

type ArticleTag struct {
	*models.Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}
