package internalhttp

import (
	"net/http"

	"github.com/sterligov/otus_highload/dating/internal/config"
)

func NewHandler(cfg *config.Config) (http.Handler, error) {

	mux := http.NewServeMux()
	//handler := HeadersMiddleware(gw)
	//handler = LoggingMiddleware(handler)
	//mux.Handle("/", handler)

	return mux, nil
}
