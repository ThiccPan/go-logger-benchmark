package helper

import "github.com/golang-jwt/jwt/v4"

type AuthJWT struct {
	secret string
}

func NewJWTGen(secret string) *AuthJWT {
	return &AuthJWT{
		secret: secret,
	}
}

type JwtCustomClaims struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (aj *AuthJWT) GenerateToken(
	id uint,
	username,
	email,
	password string,
) (string, error) {
	claims := JwtCustomClaims{
		ID:       id,
		Username: username,
		Email:    email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(aj.secret))
}
