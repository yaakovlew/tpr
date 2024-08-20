package api_common_service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

const (
	tokenTTL = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId   int    `json:"user_id"`
	UserPost string `json:"user_post"`
}

type tokenClaimsRestore struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	cost := 10
	bPassword := []byte(password)
	hash, _ := bcrypt.GenerateFromPassword(bPassword, cost)

	return string(hash)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(email)
	if err != nil {
		err = errors.New("person isn't found")
		return "", err
	}
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		err = errors.New("person isn't found")
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.Post,
	})
	return token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
}

func (s *AuthService) GenerateTokenByUserId(userId int) (string, error) {
	email, err := s.repo.GetUserMail(userId)
	if err != nil {
		err = errors.New("person isn't found")
		return "", err
	}
	user, err := s.repo.GetUser(email)
	if err != nil {
		err = errors.New("person isn't found")
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.Post,
	})
	return token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
}

func (s *AuthService) GetPost(email string) (string, error) {
	user, err := s.repo.GetUser(email)
	if err != nil {
		err = errors.New("person isn't found")
		return "", err
	}
	return user.Post, nil
}

func (s *AuthService) GenerateTokenForRestorePassword(email string) (string, error) {
	user, err := s.repo.GetUser(email)
	if err != nil {
		err = errors.New("person isn't found")
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaimsRestore{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
}

func (s *AuthService) ParseToken(accessToken string) (int, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(os.Getenv("SIGNING_KEY")), nil
	})
	if err != nil {
		return 0, "", err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, "", errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, claims.UserPost, nil
}

func (s *AuthService) ParseTokenForRestorePassword(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaimsRestore{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(os.Getenv("SIGNING_KEY")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaimsRestore)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}

func (s *AuthService) ChangePassword(userId int, changePassword model.ChangePasswordInput) error {
	email, err := s.repo.GetUserMail(userId)
	errAccess := errors.New("person isn't found")
	if err != nil {
		return errAccess
	}
	user, err := s.repo.GetUser(email)
	if err != nil {
		return errAccess
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(changePassword.OldPassword))
	if err != nil {
		return errAccess
	}
	newPassword := s.generatePasswordHash(changePassword.NewPassword)
	err = s.repo.ChangePassword(userId, newPassword)
	if err != nil {
		return errAccess
	}
	return nil
}

func (s *AuthService) GeneratePasswordResetLink(email string) (*model.PasswordResetLink, error) {
	token, err := s.GenerateTokenForRestorePassword(email)
	if err != nil {
		return nil, err
	}
	expiry := time.Now().Add(24 * time.Hour)
	link := &model.PasswordResetLink{
		Email:  email,
		Token:  token,
		Expiry: expiry,
	}
	return link, nil
}

func (s *AuthService) ChangePasswordForStudent(studentId int, newPassword string) error {
	hash := s.generatePasswordHash(newPassword)
	return s.repo.ChangePassword(studentId, hash)
}

func (s *AuthService) RestorePassword(userId int, password string) error {
	newPass := s.generatePasswordHash(password)
	return s.repo.ChangePassword(userId, newPass)
}
