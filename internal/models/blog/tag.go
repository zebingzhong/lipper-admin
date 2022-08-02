package blog

import "github.com/zebingzhong/lipper-admin/internal/models"

type Tag struct {
	*models.Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}
