package mail

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	_retries        = 5
	_maxConnections = 10
	_idleTimeout    = 15
	_waitTimeout    = 5
	_skipTls        = false
	_fromName       = "test@test.com"
)

type emailSender struct {
	fromName, fromAddress, username, password, host, port string
	maxConnections, idleTimeout, waitTimeout, retries     int
	skipTls                                               bool
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
	emailPassword := sender.password
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
	smtpAuth := smtp.PlainAuth("", emailFromAddress, emailPassword, emailHost)
	smtpServerAddress := fmt.Sprintf("%s:%s", emailHost, emailPort)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         emailHost,
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
		username:       "",
		password:       "",
		host:           "",
		port:           "",
		maxConnections: _maxConnections,
		idleTimeout:    _idleTimeout,
		waitTimeout:    _waitTimeout,
		retries:        _retries,
		skipTls:        _skipTls,
	}
}
