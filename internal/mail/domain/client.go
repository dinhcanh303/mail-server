package domain

import "time"

type Client struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	ServerID   int64     `json:"server_id"`
	TemplateID int64     `json:"template_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
