package models

import (
	"fmt"
	"time"

	"github.com/zebingzhong/lipper-admin/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	// 设置空闲连接池中连接最大数量
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	// 设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Duration(databaseSetting.MaxLifetimes))
	return db, nil
}
