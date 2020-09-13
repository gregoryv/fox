package foxhttp

import (
	"net/http"
	"time"

	"github.com/gregoryv/fox"
)

// NewRouteLog returns RouteLog using the package logger.
func NewRouteLog(l fox.Logger) *RouteLog {
	return &RouteLog{Logger: l}
}

type RouteLog struct {
	fox.Logger
}

// MiddlewareFunc returns a middleware that logs request and response status.
func (me *RouteLog) MiddlewareFunc(next http.HandlerFunc) http.Handler {
	return me.Middleware(next)
}

// Middleware returns a middleware that logs request and response status.
func (me *RouteLog) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := statusRecorder{w, 200}
		next.ServeHTTP(&rec, r)
		me.Log(r.Method, r.URL, rec.status, time.Since(start))
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.ResponseWriter.WriteHeader(code)
	rec.status = code
}
