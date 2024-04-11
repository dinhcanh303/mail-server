package app

import (
	"github.com/dinhcanh303/mail-server/cmd/mail/config"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/mail"
	"github.com/dinhcanh303/mail-server/pkg/postgres"
)

type App struct {
	Cfg *config.Config
	PG  postgres.DBEngine
	UC  mail.UseCase
}

func New(
	cfg *config.Config,
	pg postgres.DBEngine,
) *App {
	return &App{
		Cfg: cfg,
		UC:  uc,
		PG:  pg,
	}
}
