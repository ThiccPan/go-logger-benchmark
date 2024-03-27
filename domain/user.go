package domain

type User struct {
	Username string
	Email    string
	Password string
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
