package transport

import (
  "github.com/cbuschka/go-pkgdiff/internal/config"
)

type Server struct {
	transport Transport
	config    *config.ServerConfig
}

func Connect(config *config.ServerConfig) (Transport, error) {
	return Transport(&LocalTransport{}), nil
}
