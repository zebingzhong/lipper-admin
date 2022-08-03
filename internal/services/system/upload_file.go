package system

import (
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/zebingzhong/lipper-admin/internal/models/system"
	"github.com/zebingzhong/lipper-admin/pkg/upload"
)

type UploadFileService struct{}

func (e *UploadFileService) UploadFile(header *multipart.FileHeader, noSave string) (file system.FileUpload, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	fmt.Println(uploadErr)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := system.FileUpload{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		return f, nil
	}
	return
}
