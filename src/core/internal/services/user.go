package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mujak27/gamen/src/core/internal/api/dto"
	interfaces "github.com/mujak27/gamen/src/core/internal/interfaces/repository"
	"github.com/mujak27/gamen/src/core/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository interfaces.UserRepository
}

func NewUserService(userRepository interfaces.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (u *UserService) Login(request dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := u.userRepository.FindByEmail(request.Email)
	if err != nil {
		return nil, err
	}
	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password)); err != nil {
		return nil, errors.New("invalid password")
	}
	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}
	return &dto.LoginResponse{
		Status: "success",
		Token:  tokenString,
	}, nil
}

func (u *UserService) Register(request dto.RegisterRequest) (*models.User, error) {
	existingUser, err := u.userRepository.FindByEmail(request.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		ID:           uuid.New(),
		Email:        request.Email,
		Username:     request.Username,
		PasswordHash: string(hashedPassword),
		FirstName:    request.FirstName,
		LastName:     request.LastName,
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := u.userRepository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
