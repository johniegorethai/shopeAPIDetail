package http

import (
	"net/http"

	"shopeAPIDetail/pkg/grace"
)

// SkeletonHandler ...
type OrderPelapakHandler interface {
	OrderPelapakHandler(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	server       *http.Server
	OrderPelapak OrderPelapakHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	return grace.Serve(port, s.Handler())
}
