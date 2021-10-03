package _package

import (
  "github.com/cbuschka/pkgdiff/internal/server"
  "github.com/stretchr/testify/assert"
	"testing"
)

func TestDetectProvider(t *testing.T) {
	localChannel := server.Channel(&server.LocalChannel{})
	_, stderr, err := localChannel.Run("echo", "hello world")
	if err != nil {
		t.Fatal(err)
		return
	}

	_, err = DetectProvider(localChannel)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, "", stderr)
}
