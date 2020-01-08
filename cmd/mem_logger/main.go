package main

import (
	"github.com/jamieabc/ssh-mem-logger/internal/ssh"
	"io"
	"os"
)

func main() {
	s, err := ssh.NewSSH("ec2-user", "/Users/Aaron/.aws/aaron_key.pem")
	if nil != err {
		panic(err)
	}
	conn, err := s.Dial("tcp", "13.114.24.6:22")
	if nil != err {
		panic(err)
	}
	defer conn.Close()

	sess, err := conn.NewSession()
	if nil != err {
		panic(err)
	}
	defer sess.Close()

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

	err = sess.Run("ls")
	if nil != err {
		panic(err)
	}
}
