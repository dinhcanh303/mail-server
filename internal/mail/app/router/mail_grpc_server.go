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
	result, err := m.ucServer.CreateServer(ctx, &domain.Server{
		Name:           request.Name,
		Host:           request.Host,
		Port:           request.Port,
		UserName:       request.Username,
		Password:       request.Password,
		TLS:            domain.TLSType(request.Tls),
		SkipTLSVerify:  request.SkipTls,
		MaxConnections: request.MaxConnections,
		Retries:        request.Retries,
		IdleTimeout:    request.IdleTimeout,
		WaitTimeout:    request.WaitTimeout,
	})
	if err != nil {
		return nil, errors.Wrap(err, "ucServer.CreateServer failed")
	}
	return &v1.CreateServerResponse{
		Server: entityToProtobuf(result),
	}, nil
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
