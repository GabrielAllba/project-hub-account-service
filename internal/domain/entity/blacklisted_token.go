package entity

import "time"

type BlacklistedToken struct {
	ID            uint      `json:"id"`
	Token         string    `json:"token"`
	BlacklistedAt time.Time `json:"blacklisted_at"`
}

type BlacklistedTokenRepository interface {
	AddToBlacklisted(tokenString string) error
	IsTokenBlacklisted(tokenString string) (bool, error)
}

type BlacklistedTokenUseCase interface {
	BlacklistToken(tokenString string) error
}
