package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router chi.Router
	Handlers map[string]routeHandler
	WebServerPort string
}

type routeHandler struct {
	method string
	handler http.HandlerFunc
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router: chi.NewRouter(),
		Handlers: make(map[string]routeHandler),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method string, path string, handler http.HandlerFunc) {
	s.Handlers[path] = routeHandler{
		method: method,
		handler: handler,
	}
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, route := range s.Handlers {
		s.Router.Method(route.method, path, route.handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}