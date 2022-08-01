package blog

import (
	"github.com/gin-gonic/gin"
)

type Article struct{}

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {object} response.Response{data=systemRes.SysMenusResponse,msg=string} "获取用户动态路由,返回包括系统菜单详情列表"
// @Router /article/getList [get]
func (a Article) List(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"name":    "cl",
		"message": "kkk",
		"age":     "19",
	})
}
