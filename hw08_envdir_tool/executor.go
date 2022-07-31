package main

import (
	"errors"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for varName, varParams := range env {
		if varParams.NeedRemove {
			os.Unsetenv(varName)
		} else {
			os.Setenv(varName, varParams.Value)
		}
	}

	cm := cmd[0]
	command := exec.Command(cm, cmd[1:]...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			returnCode = exitError.ExitCode()
		}
	}

	return
}
