package template

import (
	"context"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
	"github.com/dinhcanh303/mail-server/pkg/redis"
	"github.com/google/wire"
	"github.com/pkg/errors"
)

type service struct {
	redis redis.RedisEngine
	repo  TemplateRepo
}

var _ UseCase = (*service)(nil)

func NewUseCase(
	redis redis.RedisEngine,
	repo TemplateRepo,
) UseCase {
	return &service{
		redis: redis,
		repo:  repo,
	}
}

var UseCaseSet = wire.NewSet(NewUseCase)

// CreateTemplate implements UseCase.
func (s *service) CreateTemplate(ctx context.Context, template *domain.Template) (*domain.Template, error) {
	template, err := s.repo.CreateTemplate(ctx, template)
	if err != nil {
		return nil, errors.Wrap(err, "service.CreateTemplate failed")
	}
	return template, nil
}

// Deletetemplate implements UseCase.
func (s *service) DeleteTemplate(ctx context.Context, id int64) error {
	err := s.repo.DeleteTemplate(ctx, id)
	if err != nil {
		return errors.Wrap(err, "service.CreateTemplate failed")
	}
	return nil
}

// GetTemplate implements UseCase.
func (s *service) GetTemplate(ctx context.Context, id int64) (*domain.Template, error) {
	template, err := s.repo.GetTemplate(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "service.Gettemplate failed")
	}
	return template, nil
}

// GetTemplates implements UseCase.
func (s *service) GetTemplates(ctx context.Context, limit int32, offset int32) ([]*domain.Template, error) {
	templates, err := s.repo.GetTemplates(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetTemplates failed")
	}
	return templates, nil
}

// GetTemplatesActive implements UseCase.
func (s *service) GetTemplatesActive(ctx context.Context) ([]*domain.Template, error) {
	templates, err := s.repo.GetTemplatesActive(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetTemplatesActive failed")
	}
	return templates, nil
}

// UpdateTemplate implements UseCase.
func (s *service) UpdateTemplate(ctx context.Context, template *domain.Template) (*domain.Template, error) {
	template, err := s.repo.UpdateTemplate(ctx, template)
	if err != nil {
		return nil, errors.Wrap(err, "service.UpdateTemplate failed")
	}
	return template, nil
}
