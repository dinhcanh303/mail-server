//go:build wireinject
// +build wireinject

package app

import (
	"github.com/dinhcanh303/mail-server/cmd/auth/config"

	configs "github.com/dinhcanh303/mail-server/pkg/config"
	"github.com/dinhcanh303/mail-server/pkg/postgres"
	"github.com/dinhcanh303/mail-server/pkg/rabbitmq"
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
		jwtFunc,
		rabbitMQFunc,
		redisEngineFunc,
		publisher.EventPublisherSet,
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
