package service

import (
	"crypto/sha1"
	"fmt"
	todo "github.com/TelitsynNikita"
	"github.com/TelitsynNikita/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "asdfag4938hdn32casdk"
	tokenTTL   = 12 * time.Hour
	signingKey = "adsf43fasdfadfsf"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId       int    `json:"user_id"`
	UserEmail    string `json:"user_email"`
	UserRole     string `json:"user_role"`
	UserFullName string `json:"full_name"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(user todo.User) (string, error) {
	user, err := s.repo.GetUser(user)
	if err != nil {
		return "", nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.Email,
		user.Role,
		user.FullName,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
