package remoteCmds_test

import (
	"github.com/jamieabc/ssh-mem-logger/internal/remoteCmds"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTopWhenOrder(t *testing.T) {
	tt, err := remoteCmds.New("top", "test", "res")
	assert.Nil(t, err, "wrong error")
	assert.Equal(t, "top -o res | grep -B 1 test", tt.String(), "wrong command")
}

func TestNewTopWhenUser(t *testing.T) {
	tt, err := remoteCmds.New("top", "test", "res", "ec2-user")
	assert.Nil(t, err, "wrong error")
	assert.Equal(t, "top -o res -U ec2-user | grep -B 1 test", tt.String(), "wrong command")
}

func TestNewTopWhenOrderAndUser(t *testing.T) {
	tt, err := remoteCmds.New("top", "test", "res", "ec2-user")
	assert.Nil(t, err, "wrong error")
	assert.Equal(t, "top -o res -U ec2-user | grep -B 1 test", tt.String(), "wrong command")
}
