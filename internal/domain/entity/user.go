package entity

import "time"

type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	GetByID(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	Create(user *User) error
}

type UserUseCase interface {
	GetUser(id uint) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) error
	VerifyPassword(email, password string) (*User, error)
}
