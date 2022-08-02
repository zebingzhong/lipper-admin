package blog

import "github.com/zebingzhong/lipper-admin/internal/models"

type Article struct {
	*models.Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverTmageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}
