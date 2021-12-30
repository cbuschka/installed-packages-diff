package _package

import (
	"github.com/cbuschka/go-pkgdiff/internal/model"
	"github.com/cbuschka/go-pkgdiff/internal/transport"
)

type PackageManager interface {
	ListPackages(transport transport.Transport) ([]model.Package, error)
	IsAvailable(transport transport.Transport) (bool, error)
}
