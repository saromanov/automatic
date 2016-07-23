package automatic

import (
	"fmt"
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

// ExecScript provides execution of the script
func ExecScript(path string) error {
	cmd := exec.Command(path)
	out, err := cmd.Output()
	if err != nil {
		return  err
	}

	fmt.Printf("%s: %s", path, string(out))

	return nil
}
