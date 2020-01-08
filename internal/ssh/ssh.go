package ssh

import (
	"fmt"
	s "golang.org/x/crypto/ssh"
	"io/ioutil"
)

type SSH interface {
	Dial(string, string) (*s.Client, error)
}

type ssh struct {
	config *s.ClientConfig
}

func (sh ssh) Dial(tcp, host string) (*s.Client, error) {
	return s.Dial(tcp, host, sh.config)
}

func NewSSH(username, keyFilenamePath string) (SSH, error) {
	if username == "" || keyFilenamePath == "" {
		return nil, fmt.Errorf("wrong input")
	}

	return &ssh{
		config: &s.ClientConfig{
			User: username,
			Auth: []s.AuthMethod{
				publicKey(keyFilenamePath),
			},
			HostKeyCallback:   s.InsecureIgnoreHostKey(),
			HostKeyAlgorithms: nil,
		},
	}, nil
}

func publicKey(path string) s.AuthMethod {
	key, err := ioutil.ReadFile(path)
	if nil != err {
		panic(err)
	}
	signer, err := s.ParsePrivateKey(key)
	if nil != err {
		panic(err)
	}

	return s.PublicKeys(signer)
}
