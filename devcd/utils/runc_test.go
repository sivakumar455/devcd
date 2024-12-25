package utils

import (
	"devcd/logger"
	"testing"
)

func TestRunCsuccess(t *testing.T) {
	logger.Info("Running Test", "test", "TestRunCsuccess")
	args := []string{"ls", "-lh"}
	err := RunC(args)
	if err != nil {
		t.Errorf("err: %v", err)
	}

}
func TestRunCfail(t *testing.T) {
	logger.Info("Running Test", "test", "TestRunCfail")
	args := []string{"lst", "-lh"}
	err := RunC(args)
	if err == nil {
		t.Error("Expected an error running test TestRunCfail, but got nil")
	}

}
