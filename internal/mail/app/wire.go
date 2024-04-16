//go:build wireinject
// +build wireinject

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

	configs "github.com/dinhcanh303/mail-server/pkg/config"
	"github.com/dinhcanh303/mail-server/pkg/mail"
	"github.com/dinhcanh303/mail-server/pkg/postgres"
	"github.com/dinhcanh303/mail-server/pkg/rabbitmq"
	"github.com/dinhcanh303/mail-server/pkg/rabbitmq/consumer"
	"github.com/dinhcanh303/mail-server/pkg/rabbitmq/publisher"
	"github.com/dinhcanh303/mail-server/pkg/redis"
	"github.com/dinhcanh303/mail-server/pkg/token"
	"github.com/google/wire"
	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

func InitApp(
	cfg *config.Config,
	cfg2 *configs.Redis,
	dbConnStr postgres.DBConnString,
	rabbitMQConnStr rabbitmq.RabbitMQConnStr,
	grpcServer *grpc.Server,
) (*App, func(), error) {
	panic(wire.Build(
		New,
		dbEngineFunc,
		// jwtFunc,
		rabbitMQFunc,
		mailServer,
		publisher.EventPublisherSet,
		consumer.EventConsumerSet,
		redisEngineFunc,
		router.MailGRPCServerSet,
		server.UseCaseSet,
		template.UseCaseSet,
		client.UseCaseSet,
		sendmail.UseCaseSet,
		history.UseCaseSet,
		repo.ServerRepoSet,
		repo.TemplateRepoSet,
		repo.ClientRepoSet,
		repo.HistoryRepoSet,
		infras.MailEventPublisherSet,
		handlers.MailEventHandlerSet,
	))
}
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
func redisEngineFunc(config *configs.Redis) (redis.RedisEngine, func(), error) {
	redis, err := redis.NewRedisClient(config)
	if err != nil {
		return nil, nil, err
	}
	return redis, func() { redis.Close() }, nil
}

func mailServer() mail.EmailSender {
	return mail.NewEmailSender()
}
