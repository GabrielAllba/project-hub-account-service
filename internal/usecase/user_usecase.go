package usecase

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	"project-hub/account-service/internal/domain/entity"
)

type userUseCase struct {
	userRepo entity.UserRepository
}

func NewUserUseCase(repo entity.UserRepository) entity.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (uc *userUseCase) GetUser(id uint) (*entity.User, error) {
	user, err := uc.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, entity.ErrNotFound
	}
	return user, nil
}

func (uc *userUseCase) GetUserByEmail(email string) (*entity.User, error) {
	user, err := uc.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, entity.ErrNotFound
	}
	return user, nil
}

func (uc *userUseCase) CreateUser(user *entity.User) error {

	existingUser, err := uc.userRepo.GetByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return entity.ErrAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.ErrInternal
	}
	user.Password = string(hashedPassword)

	return uc.userRepo.Create(user)
}

func (uc *userUseCase) VerifyPassword(email, password string) (*entity.User, error) {
	user, err := uc.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, entity.ErrUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, entity.ErrUnauthorized
	}

	return user, nil
}

func (uc *userUseCase) VerifyUser(userID uint) error {
	user, err := uc.userRepo.GetByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return entity.ErrNotFound
	}

	now := time.Now()
	user.VerifiedAt = &now

	return uc.userRepo.Update(user)
}
