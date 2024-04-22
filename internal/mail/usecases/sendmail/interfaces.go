package sendmail

import (
	"context"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
	"github.com/dinhcanh303/mail-server/pkg/rabbitmq/publisher"
)

type UseCase interface {
	TestSendMail(context.Context) error
	SendMail(ctx context.Context, history *domain.History) error
}

type MailEventPublisher interface {
	Configure(...publisher.Option)
	Publish(context.Context, []byte, string) error
}
