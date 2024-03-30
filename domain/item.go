package domain

type Item struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Stock uint
}
