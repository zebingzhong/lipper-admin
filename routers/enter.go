package routers

import "github.com/zebingzhong/lipper-admin/routers/blog"

type RouterGroup struct {
	Blog blog.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
