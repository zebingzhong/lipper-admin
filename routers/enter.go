package routers

import (
	"github.com/zebingzhong/lipper-admin/routers/blog"
	"github.com/zebingzhong/lipper-admin/routers/system"
)

type RouterGroup struct {
	Blog   blog.RouterGroup
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
