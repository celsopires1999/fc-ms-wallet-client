package webserver

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Verb     string
	Function http.HandlerFunc
}

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]Handler
	WebServerPort string
}

func NewWebServer(webServerPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]Handler),
		WebServerPort: webServerPort,
	}
}

func (s *WebServer) AddHandler(path string, handler Handler) {
	s.Handlers[path] = handler
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, handler := range s.Handlers {
		if handler.Verb == "GET" {
			s.Router.Get(path, handler.Function)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
