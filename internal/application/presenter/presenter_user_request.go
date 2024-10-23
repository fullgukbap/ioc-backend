package presenter

import (
	"ioc-backend/internal/application/domain"

	"gorm.io/gorm"
)

type UserCreateRequest struct {
	Name string `json:"name"`
}

func (r UserCreateRequest) ToDomain() *domain.User {
	return &domain.User{
		Name: r.Name,
	}
}

type UserGetReqeust struct {
	ID int `query:"id"`
}

type UserUpdateRequest struct {
	ID   uint   `query:"id"`
	Name string `json:"name"`
}

func (r UserUpdateRequest) NewUserUpdateRequest() *domain.User {
	return &domain.User{
		Model: gorm.Model{
			ID: r.ID,
		},
		Name: r.Name,
	}

}

type UserDeleteRequest struct {
	ID int `query:"id"`
}
