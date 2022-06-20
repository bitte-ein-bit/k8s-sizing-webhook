package webhook

import (
	"fmt"
	"net/http"

	"github.com/bitte-ein-bit/k8s-sizing-webhook/internal/log"
	"github.com/bitte-ein-bit/k8s-sizing-webhook/internal/mutation/mark"
	"github.com/bitte-ein-bit/k8s-sizing-webhook/internal/mutation/mem"
)

// Config is the handler configuration.
type Config struct {
	MetricsRecorder MetricsRecorder
	Marker          mark.Marker
	MemoryFixer     mem.Fixer
	Logger          log.Logger
}

func (c *Config) defaults() error {
	if c.Marker == nil {
		return fmt.Errorf("marker is required")
	}

	if c.MetricsRecorder == nil {
		c.MetricsRecorder = dummyMetricsRecorder
	}

	if c.Logger == nil {
		c.Logger = log.Dummy
	}

	return nil
}

type handler struct {
	marker      mark.Marker
	memoryFixer mem.Fixer
	handler     http.Handler
	metrics     MetricsRecorder
	logger      log.Logger
}

// New returns a new webhook handler.
func New(config Config) (http.Handler, error) {
	err := config.defaults()
	if err != nil {
		return nil, fmt.Errorf("handler configuration is not valid: %w", err)
	}

	mux := http.NewServeMux()

	h := handler{
		handler:     mux,
		marker:      config.Marker,
		memoryFixer: config.MemoryFixer,
		metrics:     config.MetricsRecorder,
		logger:      config.Logger.WithKV(log.KV{"service": "webhook-handler"}),
	}

	// Register all the routes with our router.
	err = h.routes(mux)
	if err != nil {
		return nil, fmt.Errorf("could not register routes on handler: %w", err)
	}

	// Register root handler middlware.
	h.handler = h.measuredHandler(h.handler) // Add metrics middleware.

	return h, nil
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}
