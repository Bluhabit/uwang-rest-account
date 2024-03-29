// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNamePostMention = "post_mention"

// PostMention mapped from table <post_mention>
type PostMention struct {
	ID        string    `gorm:"column:id;primaryKey" json:"id"`
	PostID    string    `gorm:"column:post_id" json:"post_id"`
	UserID    string    `gorm:"column:user_id" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Deleted   bool      `gorm:"column:deleted;not null" json:"deleted"`
}

// TableName PostMention's table name
func (*PostMention) TableName() string {
	return TableNamePostMention
}
