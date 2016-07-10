package automatic

import (
	"os/exec"
)

// execute.go provides execution of the scripts ot commands

// ExecCommand provides execution of the command
func ExecCommand(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
