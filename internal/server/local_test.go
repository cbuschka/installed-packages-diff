package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunEchoLocally(t *testing.T) {

	localChannel := LocalChannel{}
	stdout, stderr, err := localChannel.Run("echo", "hello world")
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, "", stderr)
	assert.Equal(t, "hello world\n", stdout)
}
