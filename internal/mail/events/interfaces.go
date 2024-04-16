package events

import (
	"context"

	"github.com/dinhcanh303/mail-server/internal/pkg/event"
)

type MailEventHandler interface {
	Handle(context.Context, *event.SendMailEvent) error
}
