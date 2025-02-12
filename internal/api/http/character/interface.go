package character

import (
	"context"

	"CtrlAltDestiny/internal/entity"
)

//go:generate go run go.uber.org/mock/mockgen -source=interface.go -destination=manager_mock_test.go -package=user Manager
type Usecases interface {
	GetCharacter(ctx context.Context, id int) (entity.Character, error)
	CreateCharacter(ctx context.Context, character entity.Character) error
	UpdateCharacter(ctx context.Context, character entity.Character) error
	DeleteCharacter(ctx context.Context, id int) error
}
