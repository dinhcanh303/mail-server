package history

import (
	"context"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
	"github.com/dinhcanh303/mail-server/pkg/redis"
	"github.com/google/wire"
	"github.com/pkg/errors"
)

type service struct {
	redis redis.RedisEngine
	repo  HistoryRepo
}

var _ UseCase = (*service)(nil)

var UseCaseSet = wire.NewSet(NewUseCase)

func NewUseCase(
	redis redis.RedisEngine,
	repo HistoryRepo,
) UseCase {
	return &service{
		redis: redis,
		repo:  repo,
	}
}

// CreateHistory implements UseCase.
func (s *service) CreateHistory(ctx context.Context, history *domain.History) (*domain.History, error) {
	history, err := s.repo.CreateHistory(ctx, history)
	if err != nil {
		return nil, errors.Wrap(err, "service.CreateHistory failed")
	}
	return history, nil
}

// DeleteHistory implements UseCase.
func (s *service) DeleteHistory(ctx context.Context, id int64) error {
	err := s.repo.DeleteHistory(ctx, id)
	if err != nil {
		return errors.Wrap(err, "service.DeleteHistory failed")
	}
	return nil
}

// GetHistories implements UseCase.
func (s *service) GetHistories(ctx context.Context, limit int32, offset int32) ([]*domain.History, error) {
	histories, err := s.repo.GetHistories(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetHistories failed")
	}
	return histories, nil
}

// GetHistory implements UseCase.
func (s *service) GetHistory(ctx context.Context, id int64) (*domain.History, error) {
	history, err := s.repo.GetHistory(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetHistory failed")
	}
	return history, nil
}

// UpdateHistory implements UseCase.
func (s *service) UpdateHistory(ctx context.Context, history *domain.History) (*domain.History, error) {
	history, err := s.repo.UpdateHistory(ctx, history)
	if err != nil {
		return nil, errors.Wrap(err, "service.UpdateHistory failed")
	}
	return history, nil
}
