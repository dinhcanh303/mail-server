package client

import (
	"context"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
	"github.com/dinhcanh303/mail-server/pkg/redis"
	"github.com/google/wire"
)

type service struct {
	redis redis.RedisEngine
	repo  ClientRepo
}

var _ UseCase = (*service)(nil)

func NewUseCase(
	redis redis.RedisEngine,
	repo ClientRepo,
) UseCase {
	return &service{
		redis: redis,
		repo:  repo,
	}
}

var UseCaseSet = wire.NewSet(NewUseCase)

// CreateClient implements UseCase.
func (s *service) CreateClient(context.Context, *domain.Client) (*domain.Client, error) {
	panic("unimplemented")
}

// DeleteClient implements UseCase.
func (s *service) DeleteClient(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// GetClient implements UseCase.
func (s *service) GetClient(ctx context.Context, id int64) (*domain.Client, error) {
	panic("unimplemented")
}

// GetClients implements UseCase.
func (s *service) GetClients(ctx context.Context, limit int32, offset int32) ([]*domain.Client, error) {
	panic("unimplemented")
}

// UpdateClient implements UseCase.
func (s *service) UpdateClient(context.Context, *domain.Client) (*domain.Client, error) {
	panic("unimplemented")
}
