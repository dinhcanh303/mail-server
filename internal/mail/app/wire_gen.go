// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/dinhcanh303/mail-server/cmd/mail/config"
	"github.com/dinhcanh303/mail-server/internal/mail/app/router"
	"github.com/dinhcanh303/mail-server/internal/mail/infras/repo"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/client"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/server"
	"github.com/dinhcanh303/mail-server/internal/mail/usecases/template"
	"github.com/dinhcanh303/mail-server/pkg/config"
	"github.com/dinhcanh303/mail-server/pkg/postgres"
	"github.com/dinhcanh303/mail-server/pkg/rabbitmq"
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
	clientUseCase := client.NewUseCase(redisEngine, clientRepo)
	mailServiceServer := router.NewMailGRPCServer(grpcServer, templateUseCase, useCase, clientUseCase, cfg)
	app := New(cfg, dbEngine, useCase, templateUseCase, clientUseCase, mailServiceServer)
	return app, func() {
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
