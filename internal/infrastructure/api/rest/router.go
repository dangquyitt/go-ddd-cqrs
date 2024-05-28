package rest

import (
	"github.com/dangquyitt/go-ddd-cqrs/internal/infrastructure/api/rest/handler"
	"github.com/dangquyitt/go-ddd-cqrs/internal/infrastructure/api/rest/middleware"
)

func (s *Server) routes(h *handler.Handler, m *middleware.Middleware) {
	s.echo.GET("/hello", h.HelloWold)
}
