package config

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLoadConfig(t *testing.T) {

	config, err := LoadConfig(strings.NewReader(`version: 'installed-packages-diff/3'
groups:
  db:
    type: rpm
    servers:
      - url: ssh://root@dbdev
        excludes:
          - "missing"
      - url: ssh://root@dblive

  web:
    servers:
      - url: ssh://root@webdev
        excludes:
          - "missing"
      - url: ssh://root@weblive
`))
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, 2, len(config.Groups))
}
