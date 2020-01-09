package parser_test

import (
	"github.com/jamieabc/ssh-mem-logger/internal/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testConfigPath = "fixtures/config.lua"
)

func TestConfig_Parse(t *testing.T) {
	p := parser.NewParser(testConfigPath)
	c, err := p.Parse()
	assert.Nil(t, err, "wrong parse")

	server1 := parser.Server{
		UserName: "user1",
		KeyPath:  "/home/test_key1.pem",
		IP:       "1.2.3.4",
		Port:     22,
		Name:     "server1",
	}
	server2 := parser.Server{
		UserName: "user2",
		KeyPath:  "/home/test_key2.pem",
		IP:       "5.6.7.8",
		Port:     22,
		Name:     "server2",
	}

	assert.Equal(t, 2, len(c.Servers), "wrong config count")
	assert.Equal(t, server1, c.Servers[0], "wrong first server")
	assert.Equal(t, server2, c.Servers[1], "wrong second server")
}
