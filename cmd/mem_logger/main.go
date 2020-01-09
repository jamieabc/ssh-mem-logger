package main

import (
	"fmt"
	"github.com/jamieabc/ssh-mem-logger/internal/client"
	"github.com/jamieabc/ssh-mem-logger/internal/remoteCmds"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
)

const (
	program = "bitmarkd"
)

type connection struct {
	remoteIP   string
	remotePort uint
	serverName string
	session    *ssh.Session
	info       client.Info
}

func (c connection) con() string {
	return fmt.Sprintf("%s:%d", c.remoteIP, c.remotePort)
}

func main() {
	top, err := remoteCmds.New("top", program)
	if nil != err {
		panic(err)
	}

	info := client.Info{
		Username:    "ec2-user",
		KeyFilePath: "/Users/Aaron/.aws/aaron_key.pem",
	}

	cons := []connection{
		{
			remoteIP:   "13.114.24.6",
			remotePort: 22,
			serverName: "test",
			session:    nil,
			info:       info,
		},
		{
			remoteIP:   "54.238.201.3",
			remotePort: 22,
			serverName: "test2",
			session:    nil,
			info:       info,
		},
	}

	setupConnections(cons)

	for _, c := range cons {
		fmt.Printf("server: %s\n", c.serverName)
		err = c.session.Run(top.String())
		if nil != err {
			panic(err)
		}
	}

	closeConnections(cons)
}

func closeConnections(cons []connection) {
	for _, c := range cons {
		_ = c.session.Close()
	}
}

func setupConnections(conns []connection) {
	for i := range conns {
		sh, err := client.NewSSH(conns[i].info)
		if nil != err {
			panic(err)
		}

		conn, err := sh.Connect("tcp", conns[i].con())
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

		conns[i].session = sess
	}
}
