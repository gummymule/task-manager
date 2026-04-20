package domain

import "time"

type Task struct {
	ID          int64     `json:"id" db:"id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	BoardID     int64     `json:"board_id" db:"board_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Status      string    `json:"status" db:"status"`
	CreateAt    time.Time `json:"create_at" db:"create_at"`
	UpdateAt    time.Time `json:"update_at" db:"update_at"`
}

type TaskRepository interface {
	FindAll(userID, boardID int64, page, limit int) ([]Task, error)
	FindByID(id, userID int64) (*Task, error)
	Create(task *Task) (*Task, error)
	Update(task *Task) (*Task, error)
	Delete(id, userID int64) error
}

type TaskUsecase interface {
	GetAll(userID, boardID int64, page, limit int) ([]Task, error)
	GetByID(id, userID int64) (*Task, error)
	Create(task *Task) (*Task, error)
	Update(task *Task) (*Task, error)
	Delete(id, userID int64) error
}
