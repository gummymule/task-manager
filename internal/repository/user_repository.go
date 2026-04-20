package repository

import (
	"github.com/gummymule/task-manager/internal/domain"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, name, email, password, create_at FROM users WHERE email = $1`
	err := r.db.Get(user, query, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Create(user *domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (name, email, password) 
		VALUES ($1, $2, $3) 
		RETURNING id, name, email, password, create_at
	`
	result := &domain.User{}
	err := r.db.QueryRowx(query, user.Name, user.Email, user.Password).StructScan(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
