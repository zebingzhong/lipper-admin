package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/zebingzhong/lipper-admin/global"
)

type QiNiu struct{}

//@author: lipper
//@object: *Qiniu
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error
func (*QiNiu) UploadFile(file *multipart.FileHeader) (string, string, error) {
	putPolicy := storage.PutPolicy{Scope: global.QiNiuSetting.Bucket}
	mac := qbox.NewMac(global.QiNiuSetting.AccessKey, global.QiNiuSetting.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := qiniuConfig()
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}
	f, openError := file.Open()
	if openError != nil {
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close()
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	putErr := formUploader.Put(context.Background(), &ret, upToken, fileKey, f, file.Size, &putExtra)
	if putErr != nil {
		return "", "", errors.New("function formUploader.Put() Filed, err:" + putErr.Error())
	}
	return global.QiNiuSetting.ImgPath + "/" + ret.Key, ret.Key, nil
}

//@author lipper
//@object: *Qiniu
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error
func (*QiNiu) DeleteFile(key string) error {
	mac := qbox.NewMac(global.QiNiuSetting.AccessKey, global.QiNiuSetting.SecretKey)
	cfg := qiniuConfig()
	bucketManager := storage.NewBucketManager(mac, cfg)
	if err := bucketManager.Delete(global.QiNiuSetting.Bucket, key); err != nil {
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}
	return nil
}

//@author lipper
//@object: *Qiniu
//@function: qiniuConfig
//@description: 根据配置文件进行返回七牛云的配置
//@return: *storage.Config
func qiniuConfig() *storage.Config {
	cfg := storage.Config{
		UseHTTPS:      global.QiNiuSetting.UseHttps,
		UseCdnDomains: global.QiNiuSetting.UseCdnDomains,
	}
	return &cfg
}
