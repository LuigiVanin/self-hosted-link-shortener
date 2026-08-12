package entities

import (
	"encoding/json"
	"time"
)

type Link struct {
	ID          int64  `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:""`
	Description string `json:"description" gorm:""`
	Url         string `json:"url" gorm:""`
	Code        string `json:"code" gorm:"unique;index"`

	Metadata json.RawMessage `json:"metada"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	Body      string    `json:"body"`
}
