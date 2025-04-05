package controller

import (
	"time"

	"github.com/golang-jwt/jwt/v4"

	"project-hub/account-service/internal/domain/entity"
	"project-hub/account-service/internal/dto/output_schema"
	dtoreq "project-hub/account-service/internal/dto/request"
	userresponse "project-hub/account-service/internal/dto/response/user"
	"project-hub/account-service/internal/dto/schema"
	"project-hub/account-service/pkg/config"
)

type UserController interface {
	GetUser(id uint) (userresponse.GetMeResponse, error)
	Register(req *dtoreq.RegisterRequest) (userresponse.RegisterResponse, error)
	Login(req *dtoreq.LoginRequest) (userresponse.LoginResponse, error)
	Logout(tokenString string) (userresponse.LogoutResponse, error)
	VerifyUser(userID uint) (userresponse.VerifyUserResponse, error)
}

type userController struct {
	userUseCase             entity.UserUseCase
	blacklistedTokenUseCase entity.BlacklistedTokenUseCase
	config                  *config.Config
}

func NewUserController(userUseCase entity.UserUseCase, blacklistedTokenUseCase entity.BlacklistedTokenUseCase, cfg *config.Config) UserController {
	return &userController{
		userUseCase:             userUseCase,
		blacklistedTokenUseCase: blacklistedTokenUseCase,
		config:                  cfg,
	}
}

func (c *userController) GetUser(id uint) (userresponse.GetMeResponse, error) {
	userEntity, err := c.userUseCase.GetUser(id)
	if err != nil {
		return userresponse.GetMeResponse{}, err
	}

	return userresponse.GetMeResponse{
		ErrorSchema: schema.ErrorSchema{
			ErrorCode: "00",
			ErrorMessage: schema.ErrorMessage{
				Indonesian: "Berhasil",
				English:    "Success",
			},
		},
		OutputSchema: entity.User{
			ID:        userEntity.ID,
			Email:     userEntity.Email,
			CreatedAt: userEntity.CreatedAt,
			UpdatedAt: userEntity.UpdatedAt,
		},
	}, nil
}

func (c *userController) Register(req *dtoreq.RegisterRequest) (userresponse.RegisterResponse, error) {
	newUser := &entity.User{
		Email:    req.Email,
		Password: req.Password,
	}

	err := c.userUseCase.CreateUser(newUser)
	if err != nil {
		return userresponse.RegisterResponse{}, err
	}

	createdUser, err := c.userUseCase.GetUserByEmail(newUser.Email)
	if err != nil {
		return userresponse.RegisterResponse{}, err
	}

	return userresponse.RegisterResponse{
		ErrorSchema: schema.ErrorSchema{
			ErrorCode: "00",
			ErrorMessage: schema.ErrorMessage{
				Indonesian: "Pengguna berhasil dibuat",
				English:    "User created successfully",
			},
		},
		OutputSchema: entity.User{
			ID:        createdUser.ID,
			Email:     createdUser.Email,
			CreatedAt: createdUser.CreatedAt,
			UpdatedAt: createdUser.UpdatedAt,
		},
	}, nil
}

func (c *userController) Login(req *dtoreq.LoginRequest) (userresponse.LoginResponse, error) {
	userEntity, err := c.userUseCase.VerifyPassword(req.Email, req.Password)
	if err != nil {
		return userresponse.LoginResponse{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userEntity.ID,
		"email":   userEntity.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(c.config.JWTSecret))
	if err != nil {
		return userresponse.LoginResponse{}, entity.ErrInternal
	}

	return userresponse.LoginResponse{
		ErrorSchema: schema.ErrorSchema{
			ErrorCode: "00",
			ErrorMessage: schema.ErrorMessage{
				Indonesian: "Login berhasil",
				English:    "Login successful",
			},
		},
		OutputSchema: output_schema.UserLoginOutputSchema{
			Token: tokenString,
			User: entity.User{
				ID:        userEntity.ID,
				Email:     userEntity.Email,
				CreatedAt: userEntity.CreatedAt,
				UpdatedAt: userEntity.UpdatedAt,
			},
		},
	}, nil
}

func (c *userController) Logout(tokenString string) (userresponse.LogoutResponse, error) {
	err := c.blacklistedTokenUseCase.BlacklistToken(tokenString)
	if err != nil {
		return userresponse.LogoutResponse{}, err
	}

	return userresponse.LogoutResponse{
		ErrorSchema: schema.ErrorSchema{
			ErrorCode: "00",
			ErrorMessage: schema.ErrorMessage{
				Indonesian: "Logout berhasil",
				English:    "Logout successful",
			},
		},
	}, nil
}

func (c *userController) VerifyUser(userID uint) (userresponse.VerifyUserResponse, error) {
	err := c.userUseCase.VerifyUser(userID)
	if err != nil {
		return userresponse.VerifyUserResponse{}, err
	}

	verifiedUser, err := c.userUseCase.GetUser(userID)
	if err != nil {
		return userresponse.VerifyUserResponse{}, err
	}

	return userresponse.VerifyUserResponse{
		ErrorSchema: schema.ErrorSchema{
			ErrorCode: "00",
			ErrorMessage: schema.ErrorMessage{
				Indonesian: "Verifikasi berhasil",
				English:    "Verification successful",
			},
		},
		OutputSchema: entity.User{
			ID:         verifiedUser.ID,
			Email:      verifiedUser.Email,
			VerifiedAt: verifiedUser.VerifiedAt,
			CreatedAt:  verifiedUser.CreatedAt,
			UpdatedAt:  verifiedUser.UpdatedAt,
		},
	}, nil
}
