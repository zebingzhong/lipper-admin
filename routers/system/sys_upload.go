package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zebingzhong/lipper-admin/internal/api/v1"
)

type FileUploadRouter struct{}

func (s *BaseRouter) InitFileUploadRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	uploadRouter := Router.Group("upload")
	uploadFileApi := v1.ApiGroupApp.SystemApiGroup
	{
		uploadRouter.POST("upload", uploadFileApi.UploadFile)
	}
	return uploadRouter
}
