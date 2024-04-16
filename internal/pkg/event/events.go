package event

import "github.com/dinhcanh303/mail-server/internal/mail/domain"

type SendMailEvent struct {
	ClientId int64
	History  *domain.History
}
