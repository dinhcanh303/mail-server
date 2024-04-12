package router

import (
	"context"

	v1 "github.com/dinhcanh303/mail-server/api/mail/v1"
	"github.com/dinhcanh303/mail-server/cmd/mail/config"
	"github.com/dinhcanh303/mail-server/internal/mail/domain"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/server"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/template"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type mailGRPCServer struct {
	v1.UnimplementedMailServiceServer
	ucTemp   template.UseCase
	ucServer server.UseCase
	cfg      *config.Config
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
	cfg *config.Config,
) v1.MailServiceServer {
	svc := mailGRPCServer{
		ucTemp:   ucTemp,
		ucServer: ucServer,
		cfg:      cfg,
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
	model := domain.NewServer(request.Name,
		request.Host,
		request.Port,
		request.Username,
		request.Password,
		request.Tls,
		request.SkipTls,
		request.MaxConnections,
		request.Retries,
		request.IdleTimeout,
		request.WaitTimeout)
	result, err := m.ucServer.CreateServer(ctx, model)
	if err != nil {
		return nil, errors.Wrap(err, "ucServer.CreateServer failed")
	}
	return &v1.CreateServerResponse{
		Server: entityToProtobuf(result),
	}, nil
}
func (m *mailGRPCServer) GetServer(ctx context.Context, request *v1.GetServerRequest) (*v1.GetServerResponse, error) {
	result, err := m.ucServer.GetServer(ctx, request.Id)
	if err != nil {
		return nil, errors.Wrap(err, "ucServer.GetServer failed")
	}
	return &v1.GetServerResponse{
		Server: entityToProtobuf(result),
	}, nil
}
func (m *mailGRPCServer) GetServers(ctx context.Context, request *v1.GetServersRequest) (*v1.GetServersResponse, error) {
	results, err := m.ucServer.GetServers(ctx, request.Limit, request.Offset)
	if err != nil {
		return nil, errors.Wrap(err, "ucServer.GetServers failed")
	}
	return &v1.GetServersResponse{
		Servers: lo.Map(results, func(item *domain.Server, _ int) *v1.Server {
			return entityToProtobuf(item)
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
		UserName:       request.Server.Username,
		Password:       request.Server.Password,
		TLS:            domain.TLSType(request.Server.Tls),
		SkipTLSVerify:  request.Server.SkipTls,
		MaxConnections: request.Server.MaxConnections,
		Retries:        request.Server.Retries,
		IdleTimeout:    request.Server.IdleTimeout,
		WaitTimeout:    request.Server.WaitTimeout,
	})
	if err != nil {
		return nil, errors.Wrap(err, "ucServer.UpdateServer failed")
	}
	return &v1.UpdateServerResponse{
		Server: entityToProtobuf(result),
	}, nil
}
func (m *mailGRPCServer) CreateTemplate(ctx context.Context, request *v1.CreateTemplateRequest) (*v1.CreateTemplateResponse, error) {
	model := domain.NewTemplate(request.Name, request.Status, request.Html)
	result, err := m.ucTemp.CreateTemplate(ctx, model)
	if err != nil {
		return nil, errors.Wrap(err, "ucTemp.CreateTemplate failed")
	}
	return &v1.CreateTemplateResponse{
		Template: entityTemplateToProtobuf(result),
	}, nil
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

func entityToProtobuf(entity *domain.Server) *v1.Server {
	return &v1.Server{
		Id:             entity.ID,
		Name:           entity.Name,
		Host:           entity.Host,
		Port:           entity.Port,
		Username:       entity.UserName,
		Password:       entity.Password,
		Tls:            string(entity.TLS),
		SkipTls:        entity.SkipTLSVerify,
		MaxConnections: entity.MaxConnections,
		IdleTimeout:    entity.IdleTimeout,
		Retries:        entity.Retries,
		WaitTimeout:    entity.WaitTimeout,
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
		CreatedAt: timestamppb.New(entity.CreatedAt),
		UpdatedAt: timestamppb.New(entity.UpdatedAt),
	}
}
