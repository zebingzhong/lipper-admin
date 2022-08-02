package blog

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zebingzhong/lipper-admin/internal/api/v1"
)

type ArticelRouter struct{}

func (b *ArticelRouter) InitArticleRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	articelRouter := Router.Group("article")
	articleApi := v1.ApiGroupApp.BlogApiGroup.Article
	{
		articelRouter.POST("getList", articleApi.List)
	}
	return articelRouter
}
