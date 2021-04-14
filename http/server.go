package http

import (
	"net"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	gg "DavisFrench/golang-grocery"
)

type Server struct {
	ln net.Listener

	Addr        string
	Host        string
	Recoverable bool

	groceryService gg.GroceryService
}

func NewServer(groceryService gg.GroceryService) *Server {
	return &Server{
		Recoverable:    true,
		groceryService: groceryService,
	}
}

func (s *Server) Open() error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	s.ln = ln

	return http.Serve(s.ln, s.router())
}

func (s *Server) Close() error {
	if s.ln != nil {
		s.ln.Close()
	}
	return nil
}

func (s *Server) router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	if s.Recoverable {
		r.Use(middleware.Recoverer)
	}

	r.Route("/grocery", func(r chi.Router) {
		r.Get("/ping", s.handlePing)
	})

	return r
}

func (s *Server) handlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status":"ok"}` + "\n"))
}
