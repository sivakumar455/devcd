package services

import (
	"devcd/config"
	"devcd/logger"
	"devcd/utils"
	"path/filepath"
)

func CleanFullCd() {
	CleanLogs()
	CleanNexusJars()
	CleanRuntimeJars()
}

// var runtimeHome = os.Getenv("DEVCD_RUNTIME")
var logsHome = "/logs"
var nexusHome = "/nexus-repo"
var rtExeHome = "/runtime"

func CleanLogs() {
	logDir := filepath.Join(config.DEVCD_RUNTIME, logsHome)

	err := utils.RemoveDirectoryExcludeGitIgnore(logDir)
	if err != nil {
		logger.Error("Error removing log directory", "logDir", logDir, "error", err)
	}
	logger.Info("Successfully cleaned cd logs", "logDir", logDir)
}

func CleanNexusJars() {

	nxsDir := filepath.Join(config.DEVCD_RUNTIME, nexusHome)

	err := utils.RemoveDirectoryExcludeGitIgnore(nxsDir)
	if err != nil {
		logger.Error("Error removing nexus jar directory", "nxsDir", nxsDir, "error", err)
	}
	logger.Info("Successfully cleaned cd nexus jars", "nxsDir", nxsDir)
}

func CleanRuntimeJars() {

	rtjarDir := filepath.Join(config.DEVCD_RUNTIME, rtExeHome)

	err := utils.RemoveDirectoryExcludeGitIgnore(rtjarDir)
	if err != nil {
		logger.Error("Error removing runtime jar directory", "rtjarDir", rtjarDir, "error", err)
	}
	logger.Info("Successfully cleaned cd runtime jars", "rtjarDir", rtjarDir)
}
