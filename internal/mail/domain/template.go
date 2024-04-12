package domain

import "time"

type Template struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Html      string    `json:"html"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTemplate(name string, status string, html string) *Template {
	if status == "" {
		status = "active"
	}
	return &Template{
		Name:   name,
		Status: status,
		Html:   html,
	}
}
