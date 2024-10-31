package services

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rwx03/Pastebin/backend/internal/models"
	"github.com/rwx03/Pastebin/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	jwtAccessSecret  = []byte("access-secret")
	jwtRefreshSecret = []byte("refresh-secret")
	accessTokenTTL   = 15 * time.Minute
	refreshTokenTTL  = 30 * 24 * 60 * 60 * 1000
)

type tokenClaims struct {
	jwt.Claims
	Email string `json:"email"`
}

type AuthService struct {
	tokenRepo repository.Token
	userRepo  repository.User
}

func NewAuthService(userRepo repository.User, tokenRepo repository.Token) *AuthService {
	return &AuthService{
		tokenRepo: tokenRepo,
		userRepo:  userRepo,
	}
}

func (a *AuthService) Register(email, password string) (string, string, error) {
	ctx := context.Background()

	candidate, err := a.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}

	if candidate != nil {
		return "", "", errors.New("user with this email already exists")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 3)
	if err != nil {
		return "", "", err
	}

	userID, err := a.userRepo.Create(ctx, models.User{
		Email:     email,
		Password:  string(hashPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return "", "", err
	}

	accessToken, refreshToken, err := generateTokens(email)
	if err != nil {
		return "", "", err
	}

	if _, err := a.tokenRepo.Create(ctx, models.Token{
		UserID:       userID,
		RefreshToken: refreshToken,
	}); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (a *AuthService) Login(email, password string) (string, string, error) {
	ctx := context.Background()

	user, err := a.userRepo.GetByEmail(ctx, email)
	if err != nil || user == nil {
		return "", "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", errors.New("invalid password")
	}

	accessToken, refreshToken, err := generateTokens(email)
	if err != nil {
		return "", "", err
	}

	if err := a.tokenRepo.UpdateByID(ctx, user.ID, refreshToken); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (a *AuthService) Refresh(refreshToken string) (string, string, error) {
	ctx := context.Background()

	email, err := validateToken(refreshToken, false)
	if err != nil {
		return "", "", err
	}

	accessToken, newRefreshToken, err := generateTokens(email)
	if err != nil {
		return "", "", err
	}

	if err := a.tokenRepo.UpdateByToken(ctx, refreshToken, newRefreshToken); err != nil {
		return "", "", err
	}

	return accessToken, newRefreshToken, nil
}

func generateTokens(email string) (string, string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		email,
	})

	accessTokenString, err := accessToken.SignedString(jwtAccessSecret)
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		email,
	})

	refreshTokenString, err := refreshToken.SignedString(jwtRefreshSecret)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func validateToken(tokenString string, isAccessToken bool) (string, error) {
	var secret []byte

	if isAccessToken {
		secret = jwtAccessSecret
	} else {
		secret = jwtRefreshSecret
	}

	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		return claims.Email, nil
	}

	return "", errors.New("invalid token")
}
