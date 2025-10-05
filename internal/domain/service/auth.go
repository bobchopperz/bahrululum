package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bobchopperz/bahrululum/internal/config"
	"github.com/bobchopperz/bahrululum/internal/domain/models"
	"github.com/bobchopperz/bahrululum/internal/domain/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, req *models.LoginRequest) (*models.TokenResponse, error)
	ValidateToken(tokenString string) (*Claims, error)
	GenerateToken(userID uint) (*models.TokenResponse, error)
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

type authService struct {
	userRepo  repository.UserRepository
	jwtConfig *config.JWTConfig
}

func NewAuthService(userRepo repository.UserRepository, jwtConfig *config.JWTConfig) AuthService {
	return &authService{
		userRepo:  userRepo,
		jwtConfig: jwtConfig,
	}
}

func (s *authService) Login(ctx context.Context, req *models.LoginRequest) (*models.TokenResponse, error) {
	user, err := s.userRepo.GetByNip(ctx, req.Nip)
	if err != nil {
		return nil, fmt.Errorf("Invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("Invalid credentials")
	}

	if !user.IsActive {
		return nil, fmt.Errorf("user account is inactive")
	}

	token, err := s.GenerateToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token")
	}

	token.User = user.ToResponse()

	return token, nil
}

func (s *authService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.jwtConfig.Secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token claims")
}

func (s *authService) GenerateToken(userID uint) (*models.TokenResponse, error) {
	accessClaimns := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.jwtConfig.Expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-rest-api",
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaimns)
	accessTokenString, err := accessToken.SignedString([]byte(s.jwtConfig.Secret))
	if err != nil {
		return nil, fmt.Errorf("filed to sign access token: %w", err)
	}

	refreshClaims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.jwtConfig.RefreshExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-rest-api",
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(s.jwtConfig.Secret))
	if err != nil {
		return nil, fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return &models.TokenResponse{
		AccessToken: accessTokenString,

		RefreshToken: refreshTokenString,
		ExpiresAt:    time.Now().Add(s.jwtConfig.Expiry).Unix(),
		TokenType:    "Bearer",
	}, nil
}
