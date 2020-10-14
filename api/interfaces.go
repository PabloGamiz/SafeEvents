package api

import "net/http"

// Server represents a router for all requests
type Server interface {
	Router() http.Handler
}
