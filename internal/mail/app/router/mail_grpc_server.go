package router

import (
	"context"

	v1 "github.com/dinhcanh303/mail-server/api/mail/v1"
	"github.com/dinhcanh303/mail-server/cmd/mail/config"
	"github.com/dinhcanh303/mail-server/internal/mail/domain"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/client"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/history"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/sendmail"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/server"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/template"
	"github.com/dinhcanh303/mail-server/pkg/constant"
	"github.com/dinhcanh303/mail-server/pkg/utils"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type mailGRPCServer struct {
	v1.UnimplementedMailServiceServer
	ucTemp     template.UseCase
	ucServer   server.UseCase
	ucClient   client.UseCase
	ucSendMail sendmail.UseCase
	ucHistory  history.UseCase
	cfg        *config.Config
}

var _ v1.MailServiceServer = (*mailGRPCServer)(nil)

var MailGRPCServerSet = wire.NewSet(NewMailGRPCServer)

const (
	token = "success"
)

func NewMailGRPCServer(
	grpcServer *grpc.Server,
	ucTemp template.UseCase,
	ucServer server.UseCase,
	ucClient client.UseCase,
	ucSendMail sendmail.UseCase,
	ucHistory history.UseCase,
	cfg *config.Config,
) v1.MailServiceServer {
	svc := mailGRPCServer{
		ucTemp:     ucTemp,
		ucServer:   ucServer,
		ucClient:   ucClient,
		ucSendMail: ucSendMail,
		ucHistory:  ucHistory,
		cfg:        cfg,
	}
	v1.RegisterMailServiceServer(grpcServer, &svc)
	reflection.Register(grpcServer)
	return &svc
}

func (m *mailGRPCServer) Login(ctx context.Context, request *v1.LoginRequest) (*v1.LoginResponse, error) {
	cfgAuth := m.cfg.Auth
	if request.Username != cfgAuth.Username {
		return nil, errors.New("username doesn't not exits")
	}
	if request.Password != cfgAuth.Password {
		return nil, errors.New("password mismatch")
	}
	return &v1.LoginResponse{
		AccessToken: token,
	}, nil
}
func (m *mailGRPCServer) Logout(ctx context.Context, request *v1.LogoutRequest) (*v1.LogoutResponse, error) {
	return &v1.LogoutResponse{}, nil
}
func (m *mailGRPCServer) CreateServer(ctx context.Context, request *v1.CreateServerRequest) (*v1.CreateServerResponse, error) {
	model := domain.NewServer(
		request.Server.Name,
		request.Server.Host,
		request.Server.Port,
		request.Server.AuthProtocol,
		request.Server.Username,
		request.Server.Password,
		request.Server.FromName,
		request.Server.FromAddress,
		request.Server.TlsType,
		request.Server.MaxConnections,
		request.Server.Retries,
		request.Server.IdleTimeout,
		request.Server.WaitTimeout)
	result, err := m.ucServer.CreateServer(ctx, model)
	if err != nil {
		return nil, errors.Wrap(err, "ucServer.CreateServer failed")
	}
	return &v1.CreateServerResponse{
		Server: entityServerToProtobuf(result),
	}, nil
}
func (m *mailGRPCServer) DuplicateServer(ctx context.Context, request *v1.DuplicateServerRequest) (*v1.DuplicateServerResponse, error) {

	serverId := request.Server.Id
	serverName := request.Server.Name
	if serverId != 0 && serverName != "" {
		server, err := m.ucServer.GetServer(ctx, serverId)
		if err != nil {
			return nil, errors.Wrap(err, "ucClient.DuplicateServer failed , please check client id")
		}
		result, err := m.ucServer.CreateServer(ctx, &domain.Server{
			Name:           serverName,
			Host:           server.Host,
			Port:           server.Port,
			AuthProtocol:   server.AuthProtocol,
			UserName:       server.UserName,
			Password:       server.Password,
			FromName:       server.FromName,
			FromAddress:    server.FromAddress,
			TLSType:        server.TLSType,
			MaxConnections: server.MaxConnections,
			IdleTimeout:    server.IdleTimeout,
			Retries:        server.Retries,
			WaitTimeout:    server.WaitTimeout,
			IsDefault:      server.IsDefault,
		})
		if err != nil {
			return nil, errors.Wrap(err, "ucClient.DuplicateTemplate failed")
		}
		return &v1.DuplicateServerResponse{
			Server: entityServerToProtobuf(result),
		}, nil
	}
	return nil, errors.New("ucTemp.DuplicateTemplate failed,please check id and name")
}
func (m *mailGRPCServer) GetServer(ctx context.Context, request *v1.GetServerRequest) (*v1.GetServerResponse, error) {
	result, err := m.ucServer.GetServer(ctx, request.Id)
	if err != nil {
		return nil, errors.Wrap(err, "ucServer.GetServer failed")
	}
	return &v1.GetServerResponse{
		Server: entityServerToProtobuf(result),
	}, nil
}
func (m *mailGRPCServer) GetServers(ctx context.Context, request *v1.GetServersRequest) (*v1.GetServersResponse, error) {
	results, err := m.ucServer.GetServers(ctx, request.Limit, request.Offset)
	if err != nil {
		return nil, errors.Wrap(err, "ucServer.GetServers failed")
	}
	return &v1.GetServersResponse{
		Servers: lo.Map(results, func(item *domain.Server, _ int) *v1.Server {
			return entityServerToProtobuf(item)
		}),
	}, nil
}

func (m *mailGRPCServer) DeleteServer(ctx context.Context, request *v1.DeleteServerRequest) (*v1.DeleteServerResponse, error) {
	err := m.ucServer.DeleteServer(ctx, request.Id)
	if err != nil {
		return nil, errors.Wrap(err, "ucServer.DeleteServer failed")
	}
	return &v1.DeleteServerResponse{}, nil
}

func (m *mailGRPCServer) UpdateServer(ctx context.Context, request *v1.UpdateServerRequest) (*v1.UpdateServerResponse, error) {
	result, err := m.ucServer.UpdateServer(ctx, &domain.Server{
		ID:             request.Server.Id,
		Name:           request.Server.Name,
		Host:           request.Server.Host,
		Port:           request.Server.Port,
		AuthProtocol:   request.Server.AuthProtocol,
		UserName:       request.Server.Username,
		Password:       request.Server.Password,
		FromName:       request.Server.FromName,
		FromAddress:    request.Server.FromAddress,
		TLSType:        domain.TLSType(request.Server.TlsType),
		MaxConnections: request.Server.MaxConnections,
		Retries:        request.Server.Retries,
		IdleTimeout:    request.Server.IdleTimeout,
		WaitTimeout:    request.Server.WaitTimeout,
	})
	if err != nil {
		return nil, errors.Wrap(err, "ucServer.UpdateServer failed")
	}
	return &v1.UpdateServerResponse{
		Server: entityServerToProtobuf(result),
	}, nil
}
func (m *mailGRPCServer) CreateTemplate(ctx context.Context, request *v1.CreateTemplateRequest) (*v1.CreateTemplateResponse, error) {
	model := domain.NewTemplate(request.Template.Name, request.Template.Status, request.Template.Html)
	result, err := m.ucTemp.CreateTemplate(ctx, model)
	if err != nil {
		return nil, errors.Wrap(err, "ucTemp.CreateTemplate failed")
	}
	return &v1.CreateTemplateResponse{
		Template: entityTemplateToProtobuf(result),
	}, nil
}
func (m *mailGRPCServer) DuplicateTemplate(ctx context.Context, request *v1.DuplicateTemplateRequest) (*v1.DuplicateTemplateResponse, error) {

	templateId := request.Template.Id
	templateName := request.Template.Name
	if templateId != 0 && templateName != "" {
		template, err := m.ucTemp.GetTemplate(ctx, templateId)
		if err != nil {
			return nil, errors.Wrap(err, "ucClient.DuplicateTemplate failed , please check client id")
		}
		result, err := m.ucTemp.CreateTemplate(ctx, &domain.Template{
			Name:      templateName,
			Status:    template.Html,
			Html:      template.Html,
			IsDefault: template.IsDefault,
		})
		if err != nil {
			return nil, errors.Wrap(err, "ucClient.DuplicateTemplate failed")
		}
		return &v1.DuplicateTemplateResponse{
			Template: entityTemplateToProtobuf(result),
		}, nil
	}
	return nil, errors.New("ucTemp.DuplicateTemplate failed,please check id and name")
}
func (m *mailGRPCServer) GetTemplate(ctx context.Context, request *v1.GetTemplateRequest) (*v1.GetTemplateResponse, error) {
	result, err := m.ucTemp.GetTemplate(ctx, request.Id)
	if err != nil {
		return nil, errors.Wrap(err, "ucTemp.GetTemplate failed")
	}
	return &v1.GetTemplateResponse{
		Template: entityTemplateToProtobuf(result),
	}, nil
}
func (m *mailGRPCServer) GetTemplates(ctx context.Context, request *v1.GetTemplatesRequest) (*v1.GetTemplatesResponse, error) {
	results, err := m.ucTemp.GetTemplates(ctx, request.Limit, request.Offset)
	if err != nil {
		return nil, errors.Wrap(err, "ucTemp.GetTemplates failed")
	}
	return &v1.GetTemplatesResponse{
		Templates: lo.Map(results, func(item *domain.Template, _ int) *v1.Template {
			return entityTemplateToProtobuf(item)
		}),
	}, nil
}
func (m *mailGRPCServer) GetTemplatesActive(ctx context.Context, request *v1.GetTemplatesActiveRequest) (*v1.GetTemplatesActiveResponse, error) {
	results, err := m.ucTemp.GetTemplatesActive(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "ucTemp.GetTemplates failed")
	}
	return &v1.GetTemplatesActiveResponse{
		Templates: lo.Map(results, func(item *domain.Template, _ int) *v1.Template {
			return entityTemplateToProtobuf(item)
		}),
	}, nil
}
func (m *mailGRPCServer) DeleteTemplate(ctx context.Context, request *v1.DeleteTemplateRequest) (*v1.DeleteTemplateResponse, error) {
	err := m.ucTemp.DeleteTemplate(ctx, request.Id)
	if err != nil {
		return nil, errors.Wrap(err, "ucTemp.DeleteTemplate failed")
	}
	return &v1.DeleteTemplateResponse{}, nil
}
func (m *mailGRPCServer) UpdateTemplate(ctx context.Context, request *v1.UpdateTemplateRequest) (*v1.UpdateTemplateResponse, error) {
	result, err := m.ucTemp.UpdateTemplate(ctx, &domain.Template{
		ID:     request.Template.Id,
		Name:   request.Template.Name,
		Status: request.Template.Status,
		Html:   request.Template.Html,
	})
	if err != nil {
		return nil, errors.Wrap(err, "ucTemp.UpdateTemplate failed")
	}
	return &v1.UpdateTemplateResponse{
		Template: entityTemplateToProtobuf(result),
	}, nil
}

func (m *mailGRPCServer) CreateClient(ctx context.Context, request *v1.CreateClientRequest) (*v1.CreateClientResponse, error) {
	result, err := m.ucClient.CreateClient(ctx, &domain.Client{
		Name:       request.Client.Name,
		ServerID:   request.Client.ServerId,
		TemplateID: request.Client.TemplateId,
	})
	if err != nil {
		return nil, errors.Wrap(err, "ucTemp.CreateClient failed")
	}
	return &v1.CreateClientResponse{
		Client: entityClientToProtobuf(result),
	}, nil
}
func (m *mailGRPCServer) DuplicateClient(ctx context.Context, request *v1.DuplicateClientRequest) (*v1.DuplicateClientResponse, error) {

	clientId := request.Client.Id
	clientName := request.Client.Name
	if clientId != 0 && clientName != "" {
		client, err := m.ucClient.GetClient(ctx, clientId)
		if err != nil {
			return nil, errors.Wrap(err, "ucClient.DuplicateClient failed , please check client id")
		}
		result, err := m.ucClient.CreateClient(ctx, &domain.Client{
			Name:       clientName,
			ServerID:   client.ServerID,
			TemplateID: client.TemplateID,
			IsDefault:  client.IsDefault,
		})
		if err != nil {
			return nil, errors.Wrap(err, "ucClient.DuplicateClient failed")
		}
		return &v1.DuplicateClientResponse{
			Client: entityClientToProtobuf(result),
		}, nil
	}
	return nil, errors.New("ucTemp.DuplicateClient failed,please check id and name")
}
func (m *mailGRPCServer) GetClient(ctx context.Context, request *v1.GetClientRequest) (*v1.GetClientResponse, error) {
	result, err := m.ucClient.GetClient(ctx, request.Id)
	if err != nil {
		return nil, errors.Wrap(err, "ucTemp.GetClient failed")
	}
	return &v1.GetClientResponse{
		Client: entityClientToProtobuf(result),
	}, nil
}
func (m *mailGRPCServer) GetClients(ctx context.Context, request *v1.GetClientsRequest) (*v1.GetClientsResponse, error) {
	results, err := m.ucClient.GetClients(ctx, request.Limit, request.Offset)
	if err != nil {
		return nil, errors.Wrap(err, "ucTemp.GetClients failed")
	}
	return &v1.GetClientsResponse{
		Clients: lo.Map(results, func(item *domain.Client, _ int) *v1.Client {
			return entityClientToProtobuf(item)
		}),
	}, nil
}
func (m *mailGRPCServer) DeleteClient(ctx context.Context, request *v1.DeleteClientRequest) (*v1.DeleteClientResponse, error) {
	err := m.ucClient.DeleteClient(ctx, request.Id)
	if err != nil {
		return nil, errors.Wrap(err, "ucTemp.DeleteClient failed")
	}
	return &v1.DeleteClientResponse{}, nil
}
func (m *mailGRPCServer) UpdateClient(ctx context.Context, request *v1.UpdateClientRequest) (*v1.UpdateClientResponse, error) {
	result, err := m.ucClient.UpdateClient(ctx, &domain.Client{
		ID:         request.Client.Id,
		Name:       request.Client.Name,
		ServerID:   request.Client.ServerId,
		TemplateID: request.Client.TemplateId,
	})
	if err != nil {
		return nil, errors.Wrap(err, "ucTemp.UpdateTemplate failed")
	}
	return &v1.UpdateClientResponse{
		Client: entityClientToProtobuf(result),
	}, nil
}

func (m *mailGRPCServer) GetHistories(ctx context.Context, request *v1.GetHistoriesRequest) (*v1.GetHistoriesResponse, error) {
	results, err := m.ucHistory.GetHistories(ctx, request.Limit, request.Offset)
	if err != nil {
		return nil, errors.Wrap(err, "ucServer.GetServers failed")
	}
	return &v1.GetHistoriesResponse{
		Histories: lo.Map(results, func(item *domain.History, _ int) *v1.History {
			return entityHistoryToProtobuf(item)
		}),
	}, nil
}
func (m *mailGRPCServer) TestSendMail(ctx context.Context, request *v1.TestSendMailRequest) (*v1.TestSendMailResponse, error) {
	err := m.ucSendMail.TestSendMail(ctx, request.Host, request.Port, request.AuthProtocol, request.Username, request.Password, request.TlsType,
		request.FromName, request.FromAddress, request.IdleTimeout, request.MaxConnections, request.Retries, request.WaitTimeout, request.To)
	if err != nil {
		return nil, errors.Wrap(err, "router.TestSendMail failed")
	}
	return nil, nil
}
func (m *mailGRPCServer) SendMail(ctx context.Context, request *v1.SendMailRequest) (*v1.SendMailResponse, error) {
	apiKey, err := utils.GetKeyMetadata(ctx, constant.ApiKey)
	if err != nil {
		return nil, errors.Wrap(err, "utils.GetKeyMetadata failed")
	}
	if apiKey == "" {
		return nil, errors.New("Invalid Request")
	}
	err = m.ucSendMail.SendMail(ctx, &domain.History{
		ApiKey:  apiKey,
		To:      request.To,
		Subject: request.Subject,
		Cc:      request.Cc,
		Bcc:     request.Bcc,
		Content: request.Content.AsMap(),
		Status:  "pending",
	})
	if err != nil {
		return nil, errors.Wrap(err, "router.SendMail failed")
	}
	return nil, nil
}
func entityHistoryToProtobuf(entity *domain.History) *v1.History {
	return &v1.History{
		Id:        entity.ID,
		ApiKey:    entity.ApiKey,
		Subject:   entity.Subject,
		To:        entity.To,
		Cc:        entity.Cc,
		Bcc:       entity.Bcc,
		Content:   utils.ToStruct(entity.Content),
		Status:    entity.Status,
		CreatedAt: timestamppb.New(entity.CreatedAt),
		UpdatedAt: timestamppb.New(entity.UpdatedAt),
	}
}

func entityServerToProtobuf(entity *domain.Server) *v1.Server {
	return &v1.Server{
		Id:             entity.ID,
		Name:           entity.Name,
		Host:           entity.Host,
		Port:           entity.Port,
		AuthProtocol:   entity.AuthProtocol,
		Username:       entity.UserName,
		Password:       entity.Password,
		FromName:       entity.FromName,
		FromAddress:    entity.FromAddress,
		TlsType:        string(entity.TLSType),
		MaxConnections: entity.MaxConnections,
		IdleTimeout:    entity.IdleTimeout,
		Retries:        entity.Retries,
		WaitTimeout:    entity.WaitTimeout,
		IsDefault:      entity.IsDefault,
		CreatedAt:      timestamppb.New(entity.CreatedAt),
		UpdatedAt:      timestamppb.New(entity.UpdatedAt),
	}
}
func entityTemplateToProtobuf(entity *domain.Template) *v1.Template {
	return &v1.Template{
		Id:        entity.ID,
		Name:      entity.Name,
		Status:    entity.Status,
		Html:      entity.Html,
		IsDefault: entity.IsDefault,
		CreatedAt: timestamppb.New(entity.CreatedAt),
		UpdatedAt: timestamppb.New(entity.UpdatedAt),
	}
}
func entityClientToProtobuf(entity *domain.Client) *v1.Client {
	return &v1.Client{
		Id:         entity.ID,
		Name:       entity.Name,
		ServerId:   entity.ServerID,
		TemplateId: entity.TemplateID,
		IsDefault:  entity.IsDefault,
		ApiKey:     entity.ApiKey,
		CreatedAt:  timestamppb.New(entity.CreatedAt),
		UpdatedAt:  timestamppb.New(entity.UpdatedAt),
	}
}
