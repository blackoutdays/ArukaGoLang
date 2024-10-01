package test

import (
	"fmt"
	"os/exec"
)

// RunCommand выполняет одну команду.
func RunCommand(cmd string) error {
	fmt.Printf("Running command: %s\n", cmd)
	command := exec.Command("sh", "-c", cmd)
	output, err := command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("command failed: %w, output: %s", err, string(output))
	}
	fmt.Printf("Output: %s\n", string(output))
	return nil
}

// RunCommands выполняет последовательность команд.
func RunCommands(cmds []string) error {
	for _, cmd := range cmds {
		if err := RunCommand(cmd); err != nil {
			return fmt.Errorf("Error running %s: %v", cmd, err)
		}
	}
	return nil
}
