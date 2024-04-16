package history

import (
	"context"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
)

type HistoryRepo interface {
	CreateHistory(context.Context, *domain.History) (*domain.History, error)
	UpdateHistory(context.Context, *domain.History) (*domain.History, error)
	DeleteHistory(ctx context.Context, id int64) error
	GetHistory(ctx context.Context, id int64) (*domain.History, error)
	GetHistories(ctx context.Context, limit, offset int32) ([]*domain.History, error)
}
type UseCase interface {
	CreateHistory(context.Context, *domain.History) (*domain.History, error)
	UpdateHistory(context.Context, *domain.History) (*domain.History, error)
	DeleteHistory(ctx context.Context, id int64) error
	GetHistory(ctx context.Context, id int64) (*domain.History, error)
	GetHistories(ctx context.Context, limit, offset int32) ([]*domain.History, error)
}
