package diff

import (
	"github.com/cbuschka/go-pkgdiff/internal/model"
)

type DiffEntry struct {
	Package model.Package
}

func Diff(packageLists []model.PackageList) error {
	return nil
}
