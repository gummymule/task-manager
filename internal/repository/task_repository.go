package repository

import (
	"github.com/gummymule/task-manager/internal/domain"
	"github.com/jmoiron/sqlx"
)

type taskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) domain.TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) FindAll(userID int64, page, limit int) ([]domain.Task, error) {
	tasks := []domain.Task{}
	offset := (page - 1) * limit
	query := `
		SELECT id, user_id, title, description, status, create_at, update_at
		FROM tasks
		WHERE user_id = $1
		ORDER BY create_at DESC
		LIMIT $2 OFFSET $3
	`
	err := r.db.Select(&tasks, query, userID, limit, offset)
	return tasks, err
}

func (r *taskRepository) FindByID(id, userID int64) (*domain.Task, error) {
	task := &domain.Task{}
	query := `
		SELECT id, user_id, title, description, status, create_at, update_at
		FROM tasks
		WHERE id = $1 AND user_id = $2
	`
	err := r.db.Get(task, query, id, userID)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *taskRepository) Create(task *domain.Task) (*domain.Task, error) {
	query := `
		INSERT INTO tasks (user_id, title, description, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, user_id, title, description, status, create_at, update_at
	`
	result := &domain.Task{}
	err := r.db.QueryRowx(query, task.UserID, task.Title, task.Description, task.Status).StructScan(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *taskRepository) Update(task *domain.Task) (*domain.Task, error) {
	query := `
		UPDATE tasks 
		SET title = $1, description = $2, status = $3, update_at = NOW()
		WHERE id = $4 AND user_id = $5
		RETURNING id, user_id, title, description, status, create_at, update_at
	`
	result := &domain.Task{}
	err := r.db.QueryRowx(query, task.Title, task.Description, task.Status, task.ID, task.UserID).StructScan(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *taskRepository) Delete(id, userID int64) error {
	query := `DELETE FROM tasks WHERE id = $1 AND user_id = $2`
	_, err := r.db.Exec(query, id, userID)
	return err
}
