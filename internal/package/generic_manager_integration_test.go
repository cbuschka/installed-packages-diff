package _package

import (
	"github.com/cbuschka/go-pkgdiff/internal/model"
	transportPkg "github.com/cbuschka/go-pkgdiff/internal/transport"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunRpmWithDnfLocally(t *testing.T) {

	localTransport := transportPkg.Transport(&transportPkg.LocalTransport{})

	available, err := rpmWithDnf.IsAvailable(localTransport)
	if err != nil || !available {
		return
	}

	packages, err := rpmWithDnf.ListPackages(localTransport)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, []model.Package{}, packages)
}
