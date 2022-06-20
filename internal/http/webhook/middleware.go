package webhook

import (
	"net/http"

	"github.com/slok/go-http-metrics/middleware"
	middlewarestd "github.com/slok/go-http-metrics/middleware/std"
)

// measuredHandler wraps a handler and measures the request handled
// by this handler.
func (h handler) measuredHandler(next http.Handler) http.Handler {
	mdlw := middleware.New(middleware.Config{Recorder: h.metrics})
	return middlewarestd.Handler("", mdlw, next)
}
