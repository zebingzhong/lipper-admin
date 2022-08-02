package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zebingzhong/lipper-admin/internal/services/system"
	"github.com/zebingzhong/lipper-admin/pkg/common"
)

type Upload struct{}

func (b *Upload) UploadFile(c *gin.Context) {
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		common.FailWithMessage("接收文件失败！", c)
		return
	}
	file, err = system.UploadFileService.UploadFile(header, noSave)

}
