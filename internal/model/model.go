package model

type Model struct {
	ID         uint32 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CreatedOn  uint32 `gorm:"column:created_on;default:0"`  // 创建时间
	CreatedBy  string `gorm:"column:created_by"`            // 创建人
	ModifiedOn uint32 `gorm:"column:modified_on;default:0"` // 修改时间
	ModifiedBy string `gorm:"column:modified_by"`           // 修改人
	DeletedOn  uint32 `gorm:"column:deleted_on;default:0"`  // 删除时间
	IsDeleted  uint8  `gorm:"column:is_deleted;default:0"`  // 是否删除 0 未删除 1 已删除
}
