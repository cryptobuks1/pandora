package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"pandora/test/utils"
)

type Node struct {
	cmd *exec.Cmd
	pid int
}

func NewNode() *Node {
	m := &Node{}
	m.cmd = m.prepareCommand()
	return m
}

func (*Node) prepareArgs() []string {
	args := []string{"-test.run", "^TestNodeMain$", "-test.coverprofile", "./node/coverage.out"}
	return args
}

func (n *Node) prepareCommand() *exec.Cmd {
	cmd := exec.Command("./node/node.test", n.prepareArgs()...)
	wd, _ := os.Getwd()
	abs, _ := filepath.Abs(wd)
	cmd.Dir, _ = filepath.Split(abs)
	return cmd
}

func (n *Node) Start() error {
	if err := utils.ExecCmd("node", n.cmd); err != nil {
		return err
	}
	n.pid = n.cmd.Process.Pid
	return nil
}

func (n *Node) Stop() error {
	if n == nil || n.cmd == nil || n.cmd.Process == nil {
		return nil
	}
	if err := n.cmd.Process.Signal(syscall.SIGQUIT); err != nil {
		return err
	}
	return nil
}
