package user

import (
	"context"

	"CtrlAltDestiny/internal/entity"
)

//go:generate go run go.uber.org/mock/mockgen -source=interface.go -destination=manager_mock_test.go -package=user Manager
type Manager interface {
	GetUser(ctx context.Context, id int) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) error
	UpdateUser(ctx context.Context, user entity.User) error
	DeleteUser(ctx context.Context, id int) error
}
