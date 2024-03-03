// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameUserLog = "user_log"

// UserLog mapped from table <user_log>
type UserLog struct {
	ID        string    `gorm:"column:id;primaryKey;default:f7729194-55ce-47f4-9486-a4532f17fea6" json:"id"`
	UserID    string    `gorm:"column:user_id" json:"user_id"`
	IPAddress string    `gorm:"column:ip_address" json:"ip_address"`
	LogType   string    `gorm:"column:log_type" json:"log_type"`
	Content   string    `gorm:"column:content" json:"content"`
	Device    string    `gorm:"column:device" json:"device"`
	Activity  string    `gorm:"column:activity" json:"activity"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Deleted   bool      `gorm:"column:deleted;not null" json:"deleted"`
}

// TableName UserLog's table name
func (*UserLog) TableName() string {
	return TableNameUserLog
}