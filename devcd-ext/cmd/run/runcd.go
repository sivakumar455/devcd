package run

import (
	"devcd/cmd/root"
	"devcd/logger"
	"devcd_ext/cmd/install"
)

func RunDevcd() {
	logger.Debug("Devcd external")
	root.Execute()
}
func init() {
	// Adding devcd extenions here
	root.RootCmd.AddCommand(install.InstallCmd)
}
