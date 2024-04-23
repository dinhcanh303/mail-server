package sendmail

import (
	"context"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
	"github.com/dinhcanh303/mail-server/pkg/rabbitmq/publisher"
)

type UseCase interface {
	TestSendMail(ctx context.Context, host string, port int64, authProtocol, username, password, tlsType, fromName, fromAddress string,
		idleTimeout, maxConnections, retries, waitTimeout int64, to string) error
	SendMail(ctx context.Context, history *domain.History) error
}

type MailEventPublisher interface {
	Configure(...publisher.Option)
	Publish(context.Context, []byte, string) error
}
