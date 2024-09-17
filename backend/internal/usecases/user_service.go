package usecase

import (
    "errors"
    "memorandum-backend/internal/entities"
    "golang.org/x/crypto/bcrypt"
)

type UserService struct {
    userRepo     entities.UserRepository
    tokenService entities.TokenService
}

func NewUserService(userRepo entities.UserRepository, tokenService entities.TokenService) *UserService {
    return &UserService{userRepo: userRepo, tokenService: tokenService}
}

func (s *UserService) Register(user *entities.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    return s.userRepo.Create(user)
}

func (s *UserService) Login(email, password string) (string, error) {
    user, err := s.userRepo.GetByEmail(email)
    if err != nil {
        return "", errors.New("invalid email or password")
    }
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return "", errors.New("invalid email or password")
    }
    return s.tokenService.GenerateToken(user)
}

func (s *UserService) GetUser(id uint) (*entities.User, error) {
    return s.userRepo.GetByID(id)
}

func (s *UserService) UpdateUser(user *entities.User) error {
    return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
    return s.userRepo.Delete(id)
}