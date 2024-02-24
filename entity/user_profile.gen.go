// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameUserProfile = "user_profile"

// UserProfile mapped from table <user_profile>
type UserProfile struct {
	ID        string    `gorm:"column:id;primaryKey;default:326ee074-372c-4eef-bd16-b454308e74fa" json:"id"`
	Key       string    `gorm:"column:key;not null" json:"key"`
	Value     string    `gorm:"column:value;not null" json:"value"`
	UserID    string    `gorm:"column:user_id" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Deleted   bool      `gorm:"column:deleted;not null" json:"deleted"`
}

// TableName UserProfile's table name
func (*UserProfile) TableName() string {
	return TableNameUserProfile
}
