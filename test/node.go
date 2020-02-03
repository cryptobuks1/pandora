package test

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

type node struct {
	cmd *exec.Cmd
	pid int
}

func newNode() *node {
	m := &node{}
	m.cmd = m.prepareCommand()
	return m
}

func (*node) prepareArgs() []string {
	args := []string{"-test.run", "^TestNodeMain$", "-test.coverprofile", "./test/coverage.out"}
	return args
}

func (n *node) prepareCommand() *exec.Cmd {
	cmd := exec.Command("./test/node.test", n.prepareArgs()...)
	wd, _ := os.Getwd()
	abs, _ := filepath.Abs(wd)
	cmd.Dir, _ = filepath.Split(abs)
	return cmd
}

func (n *node) start() error {
	if err := execCmd("node", n.cmd); err != nil {
		return err
	}
	n.pid = n.cmd.Process.Pid
	return nil
}

func (n *node) stop() error {
	if n == nil || n.cmd == nil || n.cmd.Process == nil {
		return nil
	}
	if err := n.cmd.Process.Signal(syscall.SIGQUIT); err != nil {
		return err
	}
	return nil
}
