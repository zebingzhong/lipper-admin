package v1

import "github.com/zebingzhong/lipper-admin/internal/api/v1/blog"

type ApiGroup struct {
	BlogApiGroup blog.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
