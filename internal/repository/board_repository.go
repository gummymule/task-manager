package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/gummymule/task-manager/internal/domain"
)

type boardRepository struct {
	db *sqlx.DB
}

func NewBoardRepository(db *sqlx.DB) domain.BoardRepository {
	return &boardRepository{db}
}

func (r *boardRepository) FindAll(userID int64) ([]domain.Board, error) {
	boards := []domain.Board{}
	query := `
		SELECT id, user_id, name, description, created_at, updated_at
		FROM boards
		WHERE user_id = $1
		ORDER BY created_at DESC
	`
	err := r.db.Select(&boards, query, userID)
	return boards, err
}

func (r *boardRepository) FindByID(id, userID int64) (*domain.Board, error) {
	board := &domain.Board{}
	query := `
		SELECT id, user_id, name, description, created_at, updated_at
		FROM boards WHERE id = $1 AND user_id = $2
	`
	err := r.db.Get(board, query, id, userID)
	if err != nil {
		return nil, err
	}
	return board, nil
}

func (r *boardRepository) Create(board *domain.Board) (*domain.Board, error) {
	query := `
		INSERT INTO boards (user_id, name, description)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, name, description, created_at, updated_at
	`

	result := &domain.Board{}
	err := r.db.QueryRowx(query, board.UserID, board.Name, board.Description).StructScan(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *boardRepository) Update(board	*domain.Board) (*domain.Board, error) {
	query := `
		UPDATE boards
		SET name = $1, description = $2, updated_at = NOW()
		WHERE id = $3 AND user_id = $4
		RETURNING id, user_id, name, description, created_at, updated_at
	`
	result := &domain.Board{}
	err := r.db.QueryRowx(query, board.Name, board.Description, board.ID, board.UserID).StructScan(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *boardRepository) Delete(id, userID int64) error {
	query := `DELETE FROM boards WHERE id = $1 AND user_id = $2`
	_, err := r.db.Exec(query, id, userID)
	return err
}