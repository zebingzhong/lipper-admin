package blog

import (
	"github.com/gin-gonic/gin"
)

type Article struct{}

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Router /article/getList [get]
func (a Article) List(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"name":    "cl",
		"message": "kkk",
		"age":     "19",
	})
}
