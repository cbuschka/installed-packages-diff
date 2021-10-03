package diff

import (
	_package "github.com/cbuschka/pkgdiff/internal/package"
	"github.com/cbuschka/pkgdiff/internal/server"
)

type DiffEntry struct {
	Package _package.Package
}

func Diff(packageLists []*server.PackageList) error {

}
