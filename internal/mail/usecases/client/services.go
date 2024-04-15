package client

import (
	"context"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
	"github.com/dinhcanh303/mail-server/pkg/redis"
	"github.com/google/wire"
	"github.com/pkg/errors"
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
func (s *service) CreateClient(ctx context.Context, client *domain.Client) (*domain.Client, error) {
	client, err := s.repo.CreateClient(ctx, client)
	if err != nil {
		return nil, errors.Wrap(err, "service.CreateClient failed")
	}
	return client, nil
}

// DeleteClient implements UseCase.
func (s *service) DeleteClient(ctx context.Context, id int64) error {
	err := s.repo.DeleteClient(ctx, id)
	if err != nil {
		return errors.Wrap(err, "service.DeleteClient failed")
	}
	return nil
}

// GetClient implements UseCase.
func (s *service) GetClient(ctx context.Context, id int64) (*domain.Client, error) {
	client, err := s.repo.GetClient(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetClient failed")
	}
	return client, nil
}

// GetClients implements UseCase.
func (s *service) GetClients(ctx context.Context, limit int32, offset int32) ([]*domain.Client, error) {
	clients, err := s.repo.GetClients(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetClients failed")
	}
	return clients, nil
}

// UpdateClient implements UseCase.
func (s *service) UpdateClient(ctx context.Context, client *domain.Client) (*domain.Client, error) {
	client, err := s.repo.UpdateClient(ctx, client)
	if err != nil {
		return nil, errors.Wrap(err, "service.UpdateClient failed")
	}
	return client, nil
}
