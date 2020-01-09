package remoteCmds

import "fmt"

type RemoteCmds interface {
	fmt.Stringer
}

func New(cmd string, args ...interface{}) (RemoteCmds, error) {
	switch cmd {
	case "top":
		return newTop(args)
	default:
		return nil, fmt.Errorf("unsupported type: %s", cmd)
	}
}
