package services

import (
	"devcd/config"
	"devcd/logger"
	"devcd/utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var msTestVer = os.Getenv("MS1_VER")
var TMP_DIR string
var RT_DIR string

type DefaultExtractms struct{}

func (ms DefaultExtractms) Extract() {
	tmpDir := os.Getenv("TMP_NXS_REPO")
	tmpFolder := "ms_img"
	err := utils.CreateTmpDirWithTS(tmpDir, tmpFolder)
	if err != nil {
		logger.Error("Err creating temp dir", "tmpDir", tmpDir)
	}

	rootDir, _ := utils.GetCwd()
	RT_DIR = filepath.Join(rootDir, "devcd-runtime", "runtime")

	extractAllMs()
	copyToRuntime()

}

func extractAllMs() {
	ms1()
}

func ms1() {
	msName := "microservice-one"
	msVersion := os.Getenv("MS1_VER")
	extractMsImg(msName, msVersion)
}

func extractMsImg(msName, msVersion string) {
	msService := fmt.Sprintf("%s-service", msName)

	msExeJar := fmt.Sprintf("%s-%s.jar", msName, msVersion)
	msExePath := fmt.Sprintf("/deploy/%s", msExeJar)
	msImg := fmt.Sprintf("%s/%s:%s", "MS_REPO", msService, msVersion)
	if !chkRuntimeJarExists(msExeJar) {
		utils.ExtractMSFromDocker(config.CONTAINER_RTE, msImg, msExePath, RT_DIR)
	}
}

func chkRuntimeJarExists(runtimeJar string) bool {
	currentDirectory, _ := os.Getwd()
	twoDirsBack := filepath.Dir(filepath.Dir(currentDirectory))
	destDir := filepath.Join(twoDirsBack, "runtime", runtimeJar)

	if _, err := os.Stat(destDir); err == nil {
		logger.Info("runtimeJar exists", "runtimeJar", runtimeJar)
		return true
	} else {
		logger.Warn("runtimeJar does not exist", "runtimeJar", runtimeJar)
		return false
	}
}

func copyToRuntime() {
	fmt.Println("Copying all jars to runtime")
	currentDirectory, _ := os.Getwd()
	twoDirsBack := filepath.Dir(filepath.Dir(currentDirectory))
	destDir := filepath.Join(twoDirsBack, "runtime")

	files, _ := filepath.Glob(filepath.Join(currentDirectory, "*"))

	for _, file := range files {
		source := file
		destination := filepath.Join(destDir, filepath.Base(file))

		if strings.HasSuffix(file, ".jar") {
			if err := utils.HandleDestination(source, destination); err != nil {
				logger.Error("Error:", "error", err)
				return
			}
			msgCopy := fmt.Sprintf("Copied %s to %s", filepath.Base(file), filepath.Base(destDir))
			logger.Info(msgCopy)
			logger.Debug("File copied successfully!")
		} else if strings.HasSuffix(file, "stubs") {
			logger.Info("Copying Wiremock stubs")
			destination := "REPLACE_PATH"
			if _, err := os.Stat(destination); err == nil {
				msg2 := fmt.Sprintf("Directory with name %s exists in %s, not copying", filepath.Base(destination), filepath.Base(destDir))
				logger.Info(msg2)
			} else {
				os.MkdirAll(destination, os.ModePerm)
				utils.CopyDirectory(source, destination)
				logger.Info(fmt.Sprintf("Copied %s to %s", filepath.Base(file), filepath.Base(destination)))
			}
		} else if strings.HasSuffix(file, "filters") {
			logger.Info("Copying other jars")
			logger.Debug("File", file)
			filterDir, _ := filepath.Glob(filepath.Join(file, "*"))

			for _, fltr := range filterDir {

				destination := filepath.Join(destDir, fmt.Sprintf("filters-%s-%s", filepath.Base(fltr), msTestVer))
				if _, err := os.Stat(destination); err == nil {
					logger.Info(fmt.Sprintf("Directory with name %s exists in %s, not copying", filepath.Base(destination), filepath.Base(destDir)))
				} else {
					os.MkdirAll(destination, os.ModePerm)
					utils.CopyDirectory(fltr, destination)
					logger.Info(fmt.Sprintf("Copied %s to %s", filepath.Base(file), filepath.Base(destination)))
				}
			}

		} else {
			logger.Warn(fmt.Sprintf("Ignoring File %s in %s", filepath.Base(file), destDir))
		}
	}
}
