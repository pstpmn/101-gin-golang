package models
import (
	"time"
)
type Articles struct {
	Aip            int 		`gorm:"column:aid;primary_key"`
	Topic          string 	`gorm:"column:topic;"`
	Content        string   `gorm:"column:content;"`
	Draft          string   `gorm:"column:draft;"`
	Hashtag        string   `gorm:"type:json;column:hashtag;"`
	Views          int      `gorm:"column:views;default:0"`
	UpdatedAt     time.Time `gorm:"column:updatedAt;"`
	CreatedAt	  time.Time `gorm:"column:createdAt;default:current_timestamp;"`
}