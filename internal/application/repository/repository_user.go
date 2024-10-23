package repository

import (
	"context"
	"ioc-backend/internal/application/domain"
	"ioc-backend/internal/application/port"
	"ioc-backend/internal/infra/storage/mysql"
)

var _ port.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	db *mysql.Mysql
}

func NewUserRepository(
	db *mysql.Mysql,
) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r UserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
func (r UserRepository) Get(ctx context.Context, id int) (*domain.User, error) {
	var gotUser domain.User
	if err := r.db.WithContext(ctx).First(&gotUser, id).Error; err != nil {
		return nil, err
	}

	return &gotUser, nil
}

func (r UserRepository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := r.db.WithContext(ctx).Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepository) Delete(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&domain.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
