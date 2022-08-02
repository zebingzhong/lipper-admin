package global

import (
	"github.com/zebingzhong/lipper-admin/pkg/logger"
	"github.com/zebingzhong/lipper-admin/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	SystemSetting   *setting.SystemSettingS
	LocalSetting    *setting.LocalSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	QiNiuSetting    *setting.QiNiu
	Logger          *logger.Logger
)
