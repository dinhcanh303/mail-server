package domain

type TLSType string

const (
	TLSOff   TLSType = "off"
	StartTLC TLSType = "start_tls"
	SSL_TLS  TLSType = "ssl/tls"
)

type Server struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Host           string  `json:"host"`
	Port           int     `json:"port"`
	UserName       string  `json:"username"`
	Password       string  `json:"password"`
	TLS            TLSType `json:"tls"`
	SkipTLSVerify  bool    `json:"skip_tls"`
	MaxConnections int     `json:"max_connections"`
	Retries        int     `json:"retries"`
	IdleTimeout    int     `json:"idle_timeout"`
	WaitTimeout    int     `json:"wait_timeout"`
}
