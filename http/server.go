package http

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"regexp"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	gg "DavisFrench/golang-grocery"
)

const (
	PRODUCECODE_FORMAT_ERROR = "Improperly formatted produce_code!\nPlease follow the format xxxx-xxxx-xxxx-xxxx, where x is an alphanumeric character"
	PRODUCECODE_REGEX        = `^[a-zA-Z\d]{4}-[a-zA-Z\d]{4}-[a-zA-Z\d]{4}-[a-zA-Z\d]{4}$`
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

func verifyProduceCodeFormat(produceCode string) (bool, error) {
	return regexp.MatchString(PRODUCECODE_REGEX, produceCode)
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
		r.Get("/produce", s.getAllProduce)
		r.Get("/produce/{produceCode}", s.getProduceByProduceCode)
		r.Delete("/produce/{produceCode}", s.deleteProduce)
		r.Post("/produce", s.addProduce)
	})

	return r
}

func (s *Server) handlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status":"ok"}` + "\n"))
}

func (s *Server) getAllProduce(w http.ResponseWriter, r *http.Request) {

	all, err := s.groceryService.GetAllProduce()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(all)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func (s *Server) getProduceByProduceCode(w http.ResponseWriter, r *http.Request) {

	produceCode := chi.URLParam(r, "produceCode")
	valid, err := verifyProduceCodeFormat(produceCode)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !valid {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(PRODUCECODE_FORMAT_ERROR))
		return
	}

	produce, err := s.groceryService.GetProduceByCode(produceCode)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(produce)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func (s *Server) deleteProduce(w http.ResponseWriter, r *http.Request) {

	produceCode := chi.URLParam(r, "produceCode")
	valid, err := verifyProduceCodeFormat(produceCode)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !valid {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(PRODUCECODE_FORMAT_ERROR))
		return
	}

	if err := s.groceryService.DeleteProduce(produceCode); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Produce successfully deleted"))
}

func (s *Server) addProduce(w http.ResponseWriter, r *http.Request) {

	/*valid, err := verifyProduceCodeFormat(produceCode)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !valid {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(PRODUCECODE_FORMAT_ERROR)
		return
	}

	if err := s.groceryService.AddProduce(produceCode); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}*/

	w.Write([]byte("Produce Successfully added!"))
}
