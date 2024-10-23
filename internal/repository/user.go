package repository

import (
	"context"
	"github.com/sunviv/k8s-demo/internal/domain"
	"github.com/sunviv/k8s-demo/internal/repository/dao"
)

type UserRepository struct {
	dao *dao.UserDao
}

func NewUserRepository(dao *dao.UserDao) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (repo *UserRepository) Create(ctx context.Context, user domain.User) error {
	return repo.dao.Insert(ctx, dao.User{
		Email:    user.Email,
		Password: user.Password,
	})
}

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := repo.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomain(u), nil
}

func (repo *UserRepository) toDomain(u dao.User) domain.User {
	return domain.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	}
}
