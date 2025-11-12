package mw_components

import (
	"log/slog"
	"net/http"
	"time"
)

func InitLoggerMiddleware(logger *slog.Logger) func(next http.Handler) http.Handler {
	mwFunc := func(next http.Handler) http.Handler {
		log := logger.With(slog.String("component", "mwLogger"))
		fn := func(w http.ResponseWriter, r *http.Request) {
			log := log.With(
				slog.String("path", r.URL.Path),
				slog.String("method", r.Method),
				slog.String("ip", r.RemoteAddr),
			)
			startTime := time.Now()
			defer func() {
				log.Info("request completed", "took time (mcs)", time.Since(startTime).Microseconds())
			}()
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
	return mwFunc
}
