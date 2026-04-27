package services

import (
	"errors"
	"os"
	"time"

	"go-blog/models"
	"go-blog/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// AuthService defines authentication operations.
type AuthService interface {
	Register(name, email, password string) (*models.User, error)
	Login(username, password string) (string, error)
}

type authService struct {
	repo repositories.UserRepository
}

// NewAuthService returns an AuthService backed by the provided repository.
func NewAuthService(repo repositories.UserRepository) AuthService {
	return &authService{repo: repo}
}

// HashPassword hashes a plain-text password using bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword validates a plain-text password against its bcrypt hash.
func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateJWT creates a signed JWT for the given user.
func GenerateJWT(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (s *authService) Register(username, email, password string) (*models.User, error) {
	hashed, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: hashed,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(username, password string) (string, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !CheckPassword(user.Password, password) {
		return "", errors.New("invalid credentials")
	}

	return GenerateJWT(*user)
}
