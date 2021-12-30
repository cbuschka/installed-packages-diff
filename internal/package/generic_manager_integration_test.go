package _package

import (
	"github.com/cbuschka/go-pkgdiff/internal/model"
	"github.com/cbuschka/go-pkgdiff/internal/transport"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunDnfLocally(t *testing.T) {

	localTransport := transport.LocalTransport{}
	packages, err := rpmWithDnf.ListPackages(transport.Transport(&localTransport))
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, []model.Package{}, packages)
}
