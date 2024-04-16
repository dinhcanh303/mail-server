package client

import (
	"context"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
	sharedkernel "github.com/dinhcanh303/mail-server/internal/pkg/shared_kernel"
)

type (
	ClientRepo interface {
		CreateClient(context.Context, *domain.Client) (*domain.Client, error)
		UpdateClient(context.Context, *domain.Client) (*domain.Client, error)
		DeleteClient(ctx context.Context, id int64) error
		GetClient(ctx context.Context, id int64) (*domain.Client, error)
		GetClients(ctx context.Context, limit, offset int32) ([]*domain.Client, error)
	}

	UseCase interface {
		CreateClient(context.Context, *domain.Client) (*domain.Client, error)
		UpdateClient(context.Context, *domain.Client) (*domain.Client, error)
		DeleteClient(ctx context.Context, id int64) error
		GetClient(ctx context.Context, id int64) (*domain.Client, error)
		GetClientEx(ctx context.Context, id int64) (*sharedkernel.ClientExtra, error)
		GetClients(ctx context.Context, limit, offset int32) ([]*domain.Client, error)
	}
)
