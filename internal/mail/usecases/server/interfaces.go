package server

import (
	"context"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
)

type ServerRepo interface {
	CreateServer(context.Context, *domain.Server) (*domain.Server, error)
	UpdateServer(context.Context, *domain.Server) (*domain.Server, error)
	DeleteServer(context.Context, *domain.Server) error
	GetServers(ctx context.Context, limit, offset int32) ([]*domain.Server, error)
}
type UseCase interface {
	CreateServer(context.Context, *domain.Server) (*domain.Server, error)
	UpdateServer(context.Context, *domain.Server) (*domain.Server, error)
	DeleteServer(context.Context, *domain.Server) error
	GetServers(ctx context.Context, limit, offset int32) ([]*domain.Server, error)
}
