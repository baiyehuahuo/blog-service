package model

import (
	"blog-service/global"
	"blog-service/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Model struct {
	ID         uint32 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedOn  uint32 `gorm:"column:created_on;default:0"`  // 创建时间
	CreatedBy  string `gorm:"column:created_by"`            // 创建人
	ModifiedOn uint32 `gorm:"column:modified_on;default:0"` // 修改时间
	ModifiedBy string `gorm:"column:modified_by"`           // 修改人
	DeletedOn  uint32 `gorm:"column:deleted_on;default:0"`  // 删除时间
	IsDeleted  uint8  `gorm:"column:is_deleted;default:0"`  // 是否删除 0 未删除 1 已删除
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingConfig) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%v&loc=Local"
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(s,
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}
	if err = db.DB().Ping(); err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true) // ?
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db, nil
}

// Create 对应的回调函数
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if scope.HasError() {
		return
	}
	now := time.Now().UnixNano()
	// FieldByName 获取当前是否包含所需字段
	if createTimeField, ok := scope.FieldByName("CreatedOn"); ok && createTimeField.IsBlank {
		_ = createTimeField.Set(now)
	}
	if modifiedField, ok := scope.FieldByName("ModifiedOn"); ok && modifiedField.IsBlank {
		_ = modifiedField.Set(now)
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope) {
	if scope.HasError() {
		return
	}
	var extraOption string
	if str, ok := scope.Get("gorm:delete_option"); ok {
		extraOption = fmt.Sprint(str)
	}
	deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
	isDelField, hasIsDelField := scope.FieldByName("IsDeleted")
	if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
		// 软删除
		now := time.Now().Unix()
		scope.Raw(fmt.Sprintf(
			"UPDATE %v SET %v=%v, %v=%v%v%v",
			scope.QuotedTableName(),
			scope.Quote(deletedOnField.DBName),
			scope.AddToVars(now),
			scope.Quote(isDelField.DBName),
			scope.AddToVars(1),
			addExtraSpaceIfExist(scope.CombinedConditionSql()),
			addExtraSpaceIfExist(extraOption),
		)).Exec()
	} else {
		// 硬删除
		scope.Raw(fmt.Sprintf(
			"DELETE FROM %v%v%v",
			scope.QuotedTableName(),
			addExtraSpaceIfExist(scope.CombinedConditionSql()),
			addExtraSpaceIfExist(extraOption),
		)).Exec()
	}
}

func addExtraSpaceIfExist(str string) string {
	if str == "" {
		return ""
	}
	return " " + str
}
