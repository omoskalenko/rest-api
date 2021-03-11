package server

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/omoskalenko/rest-api/internal/app/store"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  store.New(config.Store),
	}
}

// Start APIServer
func (s *APIServer) Start() error {
	if err := s.ConfigureLogger(); err != nil {
		return err
	}

	s.ConfigureRouter()

	if err := s.ConfigureStore(); err != nil {
		return err
	}
	s.logger.Info("Starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// ConfigureLogger ...
func (s *APIServer) ConfigureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

// ConfigureRouter ...
func (s *APIServer) ConfigureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

// ConfigureStore ...
func (s *APIServer) ConfigureStore() error {
	st := store.New(s.config.Store)

	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
