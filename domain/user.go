package domain

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Email    string
	Password string
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
