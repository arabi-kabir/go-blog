package models

type PostCategory struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `gorm:"not null" json:"title"`

	// One category has many posts
	Posts []Post `gorm:"foreignKey:PostCategoryID" json:"posts"`
}
