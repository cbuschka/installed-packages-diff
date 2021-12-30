package diff

import (
	"github.com/cbuschka/go-pkgdiff/internal/model"
	"github.com/cbuschka/go-pkgdiff/internal/transport"
)

type DiffEntry struct {
	Package model.Package
}

func Diff(packageLists []transport.PackageList) error {
	return nil
}
