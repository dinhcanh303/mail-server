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
	TLSType        TLSType   `json:"tls_type"`
	TLSSkipVerify  bool      `json:"skip_tls_verify"`
	MaxConnections int64     `json:"max_connections"`
	Retries        int64     `json:"retries"`
	IdleTimeout    int64     `json:"idle_timeout"`
	WaitTimeout    int64     `json:"wait_timeout"`
	IsDefault      bool      `json:"is_default"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func NewServer(name string, host string, port int64, username string, password string, tlsType string, tlsSkipVerify bool,
	maxConnections int64, retries int64, idleTimeout int64, waitTimeout int64,
) *Server {
	if tlsType == "" {
		tlsType = "TLS"
	}
	if maxConnections == 0 {
		maxConnections = 10
	}
	if retries == 0 {
		retries = 2
	}
	if idleTimeout == 0 {
		idleTimeout = 15
	}
	if waitTimeout == 0 {
		waitTimeout = 5
	}
	return &Server{
		Name:           name,
		Host:           host,
		Port:           port,
		UserName:       username,
		Password:       password,
		TLSType:        TLSType(tlsType),
		TLSSkipVerify:  tlsSkipVerify,
		MaxConnections: maxConnections,
		Retries:        retries,
		IdleTimeout:    idleTimeout,
		WaitTimeout:    waitTimeout,
	}
}
