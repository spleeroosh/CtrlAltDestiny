package userinfo

import (
	"context"
	"fmt"

	"CtrlAltDestiny/internal/entity"
	"CtrlAltDestiny/internal/pkg/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	repo   Repository
	log    log.Logger
	tracer trace.Tracer
}

func NewService(repo Repository, logger log.Logger) *Service {
	return &Service{
		repo:   repo,
		log:    log.WithName(logger, "userinfo"),
		tracer: otel.Tracer("Service"),
	}
}

func (s *Service) GetUser(ctx context.Context, id int) (entity.User, error) {
	ctx, span := s.tracer.Start(ctx, "Service:GetUser()")
	defer span.End()

	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return entity.User{}, fmt.Errorf("get user: %w", err)
	}

	return user, nil
}

func (s *Service) CreateUser(ctx context.Context, user entity.User) error {
	ctx, span := s.tracer.Start(ctx, "Service:CreateUser()")
	defer span.End()

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("save user: %w", err)
	}

	s.log.Info().Str("userName", user.Name).Msg("user has been created")

	return nil
}

func (s *Service) UpdateUser(ctx context.Context, user entity.User) error {
	ctx, span := s.tracer.Start(ctx, "Service:UpdateUser()")
	defer span.End()

	if err := s.repo.UpdateUser(ctx, user); err != nil {
		return fmt.Errorf("save user: %w", err)
	}

	return nil
}

func (s *Service) DeleteUser(ctx context.Context, id int) error {
	ctx, span := s.tracer.Start(ctx, "Service:DeleteUser()")
	defer span.End()

	if err := s.repo.DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}

	return nil
}
