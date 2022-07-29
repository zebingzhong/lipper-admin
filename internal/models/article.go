package models

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverTmageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}
