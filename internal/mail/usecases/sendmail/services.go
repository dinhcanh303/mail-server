package sendmail

import (
	"context"
	"encoding/json"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
	"github.com/dinhcanh303/mail-server/internal/pkg/event"
	"github.com/dinhcanh303/mail-server/pkg/redis"
	"github.com/google/wire"
	"github.com/pkg/errors"
)

type service struct {
	redis redis.RedisEngine
	// serviceHistory history.UseCase
	pub MailEventPublisher
}

var _ UseCase = (*service)(nil)

var UseCaseSet = wire.NewSet(NewUseCase)

func NewUseCase(
	redis redis.RedisEngine,
	// serviceHistory history.UseCase,
	pub MailEventPublisher,
) UseCase {
	return &service{
		redis: redis,
		// serviceHistory: serviceHistory,
		pub: pub,
	}
}

// HandleSendMail implements UseCase.

func (s *service) SendMail(ctx context.Context, clientId int64, history *domain.History) error {
	// history, err := s.serviceHistory.CreateHistory(ctx, history)
	// if err != nil {
	// return errors.Wrap(err, "service.SendMail failed")
	// }
	event := &event.SendMailEvent{
		ClientId: clientId,
		History:  history,
	}
	eventBytes, err := json.Marshal(event)
	if err != nil {
		return errors.Wrap(err, "json.Marshal[event]")
	}
	s.pub.Publish(ctx, eventBytes, "text/plain")
	return nil
}

// TestSendMail implements UseCase.
func (s *service) TestSendMail(context.Context) error {
	panic("unimplemented")
}
