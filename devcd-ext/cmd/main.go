package main

import (
	"devcd/logger"
	"devcd_ext/cmd/run"
)

// devcd-extensions to add any additional behaviour to devcd like adding installations or cleanup
func main() {
	logger.Debug("Devcd externals Main")
	run.RunDevcd()
}
