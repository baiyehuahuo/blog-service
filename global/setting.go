package global

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
)

var (
	ServerSetting   *setting.ServerSettingConfig
	AppSetting      *setting.AppSettingConfig
	DatabaseSetting *setting.DatabaseSettingConfig
	JWTSetting      *setting.JWTSettingConfig
	EmailSetting    *setting.EmailSettingConfig

	DBEngine *gorm.DB
	Logger   *logger.Logger
)
