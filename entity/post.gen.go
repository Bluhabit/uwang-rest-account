// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNamePost = "post"

// Post mapped from table <post>
type Post struct {
	ID            string    `gorm:"column:id;primaryKey" json:"id"`
	CreatedBy     string    `gorm:"column:created_by" json:"created_by"`
	PostID        string    `gorm:"column:post_id" json:"post_id"`
	Body          string    `gorm:"column:body" json:"body"`
	Location      string    `gorm:"column:location" json:"location"`
	LikesCount    int64     `gorm:"column:likes_count;not null" json:"likes_count"`
	CommentsCount int64     `gorm:"column:comments_count;not null" json:"comments_count"`
	PostType      string    `gorm:"column:post_type;not null" json:"post_type"`
	CreatedAt     time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Deleted       bool      `gorm:"column:deleted;not null" json:"deleted"`
}

// TableName Post's table name
func (*Post) TableName() string {
	return TableNamePost
}
