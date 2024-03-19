package domain

type Post struct {
	ID      uint `gorm:"primaryKey"`
	Title   string
	Content string
}
