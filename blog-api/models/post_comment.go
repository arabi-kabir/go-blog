package models

import "time"

type PostComment struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT"`
	PostID    uint      `gorm:"not null"`
	Comment   string    `gorm:"type:text;not null" json:"comment"`
	AuthorID  uint      `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}
