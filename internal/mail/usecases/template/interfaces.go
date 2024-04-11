package template

import (
	"context"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
)

type TemplateRepo interface {
	CreateTemplate(context.Context, *domain.Template) (*domain.Template, error)
	UpdateTemplate(context.Context, *domain.Template) (*domain.Template, error)
	DeleteTemplate(ctx context.Context, id int64) error
	GetTemplate(ctx context.Context, id int64) (*domain.Template, error)
	GetTemplates(ctx context.Context, limit, offset int32) ([]*domain.Template, error)
}
type UseCase interface {
	CreateTemplate(context.Context, *domain.Template) (*domain.Template, error)
	UpdateTemplate(context.Context, *domain.Template) (*domain.Template, error)
	DeleteTemplate(ctx context.Context, id int64) error
	GetTemplate(ctx context.Context, id int64) (*domain.Template, error)
	GetTemplates(ctx context.Context, limit, offset int32) ([]*domain.Template, error)
}
