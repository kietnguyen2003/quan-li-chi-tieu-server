package auth

type LoginRequest struct {
	Email    string
	Password string
}

type RegisterRequest struct {
	Email    string
	FullName string
	Password string
}

type AuthResponse struct {
	AccessToken  string
	RefreshToken string
	UserID       string
}
