// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package postgresql

import (
	"database/sql"
	"encoding/json"
	"time"
)

type MailClient struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	ServerID   int64     `json:"server_id"`
	TemplateID int64     `json:"template_id"`
	ApiKey     string    `json:"api_key"`
	IsDefault  bool      `json:"is_default"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type MailHistory struct {
	ID        int64           `json:"id"`
	ApiKey    string          `json:"api_key"`
	To        string          `json:"to_"`
	Subject   sql.NullString  `json:"subject"`
	Cc        sql.NullString  `json:"cc"`
	Bcc       sql.NullString  `json:"bcc"`
	Content   json.RawMessage `json:"content"`
	Status    sql.NullString  `json:"status"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type MailServer struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Host           string         `json:"host"`
	Port           int64          `json:"port"`
	AuthProtocol   sql.NullString `json:"auth_protocol"`
	Username       string         `json:"username"`
	Password       string         `json:"password"`
	FromName       sql.NullString `json:"from_name"`
	FromAddress    sql.NullString `json:"from_address"`
	TlsType        sql.NullString `json:"tls_type"`
	TlsSkipVerify  sql.NullBool   `json:"tls_skip_verify"`
	MaxConnections sql.NullInt64  `json:"max_connections"`
	IdleTimeout    sql.NullInt64  `json:"idle_timeout"`
	Retries        sql.NullInt64  `json:"retries"`
	WaitTimeout    sql.NullInt64  `json:"wait_timeout"`
	IsDefault      bool           `json:"is_default"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

type MailTemplate struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Html      sql.NullString `json:"html"`
	Status    sql.NullString `json:"status"`
	IsDefault bool           `json:"is_default"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
