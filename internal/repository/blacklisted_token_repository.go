package repository

import (
	"database/sql"
	"project-hub/account-service/internal/domain/entity"
)

type blacklistedTokenRepository struct {
	db *sql.DB
}

func NewBlacklistedTokenRepository(db *sql.DB) entity.BlacklistedTokenRepository {
	return &blacklistedTokenRepository{
		db: db,
	}
}

func (r *blacklistedTokenRepository) AddToBlacklisted(tokenString string) error {
	query := `INSERT INTO blacklisted_tokens (token) VALUES ($1)`
	_, err := r.db.Exec(query, tokenString)
	if err != nil {
		return err
	}
	return nil
}

func (r *blacklistedTokenRepository) IsTokenBlacklisted(tokenString string) (bool, error) {
	query := `SELECT COUNT(*) FROM blacklisted_tokens WHERE token = $1`
	var count int
	err := r.db.QueryRow(query, tokenString).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
