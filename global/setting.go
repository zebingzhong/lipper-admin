package global

import (
	"github.com/zebingzhong/lipper-admin/pkg/logger"
	"github.com/zebingzhong/lipper-admin/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
