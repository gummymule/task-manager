package domain

import "time"

type Board struct {
	ID          int64     `json:"id" db:"id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type BoardRepository interface {
	FindAll(userID int64) ([]Board, error)
	FindByID(id, userID int64) (*Board, error)
	Create(board *Board) (*Board, error)
	Update(board *Board) (*Board, error)
	Delete(id, userID int64) error
}

type BoardUsecase interface {
	GetAll(userID int64) ([]Board, error)
	GetByID(id, userID int64) (*Board, error)
	Create(board *Board) (*Board, error)
	Update(board *Board) (*Board, error)
	Delete(id, userID int64) error
}
