package utils

import (
	"devcd/logger"
)

func InstallHelmChartWithFiles(chartName string, chartPath string, files []string) error {

	logger.Info("Installing helm chart", "chartName", chartName, "chartPath", chartPath)
	logger.Info("helm chart with external files", "files", files)

	cwd, err := GetCwd()
	if err != nil {
		logger.Error("Error getting CWD")
		return err
	}

	logger.Debug("Current working dir", "CWD", cwd)

	customCommand := []string{"helm", "install", chartName, chartPath}

	if len(files) >= 1 {
		for _, file := range files {
			customCommand = append(customCommand, "-f", file)
		}
	}

	customCommand = append(customCommand, "--set", "CWD="+cwd)

	err = RunC(customCommand)
	if err != nil {
		logger.Error("Error installing chart", "chartName", chartName, "error", err)
		return err
	}
	logger.Info("Installed helm chart successfully", "chartName", chartName)
	return nil

}

func DeleteHelmChart(chartName string) error {

	logger.Info("Deleting helm chart", "chartName", chartName)
	customCommand := []string{"helm", "delete", chartName}

	err := RunC(customCommand)
	if err != nil {
		logger.Error("Error deleting helm chart", "chartName", chartName, "error", err)
		return err
	}
	logger.Info("Deleted helm chart successfully", "chartName", chartName)
	return nil
}

func InstallDeleteHelmChart(chartName, chartPath, action string, files []string) error {

	if action == "start" {
		return InstallHelmChartWithFiles(chartName, chartPath, files)
		//time.Sleep(sleepTime * time.Second)
	} else if action == "stop" {
		return DeleteHelmChart(chartName)
		//time.Sleep(sleepTime * time.Second)
	} else {
		logger.Warn("Invalid Chart run option. Choose [start/stop] action", "input action", action)
	}
	return nil
}

func InstallHelmChart(chartName string, chartPath string) {

	logger.Info("Installing helm chart", "chartName", chartName, "chartPath", chartPath)
	args := []string{"helm", "install", chartName, chartPath}
	err := RunC(args)
	if err != nil {
		logger.Error("Error installing chart", "chartName", chartName, "error", err)
	}
	logger.Info("Installed helm chart successfully", "chartName", chartName)
}
