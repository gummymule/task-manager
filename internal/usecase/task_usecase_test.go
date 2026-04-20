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

func TestCreateTask_Success(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)

	mockRepo.On("Create", mock.AnythingOfType("*domain.Task")).Return(&domain.Task{
		ID:          1,
		UserID:      1,
		Title:       "Test Task",
		Description: "Test Description",
		Status:      "to_do",
	}, nil)

	uc := usecase.NewTaskUsecase(mockRepo)
	result, err := uc.Create(&domain.Task{
		UserID:      1,
		Title:       "Test Task",
		Description: "Test Description",
	})

	assert.NoError(t, err)
	assert.Equal(t, "Test Task", result.Title)
	assert.Equal(t, "to_do", result.Status)
	mockRepo.AssertExpectations(t)
}

func TestCreateTask_EmptyTitle(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)

	uc := usecase.NewTaskUsecase(mockRepo)
	result, err := uc.Create(&domain.Task{
		UserID: 1,
		Title:  "",
	})

	assert.Error(t, err)
	assert.Equal(t, "title is required", err.Error())
	assert.Nil(t, result)
}

func TestGetByID_Success(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)

	mockRepo.On("FindByID", int64(1), int64(1)).Return(&domain.Task{
		ID:     1,
		UserID: 1,
		Title:  "Test Task",
		Status: "to_do",
	}, nil)

	uc := usecase.NewTaskUsecase(mockRepo)
	result, err := uc.GetByID(1, 1)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), result.ID)
	assert.Equal(t, "Test Task", result.Title)
	mockRepo.AssertExpectations(t)
}

func TestGetByID_NotFound(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)

	mockRepo.On("FindByID", int64(99), int64(1)).Return(nil, errors.New("not found"))

	uc := usecase.NewTaskUsecase(mockRepo)
	result, err := uc.GetByID(99, 1)

	assert.Error(t, err)
	assert.Equal(t, "task not found", err.Error())
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTask_InvalidStatus(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)

	uc := usecase.NewTaskUsecase(mockRepo)
	result, err := uc.Update(&domain.Task{
		ID:     1,
		UserID: 1,
		Title:  "Test Task",
		Status: "invalid_status",
	})

	assert.Error(t, err)
	assert.Equal(t, "status must be to_do, in_progress, or completed", err.Error())
	assert.Nil(t, result)
}

func TestDeleteTask_Success(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)

	mockRepo.On("FindByID", int64(1), int64(1)).Return(&domain.Task{
		ID:     1,
		UserID: 1,
	}, nil)
	mockRepo.On("Delete", int64(1), int64(1)).Return(nil)

	uc := usecase.NewTaskUsecase(mockRepo)
	err := uc.Delete(1, 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask_NotFound(t *testing.T) {
	mockRepo := new(mocks.TaskRepository)

	mockRepo.On("FindByID", int64(99), int64(1)).Return(nil, errors.New("not found"))

	uc := usecase.NewTaskUsecase(mockRepo)
	err := uc.Delete(99, 1)

	assert.Error(t, err)
	assert.Equal(t, "task not found", err.Error())
	mockRepo.AssertExpectations(t)
}
