package server

import (
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
	}
}

// Start APIServer
func (s *APIServer) Start() error {
	if err := s.ConfigureLogger(); err != nil {
		return err
	}

	s.logger.Info("Starting api server")

	return nil
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
