package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/dinhcanh303/mail-server/cmd/mail/config"
	"github.com/dinhcanh303/mail-server/internal/mail/app"
	configs "github.com/dinhcanh303/mail-server/pkg/config"
	"github.com/dinhcanh303/mail-server/pkg/logger"
	"github.com/dinhcanh303/mail-server/pkg/postgres"
	"github.com/dinhcanh303/mail-server/pkg/rabbitmq"
	"github.com/sirupsen/logrus"
	"go.uber.org/automaxprocs/maxprocs"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

func main() {
	// set GOMAXPROCS
	_, err := maxprocs.Set()
	if err != nil {
		slog.Error("failed set max procs", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("failed get config", err)
	}
	cfgRedis, err := configs.NewConfigRedis()
	if err != nil {
		slog.Error("Failed get config", err)
	}

	slog.Info("⚡ init app", "name", cfg.Name, "version", cfg.Version)

	// set up logrus
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logger.ConvertLogLevel(cfg.Log.Level))

	// integrate Logrus with the slog logger
	slog.New(logger.NewLogrusHandler(logrus.StandardLogger()))

	server := grpc.NewServer()

	go func() {
		defer server.GracefulStop()
		<-ctx.Done()
	}()

	cleanup := prepareApp(ctx, cancel, cfg, cfgRedis, server)

	// gRPC Server.
	address := fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)
	network := "tcp"

	l, err := net.Listen(network, address)
	if err != nil {
		slog.Error("failed to listen to address", err, "network", network, "address", address)
		cancel()
		<-ctx.Done()
	}

	slog.Info("🌏 start server...", "address", address)

	defer func() {
		if err1 := l.Close(); err != nil {
			slog.Error("failed to close", err1, "network", network, "address", address)
			<-ctx.Done()
		}
	}()

	err = server.Serve(l)
	if err != nil {
		slog.Error("failed start gRPC server", err, "network", network, "address", address)
		cancel()
		<-ctx.Done()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		cleanup()
		slog.Info("signal.Notify", v)
	case done := <-ctx.Done():
		cleanup()
		slog.Info("ctx.Done", "app done", done)
	}
}

func prepareApp(ctx context.Context, cancel context.CancelFunc, cfg *config.Config, cfgRedis *configs.Redis, server *grpc.Server) func() {
	_, cleanup, err := app.InitApp(cfg, cfgRedis, postgres.DBConnString(cfg.PG.DbURL), rabbitmq.RabbitMQConnStr(cfg.RabbitMQ.URL), server)
	if err != nil {
		slog.Error("failed init app", err)
		cancel()
		<-ctx.Done()
	}
	// a.Consumer.Configure(
	// 	consumer.ExChangeName("search-exchange"),
	// 	consumer.QueueName("search-queue"),
	// 	consumer.BindingKey("search-routing-key"),
	// 	consumer.ConsumerTag("search-consumer"),
	// )

	// go func() {
	// 	err1 := a.Consumer.StartConsumer(a.Worker)
	// 	if err1 != nil {
	// 		slog.Error("failed to start Consumer", err1)
	// 		cancel()
	// 		<-ctx.Done()
	// 	}
	// }()

	return cleanup
}
