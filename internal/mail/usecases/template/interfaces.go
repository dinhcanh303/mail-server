package template

import (
	"context"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
)

type TemplateRepo interface {
	CreateTemplate(context.Context, *domain.Template) (*domain.Template, error)
	UpdateTemplate(context.Context, *domain.Template) (*domain.Template, error)
	DeleteTemplate(context.Context, *domain.Template) error
	GetTemplates(ctx context.Context, limit, offset int32) ([]*domain.Template, error)
}
type UseCase interface {
	CreateTemplate(context.Context, *domain.Template) (*domain.Template, error)
	UpdateTemplate(context.Context, *domain.Template) (*domain.Template, error)
	DeleteTemplate(context.Context, *domain.Template) error
	GetTemplates(ctx context.Context, limit, offset int32) ([]*domain.Template, error)
}
