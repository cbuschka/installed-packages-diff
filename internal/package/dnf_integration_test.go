package _package

import (
  "github.com/cbuschka/pkgdiff/internal/server"
  "github.com/stretchr/testify/assert"
	"testing"
)

func TestRunDnfLocally(t *testing.T) {

	localChannel := &server.LocalChannel{}
	dnf := Dnf{}
  packages, err := dnf.ListPackages(server.Channel(localChannel))
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, []Package{}, packages)
}
