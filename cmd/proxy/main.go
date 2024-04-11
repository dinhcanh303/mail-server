package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	v1m "github.com/dinhcanh303/mail-server/api/mail/v1"
	"github.com/dinhcanh303/mail-server/cmd/proxy/config"
	"github.com/dinhcanh303/mail-server/pkg/logger"
	"github.com/golang/glog"
	gatewayRuntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func newGateway(
	ctx context.Context,
	cfg *config.Config,
	opts []gatewayRuntime.ServeMuxOption) (http.Handler, error) {
	mailEndpoint := fmt.Sprintf("%s:%d", cfg.MailHost, cfg.MailPort)

	mux := gatewayRuntime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := v1m.RegisterGroupServiceHandlerFromEndpoint(ctx, mux, mailEndpoint, dialOpts)
	if err != nil {
		return nil, err
	}
	return mux, nil
}
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if origin := r.Header.Get("Origin"); origin != "" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				slog.Info("Access-Control-Request-Method", r.Header.Get("Access-Control-Request-Method"))
				if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
					preflightHandler(w, r)
					return
				}
			}
			h.ServeHTTP(w, r)
		})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	headers := []string{"*"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	slog.Info("preflight request", "http_path", r.URL.Path)
}

func withLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Run Request", "http_method", r.Method, "http_url", r.URL)
		h.ServeHTTP(w, r)
	})
}
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	cfg, err := config.NewConfig()
	if err != nil {
		glog.Fatalf("Config error: %s", err)
	}
	slog.Info("Config::", cfg)
	// set up logrus
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logger.ConvertLogLevel(cfg.Log.Level))

	// integrate Logrus with the slog logger
	slog.New(logger.NewLogrusHandler(logrus.StandardLogger()))

	mux := http.NewServeMux()

	gw, err := newGateway(ctx, cfg, nil)
	if err != nil {
		slog.Error("failed to create a new gateway", err)
	}
	mux.Handle("/", gw)

	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler: allowCORS(withLogger(mux)),
	}

	//goroutine
	go func() {
		<-ctx.Done()
		slog.Info("shutting down the http server")
		if err := s.Shutdown(context.Background()); err != nil {
			slog.Error("failed to shutdown http server", err)
		}
	}()

	slog.Info("🌏 start listening...", "address", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err := s.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to listen and serve", err)
	}
}
