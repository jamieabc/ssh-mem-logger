package remoteCmds

import "fmt"

type top struct {
	process string
	order   string
	user    string
}

func (t top) String() string {
	cmd := "top"
	if t.order != "" {
		cmd = fmt.Sprintf("%s -o %s", cmd, t.order)
	}

	if t.user != "" {
		cmd = fmt.Sprintf("%s -U %s", cmd, t.user)
	}

	return fmt.Sprintf("%s | grep -B 1 %s", cmd, t.process)
}

func newTop(args []interface{}) (RemoteCmds, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("missing process name")
	}

	order := ""
	if len(args) >= 2 {
		order = args[1].(string)
	}

	user := ""
	if len(args) >= 3 {
		user = args[2].(string)
	}

	return &top{
		process: args[0].(string),
		order:   order,
		user:    user,
	}, nil
}
