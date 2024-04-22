package domain

import (
	"time"
)

type History struct {
	ID        int64                  `json:"id"`
	ApiKey    string                 `json:"api_key"`
	To        string                 `json:"to"`
	Subject   string                 `json:"subject"`
	Cc        string                 `json:"cc"`
	Bcc       string                 `json:"bcc"`
	Content   map[string]interface{} `json:"content"`
	Status    string                 `json:"status"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"update_at"`
}
