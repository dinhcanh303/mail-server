// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/dinhcanh303/mail-server/cmd/mail/config"
	"github.com/dinhcanh303/mail-server/internal/mail/app/router"
	"github.com/dinhcanh303/mail-server/internal/mail/events/handlers"
	"github.com/dinhcanh303/mail-server/internal/mail/infras"
	"github.com/dinhcanh303/mail-server/internal/mail/infras/repo"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/client"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/history"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/sendmail"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/server"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/template"
	"github.com/dinhcanh303/mail-server/pkg/config"
	"github.com/dinhcanh303/mail-server/pkg/mail"
	"github.com/dinhcanh303/mail-server/pkg/postgres"
	"github.com/dinhcanh303/mail-server/pkg/rabbitmq"
	"github.com/dinhcanh303/mail-server/pkg/rabbitmq/consumer"
	"github.com/dinhcanh303/mail-server/pkg/rabbitmq/publisher"
	"github.com/dinhcanh303/mail-server/pkg/redis"
	"github.com/dinhcanh303/mail-server/pkg/token"
	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

// Injectors from wire.go:

func InitApp(cfg *config.Config, cfg2 *configs.Redis, dbConnStr postgres.DBConnString, rabbitMQConnStr rabbitmq.RabbitMQConnStr, grpcServer *grpc.Server) (*App, func(), error) {
	dbEngine, cleanup, err := dbEngineFunc(dbConnStr)
	if err != nil {
		return nil, nil, err
	}
	redisEngine, cleanup2, err := redisEngineFunc(cfg2)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	serverRepo := repo.NewServerRepo(dbEngine)
	useCase := server.NewUseCase(redisEngine, serverRepo)
	templateRepo := repo.NewTemplateRepo(dbEngine)
	templateUseCase := template.NewUseCase(redisEngine, templateRepo)
	clientRepo := repo.NewClientRepo(dbEngine)
	clientUseCase := client.NewUseCase(redisEngine, clientRepo, useCase, templateUseCase)
	connection, cleanup3, err := rabbitMQFunc(rabbitMQConnStr)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	eventPublisher, err := publisher.NewPublisher(connection)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	mailEventPublisher := infras.NewMailEventPublisher(eventPublisher)
	sendmailUseCase := sendmail.NewUseCase(redisEngine, mailEventPublisher)
	mailServiceServer := router.NewMailGRPCServer(grpcServer, templateUseCase, useCase, clientUseCase, sendmailUseCase, cfg)
	historyRepo := repo.NewHistoryRepo(dbEngine)
	historyUseCase := history.NewUseCase(redisEngine, historyRepo)
	emailSender := mailServer()
	mailEventHandler := handlers.NewMailEventHandler(historyUseCase, emailSender, clientUseCase)
	eventConsumer, err := consumer.NewConsumer(connection)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	app := New(cfg, dbEngine, useCase, templateUseCase, clientUseCase, sendmailUseCase, mailServiceServer, mailEventHandler, eventPublisher, eventConsumer)
	return app, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

func dbEngineFunc(url postgres.DBConnString) (postgres.DBEngine, func(), error) {
	db, err := postgres.NewPostgresDB(url)
	if err != nil {
		return nil, nil, err
	}
	return db, func() { db.Close() }, nil
}

func jwtFunc() token.JWT {
	jwt := token.NewJWTMaker()
	return jwt
}

func rabbitMQFunc(url rabbitmq.RabbitMQConnStr) (*amqp091.Connection, func(), error) {
	conn, err := rabbitmq.NewRabbitMQConn(url)
	if err != nil {
		return nil, nil, err
	}
	return conn, func() { conn.Close() }, nil
}

func redisEngineFunc(config2 *configs.Redis) (redis.RedisEngine, func(), error) {
	redis2, err := redis.NewRedisClient(config2)
	if err != nil {
		return nil, nil, err
	}
	return redis2, func() {
		redis2.Close()
	}, nil
}

func mailServer() mail.EmailSender {
	return mail.NewEmailSender()
}
