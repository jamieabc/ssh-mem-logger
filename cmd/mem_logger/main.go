package main

import (
	"flag"
	"fmt"
	"github.com/jamieabc/ssh-mem-logger/internal/client"
	"github.com/jamieabc/ssh-mem-logger/internal/parser"
	"github.com/jamieabc/ssh-mem-logger/internal/remoteCmds"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
)

const (
	program = "bitmarkd"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "c", "", "-c config file path")
	flag.Parse()

	// make sure config exist
	if configPath == "" {
		panic("please input config path")
	}
}

type connection struct {
	remoteIP   string
	remotePort int
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

	p := parser.NewParser(configPath)
	c, err := p.Parse()
	if nil != err {
		panic(err)
	}

	cons := make([]connection, 0)

	for _, s := range c.Servers {
		cons = append(cons, connection{
			remoteIP:   s.IP,
			remotePort: s.Port,
			serverName: s.Name,
			session:    nil,
			info: client.Info{
				Username:    s.UserName,
				KeyFilePath: s.KeyPath,
			},
		})
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
