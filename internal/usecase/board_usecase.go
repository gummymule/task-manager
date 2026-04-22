package usecase

import (
	"errors"

	"github.com/gummymule/task-manager/internal/domain"
)

type boardUsecase struct {
	boardRepo domain.BoardRepository
}

func NewBoardUsecase(boardRepo domain.BoardRepository) domain.BoardUsecase {
	return &boardUsecase{boardRepo}
}

func (u *boardUsecase) GetAll(userID int64) ([]domain.Board, error) {
	return u.boardRepo.FindAll(userID)
}

func (u *boardUsecase) GetByID(id, userID int64) (*domain.Board, error) {
	board, err := u.boardRepo.FindByID(id, userID)
	if err != nil {
		return nil, errors.New("board not found")
	}
	return board, nil
}

func (u *boardUsecase) Create(board *domain.Board) (*domain.Board, error) {
	if board.Name == "" {
		return nil, errors.New("board name is required")
	}
	return u.boardRepo.Create(board)
}

func (u *boardUsecase) Update(board *domain.Board) (*domain.Board, error) {
	if board.Name == "" {
		return nil, errors.New("board name is required")
	}
	_, err := u.boardRepo.FindByID(board.ID, board.UserID)
	if err != nil {
		return nil, errors.New("board not found")
	}
	return u.boardRepo.Update(board)
}

func (u *boardUsecase) Delete(id, userID int64) error {
	_, err := u.boardRepo.FindByID(id, userID)
	if err != nil {
		return errors.New("board not found")
	}
	return u.boardRepo.Delete(id, userID)
}