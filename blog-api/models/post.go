package models

type Post struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Title     string `gorm:"not null" json:"title"`
	Content   string `gorm:"type:text;not null" json:"content"`
	AuthorID  uint   `gorm:"not null" json:"author_id"`
	CreatedAt int64  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime" json:"updated_at"`
}
