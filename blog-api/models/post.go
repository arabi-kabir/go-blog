package models

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Foreign key
	AuthorID       uint `gorm:"not null" json:"author_id"`
	PostCategoryID uint `gorm:"not null" json:"post_category_id"`

	// Belongs to category
	Author   User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"author"`
	Category PostCategory `gorm:"foreignKey:PostCategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"category"`
}
