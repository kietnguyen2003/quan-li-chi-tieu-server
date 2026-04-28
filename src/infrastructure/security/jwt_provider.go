package security

import (
	"errors"
	appAuth "quan-li-chi-tieu/src/application/auth"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTProvider struct {
	secret string
}

func NewJWTProvider(secret string) *JWTProvider {
	return &JWTProvider{secret: secret}
}

func (p *JWTProvider) Generate(userID uint) (*appAuth.TokenResponse, error) {

	claimsAccess := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(10 * time.Minute).Unix(),
	}

	claimsRefresh := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccess)
	signedString, err := accessToken.SignedString([]byte(p.secret))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	signedRefreshString, err := refreshToken.SignedString([]byte(p.secret))
	if err != nil {
		return nil, err
	}

	return &appAuth.TokenResponse{
		AccessToken:  signedString,
		RefreshToken: signedRefreshString,
	}, nil
}

func (p *JWTProvider) Validate(tokenString string) (*appAuth.TokenClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(p.secret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	userIDValue, ok := claims["user_id"].(float64)
	if !ok {
		return nil, errors.New("invalid user_id claim")
	}

	return &appAuth.TokenClaims{
		UserID: uint(userIDValue),
	}, nil
}
