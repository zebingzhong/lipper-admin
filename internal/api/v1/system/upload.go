package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zebingzhong/lipper-admin/internal/models/system"
	"github.com/zebingzhong/lipper-admin/internal/services"
	"github.com/zebingzhong/lipper-admin/pkg/response"
)

type Upload struct{}

// @Tags Upload
// @Summary 上传文件
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {object} response.Response{data=system.FileResponse,msg=string} "上传文件示例,返回包括文件详情"
// @Router /fileUploadAndDownload/upload [post]
func (b *Upload) UploadFile(c *gin.Context) {
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		response.FailWithMessage("接收文件失败！", c)
		return
	}
	file, err := services.ServiceGroupApp.SystemServiceGroup.UploadFileService.UploadFile(header, noSave)
	if err != nil {
		response.FailWithMessage("接收文件失败！", c)
		return
	}
	response.OkWithDetailed(system.FileResponse{File: file}, "上传成功", c)
}
