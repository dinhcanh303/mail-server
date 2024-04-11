package domain

import "time"

type TLSType string

const (
	TLSOff   TLSType = "off"
	StartTLC TLSType = "start_tls"
	SSL_TLS  TLSType = "ssl/tls"
)

type Server struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Host           string    `json:"host"`
	Port           int64     `json:"port"`
	UserName       string    `json:"username"`
	Password       string    `json:"password"`
	TLS            TLSType   `json:"tls"`
	SkipTLSVerify  bool      `json:"skip_tls"`
	MaxConnections int64     `json:"max_connections"`
	Retries        int64     `json:"retries"`
	IdleTimeout    int64     `json:"idle_timeout"`
	WaitTimeout    int64     `json:"wait_timeout"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
