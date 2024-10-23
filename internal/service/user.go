package service

import (
	"context"
	"errors"
	"github.com/sunviv/k8s-demo/internal/domain"
	"github.com/sunviv/k8s-demo/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrUserOrPasswordInvalid = errors.New("invalid user/password")
	ErrEmailDuplicate        = errors.New("email duplicate")
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s UserService) SignUp(ctx context.Context, user domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrUserOrPasswordInvalid
	}
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return s.repo.Create(ctx, user)
}

func (s UserService) SignIn(ctx context.Context, email, password string) (domain.User, error) {
	u, err := s.repo.FindByEmail(ctx, email)
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return domain.User{}, ErrUserOrPasswordInvalid
	}
	if err != nil {
		return domain.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return domain.User{}, ErrUserOrPasswordInvalid
	}
	return u, nil
}
