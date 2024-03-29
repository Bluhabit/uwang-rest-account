// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameSystemAccess = "system_access"

// SystemAccess mapped from table <system_access>
type SystemAccess struct {
	ID         string    `gorm:"column:id;primaryKey" json:"id"`
	Name       string    `gorm:"column:name;not null" json:"name"`
	Permission string    `gorm:"column:permission;not null" json:"permission"`
	Group      string    `gorm:"column:group;not null" json:"group"`
	CreatedAt  time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Deleted    bool      `gorm:"column:deleted;not null" json:"deleted"`
}

// TableName SystemAccess's table name
func (*SystemAccess) TableName() string {
	return TableNameSystemAccess
}
