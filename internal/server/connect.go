package server

import (
	"github.com/cbuschka/pkgdiff/internal/config"
	_package "github.com/cbuschka/pkgdiff/internal/package"
)

type Server struct {
	channel  Channel
	provider _package.Provider
	config   *config.ServerConfig
}

type PackageList struct {
	config   *config.ServerConfig
	Packages []_package.Package
}

func Connect(config *config.ServerConfig) (*Server, error) {
	server := Server{}
	server.channel = Channel(&LocalChannel{})
	return &server, nil
}

func (server *Server) ListPackages() (*PackageList, error) {
	packages, err := server.provider.ListPackages(server.channel)
	if err != nil {
		return nil, err
	}

	return &PackageList{config: server.config, Packages: packages}, nil
}
