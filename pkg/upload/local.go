package upload

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"

	"github.com/zebingzhong/lipper-admin/global"
	"github.com/zebingzhong/lipper-admin/pkg/common"
)

type Local struct{}

//@author lipper
//@object: *Local
//@function: DeleteFile
//@description: 上传文件
//@param: key string
//@return: error
func (*Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = common.MD5V([]byte(name))
	// 拼接文件名
	fileName := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建存储路径
	mkdirErr := os.MkdirAll(global.LocalSetting.StorePath, os.ModePerm)
	if mkdirErr != nil {
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}

	// 拼接路径和文件名
	p := global.LocalSetting.StorePath + "/" + fileName
	filePath := global.LocalSetting.Path + "/" + fileName
	// 读取文件
	f, openError := file.Open() // 读取文件
	if openError != nil {
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	// 关闭句柄
	defer f.Close()
	// 创建文件
	out, createErr := os.Create(p)
	if createErr != nil {
		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close()
	// 传输（拷贝）文件
	_, copyErr := io.Copy(out, f)
	if copyErr != nil {
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return filePath, fileName, nil
}

//@author lipper
//@object: *Local
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error
func (*Local) DeleteFile(key string) error {
	p := global.LocalSetting.StorePath + "/" + key
	if strings.Contains(p, global.LocalSetting.StorePath) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}
