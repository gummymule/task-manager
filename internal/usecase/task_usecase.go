package usecase

import (
	"errors"

	"github.com/gummymule/task-manager/internal/domain"
)

type taskUsecase struct {
	taskRepo domain.TaskRepository
}

func NewTaskUsecase(taskRepo domain.TaskRepository) domain.TaskUsecase {
	return &taskUsecase{taskRepo}
}

func (u *taskUsecase) GetAll(userID int64, page, limit int) ([]domain.Task, error) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > 100 {
		limit = 10
	}
	return u.taskRepo.FindAll(userID, page, limit)
}

func (u *taskUsecase) GetByID(id, userID int64) (*domain.Task, error) {
	task, err := u.taskRepo.FindByID(id, userID)
	if err != nil {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (u *taskUsecase) Create(task *domain.Task) (*domain.Task, error) {
	if task.Title == "" {
		return nil, errors.New("title is required")
	}
	if task.Status == "" {
		task.Status = "to_do"
	}
	return u.taskRepo.Create(task)
}

func (u *taskUsecase) Update(task *domain.Task) (*domain.Task, error) {

	if task.Title == "" {
		return nil, errors.New("title is required")
	}

	validStatus := map[string]bool{"to_do": true, "in_progress": true, "completed": true}
	if !validStatus[task.Status] {
		return nil, errors.New("status must be to_do, in_progress, or completed")
	}

	_, err := u.taskRepo.FindByID(task.ID, task.UserID)
	if err != nil {
		return nil, errors.New("task not found")
	}

	return u.taskRepo.Update(task)
}

func (u *taskUsecase) Delete(id, userID int64) error {
	_, err := u.taskRepo.FindByID(id, userID)
	if err != nil {
		return errors.New("task not found")
	}
	return u.taskRepo.Delete(id, userID)
}
