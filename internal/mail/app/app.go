package app

import (
	v1 "github.com/dinhcanh303/mail-server/api/mail/v1"
	"github.com/dinhcanh303/mail-server/cmd/mail/config"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/client"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/server"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/template"
	"github.com/dinhcanh303/mail-server/pkg/postgres"
)

type App struct {
	Cfg            *config.Config
	PG             postgres.DBEngine
	UCS            server.UseCase
	UCT            template.UseCase
	UCC            client.UseCase
	MailGRPCServer v1.MailServiceServer
}

func New(
	cfg *config.Config,
	pg postgres.DBEngine,
	ucs server.UseCase,
	uct template.UseCase,
	ucc client.UseCase,
	mailGRPCServer v1.MailServiceServer,
) *App {
	return &App{
		Cfg:            cfg,
		UCS:            ucs,
		UCT:            uct,
		UCC:            ucc,
		PG:             pg,
		MailGRPCServer: mailGRPCServer,
	}
}
