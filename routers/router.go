package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/zebingzhong/lipper-admin/docs"
)

func NewRouter() *gin.Engine {
	Router := gin.Default()

	blogRouter := RouterGroupApp.Blog

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	PublicGroup := Router.Group("/api")
	{
		blogRouter.InitArticleRouter(PublicGroup)
	}

	return Router
}
