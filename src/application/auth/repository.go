package auth

import (
	domainAuth "quan-li-chi-tieu/src/domain/auth"
)

type UserRepository interface {
	GetUserByEmail(email string) (*domainAuth.User, error)
	CreateUser(user *domainAuth.User) (uint, error)
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}

type TokenProvider interface {
	Generate(userID uint) (*TokenResponse, error)
	Validate(token string) (*TokenClaims, error)
}

type TokenClaims struct {
	UserID uint
}

type TokenResponse struct {
	AccessToken  string
	RefreshToken string
}
