package transport

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunEchoLocally(t *testing.T) {

	localChannel := LocalTransport{}
	stdout, stderr, err := localChannel.ExecCommand("echo", "hello world")
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, "", stderr)
	assert.Equal(t, "hello world\n", stdout)
}
