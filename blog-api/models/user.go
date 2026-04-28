package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`

	// One user has many posts
	Posts []Post `gorm:"foreignKey:AuthorID" json:"posts"`
}
