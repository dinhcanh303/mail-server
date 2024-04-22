package handlers

import (
	"bytes"
	"context"
	"log/slog"
	"strings"
	"text/template"

	"github.com/dinhcanh303/mail-server/internal/mail/domain"
	"github.com/dinhcanh303/mail-server/internal/mail/events"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/client"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/history"
	"github.com/dinhcanh303/mail-server/internal/pkg/event"
	"github.com/dinhcanh303/mail-server/pkg/mail"
	"github.com/google/wire"
	"github.com/pkg/errors"
)

type mailEventHandler struct {
	historyService history.UseCase
	sendMail       mail.EmailSender
	serviceClient  client.UseCase
}

var _ events.MailEventHandler = (*mailEventHandler)(nil)

var MailEventHandlerSet = wire.NewSet(NewMailEventHandler)

func NewMailEventHandler(historyService history.UseCase, sendMail mail.EmailSender, serviceClient client.UseCase) events.MailEventHandler {
	return &mailEventHandler{
		historyService: historyService,
		sendMail:       sendMail,
		serviceClient:  serviceClient,
	}
}

// Handle implements events.MailEventHandler.
func (s *mailEventHandler) Handle(ctx context.Context, event *event.SendMailEvent) error {
	slog.Info("Handle")
	client, err := s.serviceClient.GetClientEx(ctx, event.History.ApiKey)
	if err != nil {
		s.historyService.UpdateHistory(ctx, &domain.History{
			ID:     event.History.ID,
			Status: "failed",
		})
		return errors.Wrap(err, "failed to get client")
	}
	content, err := generateHTML(event.History.Content, client.Template.Html)
	if err != nil {
		return errors.Wrap(err, "failed to generate html")
	}
	to := strings.Split(event.History.To, ",")
	cc := strings.Split(event.History.Cc, ",")
	bcc := strings.Split(event.History.Bcc, ",")
	s.sendMail.Configure(
		mail.Username(client.Server.UserName),
		mail.Password(client.Server.Password),
		mail.Host(client.Server.Host),
		mail.Port(string(client.Server.Port)),
		mail.FromName(client.Server.FromName),
		mail.FromAddress(client.Server.FromAddress),
		mail.AuthProtocol(client.Server.AuthProtocol),
		mail.IDLETimeout(client.Server.IdleTimeout),
		mail.Retries(client.Server.Retries),
		mail.MaxConnections(client.Server.MaxConnections),
		mail.WaitTimeout(client.Server.WaitTimeout),
		mail.TypeTLS(string(client.Server.TLSType)),
	)

	err = s.sendMail.SendEmail(event.History.Subject, content, to, cc, bcc, []string{})
	if err != nil {
		s.historyService.UpdateHistory(ctx, &domain.History{
			ID:     event.History.ID,
			Status: "failed",
		})
		slog.Error("fail to send mail::", err)
		return errors.Wrap(err, "failed to send email")
	}
	s.historyService.UpdateHistory(ctx, &domain.History{
		ID:     event.History.ID,
		Status: "success",
	})
	return nil
}
func generateHTML(object map[string]interface{}, templateStr string) (string, error) {
	tmpl, err := template.New("htmlTemplate").Parse(templateStr)
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, object)
	if err != nil {
		return "", err
	}
	return tpl.String(), nil
}
