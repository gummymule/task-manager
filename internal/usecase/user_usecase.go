package usecase

import (
	"errors"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gummymule/task-manager/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo  domain.UserRepository
	jwtSecret string
}

func NewUserUsecase(userRepo domain.UserRepository, jwtSecret string) domain.UserUsecase {
	return &userUsecase{userRepo, jwtSecret}
}

func (u *userUsecase) Register(user *domain.User) (*domain.User, error) {
	existing, _ := u.userRepo.FindByEmail(user.Email)
	if existing != nil {
		return nil, errors.New("email already registered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashed)

	return u.userRepo.Create(user)
}

func (u *userUsecase) Login(email, password string) (string, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (u *userUsecase) Logout() error {
	return nil
}
