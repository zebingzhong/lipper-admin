package blog

import (
	"github.com/gin-gonic/gin"
	"github.com/zebingzhong/lipper-admin/internal/models/request"
)

type Article struct{}

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Router /article/getList [post]
func (a Article) List(c *gin.Context) {
	param := request.ArticleListRequest{}
	err := c.ShouldBind(&param)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  0,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  1,
		"message": "success",
	})

}
