// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameHashtag = "hashtag"

// Hashtag mapped from table <hashtag>
type Hashtag struct {
	Hashtag   string    `gorm:"column:hashtag" json:"hashtag"`
	ID        string    `gorm:"column:id;primaryKey" json:"id"`
	Value     string    `gorm:"column:value" json:"value"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Deleted   bool      `gorm:"column:deleted;not null" json:"deleted"`
}

// TableName Hashtag's table name
func (*Hashtag) TableName() string {
	return TableNameHashtag
}
