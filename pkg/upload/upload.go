package upload

import (
	"mime/multipart"

	"github.com/zebingzhong/lipper-admin/global"
)

// OSS 对象存储接口
type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

// NewOss OSS的实例化方法
func NewOss() OSS {
	switch global.SystemSetting.OSSType {
	case "local":
		return &Local{}
	case "qiniu":
		return &QiNiu{}
	default:
		return &Local{}
	}
}
