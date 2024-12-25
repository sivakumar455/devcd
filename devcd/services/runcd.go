package services

import (
	"devcd/config"
	"devcd/logger"
	"devcd/utils"
	"fmt"
	"strings"
	"time"
)

type StartStopSvcFunc func(string)

var featureCd string
var runMode string

var bsComposeHome string
var bsHelmHome string
var msComposeHome string
var msHelmHome string

func setup() {
	featureCd = config.FEATURE_CD
	runMode = config.RUN_MODE

	bsComposeHome = config.BS_COMPOSE_HOME
	bsHelmHome = config.BS_HELM_HOME
	msComposeHome = config.MS_COMPOSE_HOME
	msHelmHome = config.MS_HELM_HOME
}

func RunService(service, arg string) {
	logger.Info("Runnig for Service", "service", service, "action", arg)
	setup()
	switch service {
	case "cd":
		logger.Info("Runnig full CD")
		startStopSvc(StartStopCd, arg)
	case "bs":
		logger.Info("Runnig backing Services")
		startStopSvc(StartStopBs, arg)
	case "ms":
		logger.Info("Runnig microservices")
		startStopSvc(StartStopMs, arg)
	case "testbs":
		logger.Info("Runnig testbs for testing selected backing services")
		startStopSvc(StartStopTestBs, arg)
	case "testms":
		logger.Info("Runnig testms for testing selected microservices")
		startStopSvc(startStopTestMs, arg)
	}
}

func startStopSvc(fn StartStopSvcFunc, arg string) {
	logger.Info(fmt.Sprintf("Run CD Argument : %s", arg))

	if arg == "stop" || arg == "start" {
		fn(arg)
	} else {
		logger.Info("Not a valid argument for run arg action. Supported options are [start, stop]")
	}
}

func StartStopCd(action string) {
	logger.Info("Starting All")
	StartStopBs(action)
	time.Sleep(10 * time.Second)
	StartStopMs(action)
}

func StartStopBs(action string) {

	if runMode == "compose" {
		logger.Info("ComposeBS List", "ComposeBsLst", config.ComposeCfg.ComposeBs)
		StartStopComposeList(action, bsComposeHome, config.ComposeCfg.ComposeBs)
	}
	if runMode == "helm" {
		var globalFiles []string
		for _, globalFile := range config.HelmCfg.ValuesBs {
			globalFilePath := fmt.Sprintf("%s%s", bsHelmHome, globalFile)
			globalFiles = append(globalFiles, globalFilePath)
		}
		InstallDeleteHelmChartList(action, bsHelmHome, config.HelmCfg.HelmBs, globalFiles)
	}
}

func StartStopTestBs(action string) {
	if runMode == "compose" {
		logger.Info("ComposeBS List", "config.ComposeCfg.ComposeTestBs", config.ComposeCfg.ComposeTestBs)
		StartStopComposeList(action, bsComposeHome, config.ComposeCfg.ComposeTestBs)
	}
	if config.RUN_MODE == "helm" {
		var globalFiles []string
		for _, globalFile := range config.HelmCfg.ValuesBs {
			globalFilePath := fmt.Sprintf("%s%s", bsHelmHome, globalFile)
			globalFiles = append(globalFiles, globalFilePath)
		}
		InstallDeleteHelmChartList(action, bsHelmHome, config.HelmCfg.HelmTestBs, globalFiles)
	}
}

func StartStopComposeList(action string, composeHomePath string, composeLst []string) {
	logger.Debug("ComposeBS List", "composeLst", composeLst)
	for _, composeFilePath := range composeLst {
		fmt.Println()
		cName := strings.Split(composeFilePath, "/")[1]
		logger.Info(fmt.Sprintf("Deplying compose: %v", cName))
		logger.Info("Running Docker Compose for BS under folder", "cName", cName, "composeFilePath", composeFilePath)
		containerName := fmt.Sprintf("%s-%s", cName, featureCd)
		composeFile := fmt.Sprintf("%s%s", composeHomePath, composeFilePath)
		utils.StartStopCompose(config.ContainerEngine, containerName, composeFile, action)
		fmt.Println()
	}
}

func StartStopMs(action string) {
	logger.Info("MS action ", "action", action)

	if runMode == "compose" {
		StartStopComposeList(action, msComposeHome, config.ComposeCfg.ComposeMs)
	}
	if runMode == "helm" {
		var globalFiles []string
		for _, globalFile := range config.HelmCfg.ValuesMs {
			globalFilePath := fmt.Sprintf("%s%s", msHelmHome, globalFile)
			globalFiles = append(globalFiles, globalFilePath)
		}
		InstallDeleteHelmChartList(action, msHelmHome, config.HelmCfg.HelmMs, globalFiles)
	}
}

func startStopTestMs(action string) {

	logger.Info("MS action ", "action", action)
	if runMode == "compose" {
		StartStopComposeList(action, msComposeHome, config.ComposeCfg.ComposeTestMs)
	}
	if runMode == "helm" {
		var globalFiles []string
		for _, globalFile := range config.HelmCfg.ValuesMs {
			globalFilePath := fmt.Sprintf("%s%s", msHelmHome, globalFile)
			globalFiles = append(globalFiles, globalFilePath)
		}
		InstallDeleteHelmChartList(action, msHelmHome, config.HelmCfg.HelmTestMs, globalFiles)
	}
}

func InstallDeleteHelmChartList(action string, helmHomePath string, helmLst []string, globalValues []string) {
	fmt.Println()
	for _, helmFilePath := range helmLst {
		fmt.Println()
		fmt.Println(strings.Repeat("#", 50))
		bsName := strings.Split(helmFilePath, "/")[1]
		logger.Info(fmt.Sprintf("Deploying chart: %v", bsName))
		logger.Info("Deploying Helm Chart for BS under folder", "bsName", bsName, "helmFilePath", helmFilePath)
		composeFile := fmt.Sprintf("%s%s", helmHomePath, helmFilePath)
		err := utils.InstallDeleteHelmChart(bsName, composeFile, action, globalValues)
		if err != nil {
			logger.Error("Error deploying helm chart ", "chart", bsName, "error", err)
		}
		fmt.Println(strings.Repeat("#", 50))
		fmt.Println()
	}
}
