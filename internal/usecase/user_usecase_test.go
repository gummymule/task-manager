package usecase_test

import (
	"errors"
	"testing"

	"github.com/gummymule/task-manager/internal/domain"
	"github.com/gummymule/task-manager/internal/mocks"
	"github.com/gummymule/task-manager/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	mockRepo.On("FindByEmail", "john@example.com").Return(nil, errors.New("not found"))
	mockRepo.On("Create", mock.AnythingOfType("*domain.User")).Return(&domain.User{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	}, nil)

	uc := usecase.NewUserUsecase(mockRepo, "secret")
	result, err := uc.Register(&domain.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	})

	assert.NoError(t, err)
	assert.Equal(t, "John Doe", result.Name)
	assert.Equal(t, "john@example.com", result.Email)
	mockRepo.AssertExpectations(t)
}

func TestRegister_EmailAlreadyExists(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	mockRepo.On("FindByEmail", "john@example.com").Return(&domain.User{
		ID:    1,
		Email: "john@example.com",
	}, nil)

	uc := usecase.NewUserUsecase(mockRepo, "secret")
	result, err := uc.Register(&domain.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	})

	assert.Error(t, err)
	assert.Equal(t, "email already registered", err.Error())
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestLogin_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	// Generate hashed password untuk test
	hashedPassword := "$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi" // "password"

	mockRepo.On("FindByEmail", "john@example.com").Return(&domain.User{
		ID:       1,
		Email:    "john@example.com",
		Password: hashedPassword,
	}, nil)

	uc := usecase.NewUserUsecase(mockRepo, "secret")
	token, err := uc.Login("john@example.com", "password")

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestLogin_WrongPassword(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	hashedPassword := "$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi" // "password"

	mockRepo.On("FindByEmail", "john@example.com").Return(&domain.User{
		ID:       1,
		Email:    "john@example.com",
		Password: hashedPassword,
	}, nil)

	uc := usecase.NewUserUsecase(mockRepo, "secret")
	token, err := uc.Login("john@example.com", "wrongpassword")

	assert.Error(t, err)
	assert.Equal(t, "invalid email or password", err.Error())
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestLogin_EmailNotFound(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	mockRepo.On("FindByEmail", "notfound@example.com").Return(nil, errors.New("not found"))

	uc := usecase.NewUserUsecase(mockRepo, "secret")
	token, err := uc.Login("notfound@example.com", "password123")

	assert.Error(t, err)
	assert.Equal(t, "invalid email or password", err.Error())
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}
