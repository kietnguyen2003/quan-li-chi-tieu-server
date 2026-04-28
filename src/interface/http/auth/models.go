package auth

import (
	authApp "quan-li-chi-tieu/src/application/auth"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerRequest struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}
type authHTTPResponse struct {
	UserID       string `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func toLoginApplication(req loginRequest) authApp.LoginRequest {
	return authApp.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}
}

func toRegisterApplication(req registerRequest) authApp.RegisterRequest {
	return authApp.RegisterRequest{
		Email:    req.Email,
		FullName: req.FullName,
		Password: req.Password,
	}
}

func toAuthHTTPResponse(resp *authApp.AuthResponse) authHTTPResponse {
	return authHTTPResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		UserID:       resp.UserID,
	}
}
