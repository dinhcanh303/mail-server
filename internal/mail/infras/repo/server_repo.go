package repo

import (
	"context"
	"database/sql"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
	"github.com/dinhcanh303/mail-server/internal/mail/infras/postgresql"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/server"
	"github.com/dinhcanh303/mail-server/pkg/postgres"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type serverRepo struct {
	pg postgres.DBEngine
}

var _ server.ServerRepo = (*serverRepo)(nil)

func NewServerRepo(pg postgres.DBEngine) server.ServerRepo {
	return &serverRepo{pg: pg}
}

var ServerRepoSet = wire.NewSet(NewServerRepo)

// CreateServer implements server.ServerRepo.
func (s *serverRepo) CreateServer(ctx context.Context, server *domain.Server) (*domain.Server, error) {
	db := s.pg.GetDB()
	querier := postgresql.New(db)
	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.CreateServer db failed")
	}
	qtx := querier.WithTx(tx)
	result, err := qtx.CreateServer(ctx, postgresql.CreateServerParams{
		Name: server.Name,
		Host: server.Host,
		Port: server.Port,
		AuthProtocol: sql.NullString{
			String: server.AuthProtocol,
			Valid:  server.AuthProtocol != "",
		},
		Username: server.UserName,
		Password: server.Password,
		FromName: sql.NullString{
			String: server.FromName,
			Valid:  server.FromName != "",
		},
		FromAddress: sql.NullString{
			String: server.FromAddress,
			Valid:  server.FromAddress != "",
		},
		TlsType: sql.NullString{
			String: string(server.TLSType),
			Valid:  string(server.TLSType) != "",
		},
		TlsSkipVerify: sql.NullBool{
			Bool:  server.TLSSkipVerify,
			Valid: server.TLSSkipVerify,
		},
		MaxConnections: sql.NullInt64{
			Int64: server.MaxConnections,
			Valid: server.MaxConnections > 0,
		},
		IdleTimeout: sql.NullInt64{
			Int64: server.IdleTimeout,
			Valid: server.IdleTimeout > 0,
		},
		Retries: sql.NullInt64{
			Int64: server.Retries,
			Valid: server.Retries > 0,
		},
		WaitTimeout: sql.NullInt64{
			Int64: server.WaitTimeout,
			Valid: server.WaitTimeout > 0,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.CreateServer failed")
	}
	return repoServerToDomainServer(result), tx.Commit()
}

// DeleteServer implements server.ServerRepo.
func (s *serverRepo) DeleteServer(ctx context.Context, id int64) error {
	db := s.pg.GetDB()
	querier := postgresql.New(db)
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "serverRepo.DeleteServer db failed")
	}
	qtx := querier.WithTx(tx)
	err = qtx.DeleteServer(ctx, id)
	if err != nil {
		return errors.Wrap(err, "serverRepo.DeleteServer failed")
	}
	return tx.Commit()
}

// GetServer implements server.ServerRepo.
func (s *serverRepo) GetServer(ctx context.Context, id int64) (*domain.Server, error) {
	db := s.pg.GetDB()
	querier := postgresql.New(db)
	result, err := querier.GetServer(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.GetServer failed")
	}
	return repoServerToDomainServer(result), nil
}

// GetServers implements server.ServerRepo.
func (s *serverRepo) GetServers(ctx context.Context, limit int32, offset int32) ([]*domain.Server, error) {
	db := s.pg.GetDB()
	querier := postgresql.New(db)
	results, err := querier.GetServers(ctx, postgresql.GetServersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.GetServers failed")
	}
	return lo.Map(results, func(item postgresql.MailServer, _ int) *domain.Server {
		return repoServerToDomainServer(item)
	}), nil
}

// UpdateServer implements server.ServerRepo.
func (s *serverRepo) UpdateServer(ctx context.Context, server *domain.Server) (*domain.Server, error) {
	db := s.pg.GetDB()
	querier := postgresql.New(db)
	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.UpdateServer db failed")
	}
	qtx := querier.WithTx(tx)
	result, err := qtx.UpdateServer(ctx, postgresql.UpdateServerParams{
		ID: server.ID,
		Name: sql.NullString{
			String: server.Name,
			Valid:  server.Name != "",
		},
		Host: sql.NullString{
			String: server.Host,
			Valid:  server.Host != "",
		},
		Port: sql.NullInt64{
			Int64: server.Port,
			Valid: server.Port > 0,
		},
		AuthProtocol: sql.NullString{
			String: server.AuthProtocol,
			Valid:  server.AuthProtocol != "",
		},
		Username: sql.NullString{
			String: server.UserName,
			Valid:  server.UserName != "",
		},
		Password: sql.NullString{
			String: server.Password,
			Valid:  server.Password != "",
		},
		FromName: sql.NullString{
			String: server.FromName,
			Valid:  server.FromName != "",
		},
		FromAddress: sql.NullString{
			String: server.FromAddress,
			Valid:  server.FromAddress != "",
		},
		TlsType: sql.NullString{
			String: string(server.TLSType),
			Valid:  string(server.TLSType) == "",
		},
		TlsSkipVerify: sql.NullBool{
			Bool:  server.TLSSkipVerify,
			Valid: server.TLSSkipVerify,
		},
		MaxConnections: sql.NullInt64{
			Int64: server.MaxConnections,
			Valid: server.MaxConnections > 0,
		},
		IdleTimeout: sql.NullInt64{
			Int64: server.IdleTimeout,
			Valid: server.IdleTimeout > 0,
		},
		Retries: sql.NullInt64{
			Int64: server.Retries,
			Valid: server.Retries > 0,
		},
		WaitTimeout: sql.NullInt64{
			Int64: server.WaitTimeout,
			Valid: server.WaitTimeout > 0,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.UpdateServer failed")
	}
	return repoServerToDomainServer(result), tx.Commit()
}
func repoServerToDomainServer(result postgresql.MailServer) *domain.Server {
	return &domain.Server{
		ID:             result.ID,
		Name:           result.Name,
		Host:           result.Host,
		Port:           result.Port,
		AuthProtocol:   result.AuthProtocol.String,
		UserName:       result.Username,
		Password:       result.Password,
		FromName:       result.FromName.String,
		FromAddress:    result.FromAddress.String,
		TLSType:        domain.TLSType(result.TlsType.String),
		TLSSkipVerify:  result.TlsSkipVerify.Bool,
		MaxConnections: result.MaxConnections.Int64,
		Retries:        result.Retries.Int64,
		IdleTimeout:    result.IdleTimeout.Int64,
		WaitTimeout:    result.WaitTimeout.Int64,
		IsDefault:      result.IsDefault,
		CreatedAt:      result.CreatedAt,
		UpdatedAt:      result.UpdatedAt,
	}
}
