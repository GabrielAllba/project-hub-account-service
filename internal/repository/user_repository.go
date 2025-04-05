package repository

import (
	"database/sql"
	"errors"

	"project-hub/account-service/internal/domain/entity"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) entity.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetByID(id uint) (*entity.User, error) {
	query := `SELECT id, email, password, created_at, updated_at FROM users WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var user entity.User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*entity.User, error) {
	query := `SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1`
	row := r.db.QueryRow(query, email)

	var user entity.User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Create(user *entity.User) error {
	query := `
		INSERT INTO users (email, password) 
		VALUES ($1, $2) 
		RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(query, user.Email, user.Password).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
