package category

import "github.com/zebingzhong/lipper-admin/internal/models"

type FoodCategory struct {
	*models.Model
	CategoryCode  string `json:"category_code" gorm:"size:64;comment:分类编码" `
	CategoryName  string `json:"category_name" gorm:"size:64;comment:分类名称" `
	CategoryImage string `json:"category_image" gorm:"type:text;comment:分类图片" `
	CategoryLevel int8   `json:"category_level" gorm:"comment:分类级别" `
	CategorySort  uint   `json:"category_sort" gorm:"comment:排序" `
}
