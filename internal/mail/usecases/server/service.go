package server

import (
	"context"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
	"github.com/dinhcanh303/mail-server/pkg/redis"
	"github.com/google/wire"
	"github.com/pkg/errors"
)

type service struct {
	redis redis.RedisEngine
	repo  ServerRepo
}

var _ UseCase = (*service)(nil)

func NewUseCase(
	redis redis.RedisEngine,
	repo ServerRepo,
) UseCase {
	return &service{
		redis: redis,
		repo:  repo,
	}
}

var UseCaseSet = wire.NewSet(NewUseCase)

// CreateServer implements UseCase.
func (s *service) CreateServer(ctx context.Context, server *domain.Server) (*domain.Server, error) {
	server, err := s.repo.CreateServer(ctx, server)
	if err != nil {
		return nil, errors.Wrap(err, "service.CreateServer failed")
	}
	return server, nil
}

// DeleteServer implements UseCase.
func (s *service) DeleteServer(ctx context.Context, id int64) error {
	err := s.repo.DeleteServer(ctx, id)
	if err != nil {
		return errors.Wrap(err, "service.CreateServer failed")
	}
	return nil
}

// GetServer implements UseCase.
func (s *service) GetServer(ctx context.Context, id int64) (*domain.Server, error) {
	server, err := s.repo.GetServer(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetServer failed")
	}
	return server, nil
}

// GetServers implements UseCase.
func (s *service) GetServers(ctx context.Context, limit int32, offset int32) ([]*domain.Server, error) {
	servers, err := s.repo.GetServers(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetServers failed")
	}
	return servers, nil
}

// UpdateServer implements UseCase.
func (s *service) UpdateServer(ctx context.Context, server *domain.Server) (*domain.Server, error) {
	server, err := s.repo.UpdateServer(ctx, server)
	if err != nil {
		return nil, errors.Wrap(err, "service.UpdateServer failed")
	}
	return server, nil
}
