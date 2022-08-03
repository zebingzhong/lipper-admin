package v1

import (
	"github.com/zebingzhong/lipper-admin/internal/api/v1/blog"
	"github.com/zebingzhong/lipper-admin/internal/api/v1/system"
)

type ApiGroup struct {
	BlogApiGroup   blog.ApiGroup
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
