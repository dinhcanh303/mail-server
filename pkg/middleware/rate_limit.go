package middleware

import (
	"net/http"

	configs "github.com/dinhcanh303/mail-server/pkg/config"
	"golang.org/x/time/rate"
)

// Simple Rate Limit With All Request Per Second For A API
func RateLimit(cfg *configs.Request, next http.Handler) http.Handler {
	limiter := rate.NewLimiter(rate.Limit(cfg.RequestPerSecond), cfg.RequestBurst)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() {
			next.ServeHTTP(w, r)
		} else {
			return
		}
	})
}
