package infras

import (
	"context"

	"github.com/dinhcanh303/mail-server/internal/mail/usecases/sendmail"
	"github.com/dinhcanh303/mail-server/pkg/rabbitmq/publisher"
	"github.com/google/wire"
)

type mailEventPublisher struct {
	pub publisher.EventPublisher
}

var _ sendmail.MailEventPublisher = (*mailEventPublisher)(nil)

var MailEventPublisherSet = wire.NewSet(NewMailEventPublisher)

func NewMailEventPublisher(pub publisher.EventPublisher) sendmail.MailEventPublisher {
	return &mailEventPublisher{
		pub: pub,
	}
}

// Configure implements sendmail.MailEventPublisher.
func (m *mailEventPublisher) Configure(opts ...publisher.Option) {
	m.pub.Configure(opts...)
}

// Publish implements sendmail.MailEventPublisher.
func (m *mailEventPublisher) Publish(ctx context.Context, content []byte, contentType string) error {
	return m.pub.Publish(ctx, content, contentType)
}
