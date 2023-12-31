package dbplugin

import (
	"gorm.io/plugin/soft_delete"
)

type M struct {
	ID        uint                  `gorm:"primaryKey"`
	CreatedAt int                   `gorm:"autoCreateTime:milli"`
	UpdatedAt int                   `gorm:"autoUpdateTime:milli"`
	DeletedAt soft_delete.DeletedAt `gorm:"index;softDelete:milli"`
}
