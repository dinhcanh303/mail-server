package mail

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	_retries        = 5
	_maxConnections = 10
	_idleTimeout    = 15
	_waitTimeout    = 5
	_tlsType        = "TLS"
	_skipTls        = false
	_fromName       = "test"
	_fromAddress    = "noreply@test.com"
	_host           = "google.com"
	_port           = "465"
	_authProtocol   = "plain"
)

type emailSender struct {
	username, password, host, port, authProtocol, fromName, fromAddress, tlsType string
	maxConnections, idleTimeout, waitTimeout, retries                            int
	skipTls                                                                      bool
}

func (c *emailSender) Configure(opts ...Option) EmailSender {
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// SendEmail implements EmailSender.
func (sender *emailSender) SendEmail(subject string, content string, to []string, cc []string, bcc []string, attachFiles []string) error {
	emailFromName := sender.fromName
	emailFromAddress := sender.fromAddress
	username := sender.username
	password := sender.password
	emailHost := sender.host
	emailPort := sender.port
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s<%s>", emailFromName, emailFromAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc
	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attach file %s: %w", f, err)
		}
	}
	var smtpAuth smtp.Auth
	switch sender.authProtocol {
	case "plain":
		smtpAuth = smtp.PlainAuth("", emailFromName, password, emailHost)
	case "cram":
		smtpAuth = smtp.CRAMMD5Auth(username, password)
	case "", "none":
	default:
		return errors.New("auth protocol not supported")
	}
	smtpServerAddress := fmt.Sprintf("%s:%s", emailHost, emailPort)
	var tlsConfig *tls.Config
	if sender.tlsType != "none" {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: sender.skipTls,
			ServerName:         emailHost,
		}
	}
	if !sender.skipTls {
		return e.SendWithStartTLS(smtpServerAddress, smtpAuth, tlsConfig)
	}
	return e.Send(smtpServerAddress, smtpAuth)
}

var _ EmailSender = (*emailSender)(nil)

func NewEmailSender() EmailSender {
	return &emailSender{
		fromName:       _fromName,
		fromAddress:    _fromName,
		authProtocol:   _authProtocol,
		username:       "",
		password:       "",
		host:           _host,
		port:           _port,
		tlsType:        _tlsType,
		maxConnections: _maxConnections,
		idleTimeout:    _idleTimeout,
		waitTimeout:    _waitTimeout,
		retries:        _retries,
		skipTls:        _skipTls,
	}
}
