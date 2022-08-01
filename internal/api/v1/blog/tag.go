package blog

import (
	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t *Tag) Get(c *gin.Context) {}

// ListAccounts godoc
// @Summary     List accounts
// @Description get accounts
// @Tags        accounts
// @Accept      json
// @Produce     json
// @Param       q   query    string false "name search by q" Format(email)
// @Success     200 {array}  model.Account
// @Failure     400 {object} httputil.HTTPError
// @Failure     404 {object} httputil.HTTPError
// @Failure     500 {object} httputil.HTTPError
// @Router      /tags [get]
func (t *Tag) List(c *gin.Context) {}

func (t *Tag) Create(c *gin.Context) {

}

func (t *Tag) Update(c *gin.Context) {}

func (t *Tag) Delete(c *gin.Context) {}
