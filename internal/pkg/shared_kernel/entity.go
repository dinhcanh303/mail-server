package sharedkernel

import (
	"time"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
)

type ClientExtra struct {
	ID        int64           `json:"id"`
	Name      string          `json:"name"`
	Server    domain.Server   `json:"server"`
	Template  domain.Template `json:"template"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
