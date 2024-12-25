package utils

import (
	"bytes"
	"devcd/logger"
	"os/exec"
)

// Runs unix commands which recives as slice of strings and first arg would be the unix cmd
func RunC(args []string) error {
	logger.Debug("RunC Command Args", "args", args)
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		logger.Error("Failed to execute RunC command", "output", stderr.String())
		logger.Error("Failed to execute RunC command", "error", err)
		return err
	}

	logger.Debug("RunC Command executed successfully", "ouput", stdout.String())
	return nil
}
