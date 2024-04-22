package repo

import (
	"context"
	"database/sql"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
	"github.com/dinhcanh303/mail-server/internal/mail/infras/postgresql"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/client"
	"github.com/dinhcanh303/mail-server/pkg/postgres"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/samber/lo"
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
func (c *clientRepo) CreateClient(ctx context.Context, client *domain.Client) (*domain.Client, error) {
	db := c.pg.GetDB()
	querier := postgresql.New(db)
	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.CreateClient db failed")
	}
	qtx := querier.WithTx(tx)
	result, err := qtx.CreateClient(ctx, postgresql.CreateClientParams{
		Name:       client.Name,
		ServerID:   client.ServerID,
		TemplateID: client.TemplateID,
		ApiKey:     client.ApiKey,
	})
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.CreateServer failed")
	}
	return repoToDOmainClient(result), tx.Commit()
}

// DeleteClient implements client.ClientRepo.
func (c *clientRepo) DeleteClient(ctx context.Context, id int64) error {
	db := c.pg.GetDB()
	querier := postgresql.New(db)
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "serverRepo.CreateServer db failed")
	}
	qtx := querier.WithTx(tx)
	err = qtx.DeleteClient(ctx, id)
	if err != nil {
		return errors.Wrap(err, "serverRepo.DeleteClient failed")
	}
	return tx.Commit()
}

// GetClientByApiKey implements client.ClientRepo.
func (c *clientRepo) GetClientByApiKey(ctx context.Context, apiKey string) (*domain.Client, error) {
	db := c.pg.GetDB()
	querier := postgresql.New(db)
	result, err := querier.GetClientByApiKey(ctx, apiKey)
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.GetClientByApiKey failed")
	}
	return repoToDOmainClient(result), nil
}

// GetClient implements client.ClientRepo.
func (c *clientRepo) GetClient(ctx context.Context, id int64) (*domain.Client, error) {
	db := c.pg.GetDB()
	querier := postgresql.New(db)
	result, err := querier.GetClient(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.GetClient failed")
	}
	return repoToDOmainClient(result), nil
}

// GetClients implements client.ClientRepo.
func (c *clientRepo) GetClients(ctx context.Context, limit int32, offset int32) ([]*domain.Client, error) {
	db := c.pg.GetDB()
	querier := postgresql.New(db)
	results, err := querier.GetClients(ctx, postgresql.GetClientsParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.GetClients failed")
	}
	return lo.Map(results, func(item postgresql.MailClient, _ int) *domain.Client {
		return repoToDOmainClient(item)

	}), nil
}

// UpdateClient implements client.ClientRepo.
func (c *clientRepo) UpdateClient(ctx context.Context, client *domain.Client) (*domain.Client, error) {
	db := c.pg.GetDB()
	querier := postgresql.New(db)
	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.UpdateClient db failed")
	}
	qtx := querier.WithTx(tx)
	result, err := qtx.UpdateClient(ctx, postgresql.UpdateClientParams{
		ID: client.ID,
		Name: sql.NullString{
			String: client.Name,
			Valid:  client.Name != "",
		},
		ServerID: sql.NullInt64{
			Int64: client.ServerID,
			Valid: client.ServerID != 0,
		},
		TemplateID: sql.NullInt64{
			Int64: client.TemplateID,
			Valid: client.TemplateID != 0,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.UpdateClient failed")
	}
	return repoToDOmainClient(result), tx.Commit()
}

func repoToDOmainClient(entity postgresql.MailClient) *domain.Client {
	return &domain.Client{
		ID:         entity.ID,
		Name:       entity.Name,
		ServerID:   entity.ServerID,
		TemplateID: entity.TemplateID,
		IsDefault:  entity.IsDefault,
		ApiKey:     entity.ApiKey,
		CreatedAt:  entity.CreatedAt,
		UpdatedAt:  entity.UpdatedAt,
	}
}
