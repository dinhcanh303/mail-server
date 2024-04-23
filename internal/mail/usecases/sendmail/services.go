package sendmail

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/history"
	"github.com/dinhcanh303/mail-server/internal/pkg/event"
	"github.com/dinhcanh303/mail-server/pkg/mail"
	"github.com/dinhcanh303/mail-server/pkg/redis"
	"github.com/google/wire"
	"github.com/pkg/errors"
)

type service struct {
	redis          redis.RedisEngine
	serviceHistory history.UseCase
	pub            MailEventPublisher
	sendMail       mail.EmailSender
}

var _ UseCase = (*service)(nil)

var UseCaseSet = wire.NewSet(NewUseCase)

func NewUseCase(
	redis redis.RedisEngine,
	serviceHistory history.UseCase,
	pub MailEventPublisher,
	sendMail mail.EmailSender,
) UseCase {
	return &service{
		redis:          redis,
		serviceHistory: serviceHistory,
		pub:            pub,
		sendMail:       sendMail,
	}
}

// HandleSendMail implements UseCase.

func (s *service) SendMail(ctx context.Context, history *domain.History) error {
	slog.Info("History", history)
	history, err := s.serviceHistory.CreateHistory(ctx, history)
	if err != nil {
		return errors.Wrap(err, "service.SendMail failed")
	}
	event := &event.SendMailEvent{
		History: history,
	}
	eventBytes, err := json.Marshal(event)
	if err != nil {
		return errors.Wrap(err, "json.Marshal[event]")
	}
	s.pub.Publish(ctx, eventBytes, "text/plain")
	return nil
}

// TestSendMail implements UseCase.
func (s *service) TestSendMail(ctx context.Context, host string, port int64, authProtocol, username, password, tlsType, fromName, fromAddress string,
	idleTimeout, maxConnections, retries, waitTimeout int64, to string) error {
	s.sendMail.Configure(
		mail.Username(username),
		mail.Password(password),
		mail.Host(host),
		mail.Port(string(port)),
		mail.FromName(fromName),
		mail.FromAddress(fromAddress),
		mail.AuthProtocol(authProtocol),
		mail.IDLETimeout(idleTimeout),
		mail.Retries(retries),
		mail.MaxConnections(maxConnections),
		mail.WaitTimeout(waitTimeout),
		mail.TypeTLS(tlsType),
	)
	err := s.sendMail.SendEmail("Mail Test", `<h1>Hello world</h1>
	<p>This is a test message from <a href="https://github.com/dinhcanh303">Foden Ngo</a></p>`, []string{to}, []string{}, []string{}, []string{})
	if err != nil {
		return errors.Wrap(err, "service.TestSendMail failed")
	}
	return nil
}
