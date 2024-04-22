package mail

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log/slog"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	_retries        = 5
	_maxConnections = 10
	_idleTimeout    = 15
	_waitTimeout    = 5
	_tlsType        = "STARTTLS"
	_fromName       = "tlcmodular"
	_fromAddress    = "canhngo@tlcmodular.com"
	_host           = "smtp.office365.com"
	_port           = "25"
	_authProtocol   = "plain"
	_password       = "SwordfishStanley!#$@"
)

type emailSender struct {
	username, password, host, port, authProtocol, fromName, fromAddress, tlsType string
	maxConnections, idleTimeout, waitTimeout, retries                            int64
}

func (c *emailSender) Configure(opts ...Option) EmailSender {
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// SendEmail implements EmailSender.
func (sender *emailSender) SendEmail(subject string, content string, to []string, cc []string, bcc []string, attachFiles []string) error {
	// emailFromName := sender.fromName
	emailFromAddress := sender.fromAddress
	username := sender.username
	password := sender.password
	emailHost := sender.host
	emailPort := sender.port
	e := email.NewEmail()
	// e.From = fmt.Sprintf("%s<%s>", emailFromName, emailFromAddress)
	e.From = "info@tlcmodular.com"
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
		smtpAuth = smtp.PlainAuth("", emailFromAddress, password, emailHost)
	case "cram":
		smtpAuth = smtp.CRAMMD5Auth(username, password)
	case "login":
		smtpAuth = LoginAuth(emailFromAddress, password)
	case "", "none":
	default:
		return errors.New("auth protocol not supported")
	}
	smtpServerAddress := fmt.Sprintf("%s:%s", emailHost, emailPort)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         emailHost,
	}
	var err error
	var _retries = sender.retries
	for _retries > 0 {
		switch sender.tlsType {
		case "STARTTLS":
			slog.Info("STARTTLS send mail")
			err = e.SendWithStartTLS(smtpServerAddress, smtpAuth, tlsConfig)
		case "TLS":
			slog.Info("TLS send mail")
			err = e.SendWithTLS(smtpServerAddress, smtpAuth, tlsConfig)
		default:
			slog.Info("Default send mail")
			err = e.Send(smtpServerAddress, smtpAuth)
		}
		if err == nil {
			break
		}
		_retries--
	}
	if err != nil {
		return err
	}
	return nil

}

var _ EmailSender = (*emailSender)(nil)

func NewEmailSender() EmailSender {
	return &emailSender{
		fromName:       _fromAddress,
		fromAddress:    _fromAddress,
		authProtocol:   _authProtocol,
		username:       _fromAddress,
		password:       _password,
		host:           _host,
		port:           _port,
		tlsType:        _tlsType,
		maxConnections: _maxConnections,
		idleTimeout:    _idleTimeout,
		waitTimeout:    _waitTimeout,
		retries:        _retries,
	}
}
