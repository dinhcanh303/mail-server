package repo

import (
	"context"
	"database/sql"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
	"github.com/dinhcanh303/mail-server/internal/mail/infras/postgresql"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/history"
	"github.com/dinhcanh303/mail-server/pkg/postgres"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type historyRepo struct {
	pg postgres.DBEngine
}

var _ history.HistoryRepo = (*historyRepo)(nil)

var HistoryRepoSet = wire.NewSet(NewHistoryRepo)

func NewHistoryRepo(pg postgres.DBEngine) history.HistoryRepo {
	return &historyRepo{
		pg: pg,
	}
}

// CreateHistory implements history.HistoryRepo.
func (h *historyRepo) CreateHistory(ctx context.Context, history *domain.History) (*domain.History, error) {
	db := h.pg.GetDB()
	querier := postgresql.New(db)
	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.CreateHistory db failed")
	}
	qtx := querier.WithTx(tx)
	result, err := qtx.CreateHistory(ctx, postgresql.CreateHistoryParams{
		From: history.From,
		To:   history.To,
		Subject: sql.NullString{
			String: history.Subject,
			Valid:  history.Subject != "",
		},
		Cc: sql.NullString{
			String: history.Cc,
			Valid:  history.Cc != "",
		},
		Bcc: sql.NullString{
			String: history.Bcc,
			Valid:  history.Bcc != "",
		},
		Status: sql.NullString{
			String: history.Status,
			Valid:  history.Status != "",
		},
		// Content:,
	})
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.CreateHistory failed")
	}
	return repoHistoryToDomainHistory(result), tx.Commit()
}

// DeleteHistory implements history.HistoryRepo.
func (h *historyRepo) DeleteHistory(ctx context.Context, id int64) error {
	db := h.pg.GetDB()
	querier := postgresql.New(db)
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "serverRepo.DeleteHistory db failed")
	}
	qtx := querier.WithTx(tx)
	err = qtx.DeleteHistory(ctx, id)
	if err != nil {
		return errors.Wrap(err, "serverRepo.DeleteHistory failed")
	}
	return tx.Commit()
}

// GetHistories implements history.HistoryRepo.
func (h *historyRepo) GetHistories(ctx context.Context, limit int32, offset int32) ([]*domain.History, error) {
	db := h.pg.GetDB()
	querier := postgresql.New(db)
	results, err := querier.GetHistories(ctx, postgresql.GetHistoriesParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.GetHistory failed")
	}
	return lo.Map(results, func(item postgresql.MailHistory, _ int) *domain.History {
		return repoHistoryToDomainHistory(item)
	}), nil
}

// GetHistory implements history.HistoryRepo.
func (h *historyRepo) GetHistory(ctx context.Context, id int64) (*domain.History, error) {
	db := h.pg.GetDB()
	querier := postgresql.New(db)
	result, err := querier.GetHistory(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.GetHistory failed")
	}
	return repoHistoryToDomainHistory(result), nil
}

// UpdateHistory implements history.HistoryRepo.
func (h *historyRepo) UpdateHistory(ctx context.Context, history *domain.History) (*domain.History, error) {
	db := h.pg.GetDB()
	querier := postgresql.New(db)
	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.UpdateHistory db failed")
	}
	qtx := querier.WithTx(tx)
	result, err := qtx.UpdateHistory(ctx, postgresql.UpdateHistoryParams{
		ID: history.ID,
		From: sql.NullString{
			String: history.From,
			Valid:  history.From != "",
		},
		To: sql.NullString{
			String: history.To,
			Valid:  history.To != "",
		},
		Subject: sql.NullString{
			String: history.Subject,
			Valid:  history.Subject != "",
		},
		Cc: sql.NullString{
			String: history.Cc,
			Valid:  history.Cc != "",
		},
		Bcc: sql.NullString{
			String: history.Bcc,
			Valid:  history.Bcc != "",
		},
		Status: sql.NullString{
			String: history.Status,
			Valid:  history.Status != "",
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.UpdateHistory failed")
	}
	return repoHistoryToDomainHistory(result), tx.Commit()
}

func repoHistoryToDomainHistory(history postgresql.MailHistory) *domain.History {
	return &domain.History{
		ID:        history.ID,
		From:      history.From,
		To:        history.To,
		Subject:   history.Subject.String,
		Cc:        history.Cc.String,
		Bcc:       history.Bcc.String,
		Status:    history.Status.String,
		CreatedAt: history.CreatedAt,
		UpdatedAt: history.UpdatedAt,
	}
}
