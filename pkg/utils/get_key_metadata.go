package utils

import (
	"context"
	"errors"
	"log/slog"

	"google.golang.org/grpc/metadata"
)

func GetKeyMetadata(ctx context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("missing metadata")
	}
	slog.Info("Metadata::", md)
	values := md.Get(key)
	if len(values) > 0 {
		return values[0], nil
	}
	return "", nil
}
