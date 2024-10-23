package service

import (
	"context"
	"ioc-backend/internal/application/domain"
	"ioc-backend/internal/application/port"
)

var _ port.UserService = (*UserService)(nil)

type UserService struct {
	userRepository port.UserRepository
}

func NewUserService(
	userRepository port.UserRepository,
) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s UserService) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	createdUser, err := s.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s UserService) Get(ctx context.Context, id int) (*domain.User, error) {
	gotUser, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return gotUser, nil
}

func (s UserService) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	updatedUser, err := s.userRepository.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s UserService) Delete(ctx context.Context, id int) error {
	err := s.userRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
