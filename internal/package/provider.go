package _package

import "github.com/cbuschka/pkgdiff/internal/server"

type Provider interface {
	ListPackages(channel server.Channel) ([]Package, error)
	IsAvailable(channel server.Channel) (bool, error)
}
