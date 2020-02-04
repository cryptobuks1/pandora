package utils

import (
	"bufio"
	"fmt"
	"os/exec"
)

func ExecCmd(logPrefix string, cmd *exec.Cmd) error {
	stdoutReader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderrReader, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	stdoutScanner := bufio.NewScanner(stdoutReader)
	go func() {
		for stdoutScanner.Scan() {
			fmt.Printf("%s | stdout ::> %s\n", logPrefix, stdoutScanner.Text())
		}
	}()

	stderrScanner := bufio.NewScanner(stderrReader)
	go func() {
		for stderrScanner.Scan() {
			fmt.Printf("%s | stderr ::> %s\n", logPrefix, stderrScanner.Text())
		}
	}()

	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}
