package domain

import "time"

type User struct {
	ID       int64     `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"password" db:"password"`
	CreateAt time.Time `json:"create_at" db:"create_at"`
}

type UserRepository interface {
	FindByEmail(email string) (*User, error)
	Create(user *User) (*User, error)
}

type UserUsecase interface {
	Register(user *User) (*User, error)
	Login(email, password string) (string, error)
	Logout() error
}
