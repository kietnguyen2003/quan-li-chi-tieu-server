package auth

import (
	"fmt"
	domainAuth "quan-li-chi-tieu/src/domain/auth"
)

type AuthUseCase struct {
	userRepo       UserRepository
	passwordHasher PasswordHasher
	tokenProvider  TokenProvider
}

func NewAuthUseCase(userRepo UserRepository, passwordHasher PasswordHasher, tokenProvider TokenProvider) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
		tokenProvider:  tokenProvider,
	}
}

func (uc *AuthUseCase) Login(req LoginRequest) (*AuthResponse, error) {
	user, err := uc.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	err = uc.passwordHasher.Compare(user.Password, req.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	token, err := uc.tokenProvider.Generate(user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token")
	}

	return &AuthResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		UserID:       fmt.Sprintf("%d", user.ID),
	}, nil
}

func (uc *AuthUseCase) Register(req RegisterRequest) (*AuthResponse, error) {
	hashedPassword, err := uc.passwordHasher.Hash(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &domainAuth.User{
		Email:    req.Email,
		FullName: req.FullName,
		Password: hashedPassword,
	}

	userID, err := uc.userRepo.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	token, err := uc.tokenProvider.Generate(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &AuthResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		UserID:       fmt.Sprintf("%d", userID),
	}, nil
}
