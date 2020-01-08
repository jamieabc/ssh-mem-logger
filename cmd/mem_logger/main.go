package main

import (
	"github.com/jamieabc/ssh-mem-logger/internal/client"
	"github.com/jamieabc/ssh-mem-logger/internal/remoteCmds"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
)

const (
	program = "bitmarkd"
)

func main() {
	top, err := remoteCmds.New("top", program)
	if nil != err {
		panic(err)
	}

	sessions := newConnections([]string{"13.114.24.6:22", "54.238.201.3:22"})

	for _, sess := range sessions {
		err = sess.Run(top.String())
		if nil != err {
			panic(err)
		}
	}

}

func newConnections(conns []string) []*ssh.Session {
	result := make([]*ssh.Session, 0)

	for _, c := range conns {
		sh, err := client.NewSSH("ec2-user", "/Users/Aaron/.aws/aaron_key.pem")
		if nil != err {
			panic(err)
		}

		conn, err := sh.Connect("tcp", c)
		if nil != err {
			panic(err)
		}

		sess, err := conn.NewSession()
		if nil != err {
			panic(err)
		}

		stdout, err := sess.StdoutPipe()
		if nil != err {
			panic(err)
		}
		go io.Copy(os.Stdout, stdout)

		stderr, err := sess.StderrPipe()
		if nil != err {
			panic(err)
		}
		go io.Copy(os.Stderr, stderr)

		result = append(result, sess)
	}

	return result
}
