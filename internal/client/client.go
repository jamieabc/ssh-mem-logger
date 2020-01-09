package client

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

type SSH interface {
	Connect(string, string) (*ssh.Client, error)
}

type client struct {
	config *ssh.ClientConfig
}

type Info struct {
	Username    string
	KeyFilePath string
}

func (c client) Connect(tcp, host string) (*ssh.Client, error) {
	return ssh.Dial(tcp, host, c.config)
}

func NewSSH(info Info) (SSH, error) {
	if info.Username == "" || info.KeyFilePath == "" {
		return nil, fmt.Errorf("wrong input")
	}

	return &client{
		config: &ssh.ClientConfig{
			User: info.Username,
			Auth: []ssh.AuthMethod{
				publicKey(info.KeyFilePath),
			},
			HostKeyCallback:   ssh.InsecureIgnoreHostKey(),
			HostKeyAlgorithms: nil,
		},
	}, nil
}

func publicKey(path string) ssh.AuthMethod {
	key, err := ioutil.ReadFile(path)
	if nil != err {
		panic(err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if nil != err {
		panic(err)
	}

	return ssh.PublicKeys(signer)
}
