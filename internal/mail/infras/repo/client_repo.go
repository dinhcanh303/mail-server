package repo

import (
	"context"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/client"
	"github.com/dinhcanh303/mail-server/pkg/postgres"
	"github.com/google/wire"
)

type clientRepo struct {
	pg postgres.DBEngine
}

var _ client.ClientRepo = (*clientRepo)(nil)

func NewClientRepo(pg postgres.DBEngine) client.ClientRepo {
	return &clientRepo{pg: pg}
}

var ClientRepoSet = wire.NewSet(NewClientRepo)

// CreateClient implements client.ClientRepo.
func (c *clientRepo) CreateClient(context.Context, *domain.Client) (*domain.Client, error) {
	panic("unimplemented")
}

// DeleteClient implements client.ClientRepo.
func (c *clientRepo) DeleteClient(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// GetClient implements client.ClientRepo.
func (c *clientRepo) GetClient(ctx context.Context, id int64) (*domain.Client, error) {
	panic("unimplemented")
}

// GetClients implements client.ClientRepo.
func (c *clientRepo) GetClients(ctx context.Context, limit int32, offset int32) ([]*domain.Client, error) {
	panic("unimplemented")
}

// UpdateClient implements client.ClientRepo.
func (c *clientRepo) UpdateClient(context.Context, *domain.Client) (*domain.Client, error) {
	panic("unimplemented")
}
