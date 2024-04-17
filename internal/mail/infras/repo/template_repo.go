package repo

import (
	"context"
	"database/sql"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
	"github.com/dinhcanh303/mail-server/internal/mail/infras/postgresql"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/template"
	"github.com/dinhcanh303/mail-server/pkg/postgres"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

type templateRepo struct {
	pg postgres.DBEngine
}

var _ template.UseCase = (*templateRepo)(nil)

func NewTemplateRepo(pg postgres.DBEngine) template.TemplateRepo {
	return &templateRepo{pg: pg}
}

var TemplateRepoSet = wire.NewSet(NewTemplateRepo)

// CreateTemplate implements template.UseCase.
func (t *templateRepo) CreateTemplate(ctx context.Context, template *domain.Template) (*domain.Template, error) {
	db := t.pg.GetDB()
	querier := postgresql.New(db)
	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.CreateServer db failed")
	}
	qtx := querier.WithTx(tx)
	result, err := qtx.CreateTemplate(ctx, postgresql.CreateTemplateParams{
		Html: sql.NullString{
			String: template.Html,
			Valid:  template.Html != "",
		},
		Name: template.Name,
		Status: sql.NullString{
			String: template.Status,
			Valid:  template.Status != "",
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.CreateServer failed")
	}
	return repoTemplateToDomainTemplate(result), tx.Commit()
}

// DeleteTemplate implements template.UseCase.
func (t *templateRepo) DeleteTemplate(ctx context.Context, id int64) error {
	db := t.pg.GetDB()
	querier := postgresql.New(db)
	tx, err := db.Begin()
	if err != nil {
		return errors.Wrap(err, "templateRepo.DeleteTemplate db failed")
	}
	qtx := querier.WithTx(tx)
	err = qtx.DeleteTemplate(ctx, id)
	if err != nil {
		return errors.Wrap(err, "templateRepo.DeleteTemplate failed")
	}
	return tx.Commit()
}

// GetTemplate implements template.UseCase.
func (t *templateRepo) GetTemplate(ctx context.Context, id int64) (*domain.Template, error) {
	db := t.pg.GetDB()
	querier := postgresql.New(db)
	result, err := querier.GetTemplate(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "templateRepo.GetTemplate failed")
	}
	return repoTemplateToDomainTemplate(result), nil
}

// GetTemplates implements template.UseCase.
func (t *templateRepo) GetTemplates(ctx context.Context, limit int32, offset int32) ([]*domain.Template, error) {
	db := t.pg.GetDB()
	querier := postgresql.New(db)
	results, err := querier.GetTemplates(ctx, postgresql.GetTemplatesParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, errors.Wrap(err, "templateRepo.GetTemplates failed")
	}
	return lo.Map(results, func(item postgresql.MailTemplate, _ int) *domain.Template {
		return repoTemplateToDomainTemplate(item)
	}), nil
}

// GetTemplatesActive implements template.UseCase.
func (t *templateRepo) GetTemplatesActive(ctx context.Context) ([]*domain.Template, error) {
	db := t.pg.GetDB()
	querier := postgresql.New(db)
	results, err := querier.GetTemplatesActive(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "templateRepo.GetTemplatesActive failed")
	}
	return lo.Map(results, func(item postgresql.MailTemplate, _ int) *domain.Template {
		return repoTemplateToDomainTemplate(item)
	}), nil
}

// UpdateTemplate implements template.UseCase.
func (t *templateRepo) UpdateTemplate(ctx context.Context, template *domain.Template) (*domain.Template, error) {
	db := t.pg.GetDB()
	querier := postgresql.New(db)
	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.UpdateServer db failed")
	}
	qtx := querier.WithTx(tx)
	result, err := qtx.UpdateTemplate(ctx, postgresql.UpdateTemplateParams{
		ID: template.ID,
		Name: sql.NullString{
			String: template.Name,
			Valid:  template.Name != "",
		},
		Html: sql.NullString{
			String: template.Html,
			Valid:  template.Html != "",
		},
		Status: sql.NullString{
			String: template.Status,
			Valid:  template.Status != "",
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "serverRepo.UpdateServer failed")
	}
	return repoTemplateToDomainTemplate(result), tx.Commit()
}

func repoTemplateToDomainTemplate(result postgresql.MailTemplate) *domain.Template {
	return &domain.Template{
		ID:        result.ID,
		Name:      result.Name,
		Status:    result.Status.String,
		Html:      result.Html.String,
		IsDefault: result.IsDefault,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}
}
