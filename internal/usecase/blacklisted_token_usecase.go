package usecase

import (
	"project-hub/account-service/internal/domain/entity"
)

type blacklistedTokenUseCase struct {
	blacklistedTokenRepo entity.BlacklistedTokenRepository
}

func NewBlacklistedTokenUseCase(repo entity.BlacklistedTokenRepository) entity.BlacklistedTokenUseCase {
	return &blacklistedTokenUseCase{
		blacklistedTokenRepo: repo,
	}
}

func (uc *blacklistedTokenUseCase) BlacklistToken(tokenString string) error {
	err := uc.blacklistedTokenRepo.AddToBlacklisted(tokenString)
	if err != nil {
		return entity.ErrInternal
	}
	return nil
}
