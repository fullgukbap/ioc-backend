package port

import (
	"context"
	"ioc-backend/internal/application/domain"
)

type UserService interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	Get(ctx context.Context, id int) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) (*domain.User, error)
	Delete(ctx context.Context, id int) error
}

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	Get(ctx context.Context, id int) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) (*domain.User, error)
	Delete(ctx context.Context, id int) error
}
