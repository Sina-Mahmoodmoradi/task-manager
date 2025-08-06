package service

import (
    "errors"
    "time"
    "golang.org/x/crypto/bcrypt"

    "github.com/Sina-Mahmoodmoradi/task-manager/internal/domain"
    "github.com/Sina-Mahmoodmoradi/task-manager/internal/core/port"
)



type userService struct {
    userRepo   domain.UserRepository
    tokenManager port.TokenManager
}

func NewUserService(userRepo domain.UserRepository, tokenManager port.TokenManager) domain.UserService {
    return &userService{userRepo: userRepo, tokenManager: tokenManager}
}

func (s *userService) Register(username, password string) (string, error) {
    existing, err := s.userRepo.GetByUsername(username)
    if err == nil && existing != nil {
        return "", errors.New("username already exists")
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }

    user := &domain.User{
        Username: username,
        Password: string(hashedPassword),
    }

    err = s.userRepo.Create(user)
    if err != nil {
        return "", err
    }


	// NOTE: Here you would generate a JWT or other token.
    // For now, we’ll just return a placeholder string.
    token, err := s.tokenManager.CreateToken(user.ID, time.Hour)
    if err != nil {
        return "", err
    }

    return token, nil
}


func (s *userService) Login(username, password string) (string, error) {
    user, err := s.userRepo.GetByUsername(username)
    if err != nil || user == nil {
        return "", errors.New("invalid credentials")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", errors.New("invalid credentials")
    }

    token, err := s.tokenManager.CreateToken(user.ID, time.Hour)
    if err != nil {
        return "", err
    }

    return token, nil
}


func (s *userService) GetByID(id uint) (*domain.User, error) {
    return s.userRepo.GetByID(id)
}


func (s *userService) GetByUsername(username string) (*domain.User, error) {
    return s.userRepo.GetByUsername(username)
}

