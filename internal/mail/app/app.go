package app

import (
	"context"
	"encoding/json"

	v1 "github.com/dinhcanh303/mail-server/api/mail/v1"
	"github.com/dinhcanh303/mail-server/cmd/mail/config"
	"github.com/dinhcanh303/mail-server/internal/mail/events"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/client"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/sendmail"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/server"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/template"
	"github.com/dinhcanh303/mail-server/internal/pkg/event"
	"github.com/dinhcanh303/mail-server/pkg/postgres"
	"github.com/dinhcanh303/mail-server/pkg/rabbitmq/consumer"
	"github.com/dinhcanh303/mail-server/pkg/rabbitmq/publisher"
	"github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type App struct {
	Cfg              *config.Config
	PG               postgres.DBEngine
	UCS              server.UseCase
	UCT              template.UseCase
	UCC              client.UseCase
	UCSM             sendmail.UseCase
	MailGRPCServer   v1.MailServiceServer
	mailEventHandler events.MailEventHandler
	Publisher        publisher.EventPublisher
	Consumer         consumer.EventConsumer
}

func New(
	cfg *config.Config,
	pg postgres.DBEngine,
	ucs server.UseCase,
	uct template.UseCase,
	ucc client.UseCase,
	ucsm sendmail.UseCase,
	mailGRPCServer v1.MailServiceServer,
	mailEventHandler events.MailEventHandler,
	publisher publisher.EventPublisher,
	consumer consumer.EventConsumer,
) *App {
	return &App{
		Cfg:              cfg,
		UCS:              ucs,
		UCT:              uct,
		UCC:              ucc,
		UCSM:             ucsm,
		PG:               pg,
		MailGRPCServer:   mailGRPCServer,
		mailEventHandler: mailEventHandler,
		Publisher:        publisher,
		Consumer:         consumer,
	}
}

func (a *App) Worker(ctx context.Context, messages <-chan amqp091.Delivery) {
	for delivery := range messages {
		slog.Info("processDeliveries", "delivery_tag", delivery.DeliveryTag)
		slog.Info("received", "delivery_type", delivery.Type)
		switch delivery.Type {
		case "sendmail":
			var payload event.SendMailEvent
			err := json.Unmarshal(delivery.Body, &payload)
			if err != nil {
				slog.Error("failed to unmarshal message", err)
			}
			err = a.mailEventHandler.Handle(ctx, &payload)
			if err != nil {
				if err = delivery.Reject(false); err != nil {
					slog.Error("failed to reject delivery", err)
				}
				slog.Error("failed to process delivery", err)
			} else {
				err = delivery.Ack(false)
				if err != nil {
					slog.Error("failed to acknowledge delivery", err)
				}
			}
		}
	}
}
