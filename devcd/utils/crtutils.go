package utils

import (
	"bytes"
	"devcd/logger"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func CheckContainerRtStatus(crt IContainerRuntime) error {
	return crt.CheckContainerRtStatus()
}

func PullContainerImage(crt IContainerRuntime, imageName string) error {
	return crt.PullImage(imageName)
}

func BuildContainerImage(crt IContainerRuntime, imageName string, imgFilePath string) error {
	return crt.BuildImage(imageName, imgFilePath)
}

func RemoveContainerImage(crt IContainerRuntime, imageName string) error {
	return crt.RemoveImage(imageName)
}

func RunDockerComposeUp(crt IContainerRuntime, containerName string, composeFile string) error {
	return crt.RunDockerComposeUp(containerName, composeFile)
}

func RunDockerComposeDown(crt IContainerRuntime, containerName string, composeFile string) error {
	return crt.RunDockerComposeDown(containerName, composeFile)
}

func IsContainerRunning(crt IContainerRuntime, containerName string) bool {
	return crt.IsContainerRunning(containerName)
}

func StopContainer(crt IContainerRuntime, containerName string) error {
	return crt.StopContainer(containerName)
}

func RemoveContainer(crt IContainerRuntime, containerName string) error {
	return crt.RemoveContainer(containerName)
}

func CreateVolume(crt IContainerRuntime, volumeName string) error {
	return crt.CreateVolume(volumeName)
}

func DeleteVolume(crt IContainerRuntime, volumeName string) error {
	return crt.DeleteVolume(volumeName)
}

func CreateNetwork(crt IContainerRuntime, networkName string) error {
	return crt.CreateNetwork(networkName)
}

func DeleteNetwork(crt IContainerRuntime, networkName string) error {
	return crt.DeleteNetwork(networkName)
}

func StartStopCompose(crt IContainerRuntime, containerName, composeFile, action string) {

	if action == "start" {
		RunDockerComposeUp(crt, containerName, composeFile)
		//time.Sleep(sleepTime * time.Second)
	} else if action == "stop" {
		RunDockerComposeDown(crt, containerName, composeFile)
		//time.Sleep(sleepTime * time.Second)
	} else {
		logger.Info("Invalid Compose run option. Choose [start/stop] ", "action", action)
	}
}

func BuildDockerImageWithArg(containerRtEngine string, dockerImage string, dockerFilePath string, buildArg string) {

	// Build the Container Image
	var stdout, stderr bytes.Buffer

	logger.Info("Building Container Runtime image...", "dockerImage", dockerImage)
	cmd := exec.Command(containerRtEngine, "build", "--build-arg", buildArg, "-t", dockerImage, dockerFilePath)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		logger.Error("Error executing command", "error", err)
		logger.Error("Standard error", "out", stderr.String())
		syscall.Exit(1)
	}
	logger.Info("Standard output", "out", stdout.String())
	logger.Info("Container Runtime build ran successfully", "dockerImage", dockerImage)
}

func ExtractMSFromDocker(containerRtEngine, msImg, msExePath string) {

	// Pull and extract
	cmdPull := exec.Command(containerRtEngine, "pull", msImg)
	cmdPull.Stdout = os.Stdout
	cmdPull.Stderr = os.Stderr
	if err := cmdPull.Run(); err != nil {
		logger.Error("Error pulling image:", "error", err)
		return
	}

	containerIDCmd := exec.Command(containerRtEngine, "create", msImg, "bin/bash")
	containerIDOutput, err := containerIDCmd.Output()
	if err != nil {
		logger.Error("Error creating container:", "error", err)
		return
	}
	containerID := strings.TrimSpace(string(containerIDOutput))
	logger.Debug(fmt.Sprintf("containerID: %s", containerID))

	cmdCopy := exec.Command(containerRtEngine, "cp", fmt.Sprintf("%s:%s", containerID, msExePath), ".")
	cmdCopy.Stdout = os.Stdout
	cmdCopy.Stderr = os.Stderr
	if err := cmdCopy.Run(); err != nil {
		logger.Error("Executing cmd", "cmd", cmdCopy.String())
		logger.Error("Error copying file from container:", "error", err)
		return
	}

	// Clean up
	cmdRemove := exec.Command(containerRtEngine, "rm", containerID)
	cmdRemove.Stdout = os.Stdout
	cmdRemove.Stderr = os.Stderr
	if err := cmdRemove.Run(); err != nil {
		logger.Error("Error removing container:", "error", err)
		return
	}

	cmdRemoveImage := exec.Command(containerRtEngine, "rmi", msImg)
	cmdRemoveImage.Stdout = os.Stdout
	cmdRemoveImage.Stderr = os.Stderr
	if err := cmdRemoveImage.Run(); err != nil {
		logger.Error("Error removing image:", "error", err)
		return
	}
}

func ExtractMSFromDockerToPath(containerRtEngine, msImg, msExePath, destPath string) {

	// Pull and extract
	cmdPull := exec.Command(containerRtEngine, "pull", msImg)
	cmdPull.Stdout = os.Stdout
	cmdPull.Stderr = os.Stderr
	if err := cmdPull.Run(); err != nil {
		logger.Error("Error pulling image:", "error", err)
		return
	}

	containerIDCmd := exec.Command(containerRtEngine, "create", msImg, "bin/bash")
	containerIDOutput, err := containerIDCmd.Output()
	if err != nil {
		logger.Error("Error creating container:", "error", err)
		return
	}
	containerID := strings.TrimSpace(string(containerIDOutput))
	logger.Debug(fmt.Sprintf("containerID: %s", containerID))

	cmdCopy := exec.Command(containerRtEngine, "cp", fmt.Sprintf("%s:%s", containerID, msExePath), destPath)
	cmdCopy.Stdout = os.Stdout
	cmdCopy.Stderr = os.Stderr
	if err := cmdCopy.Run(); err != nil {
		logger.Error("Executing cmd", "cmd", cmdCopy.String())
		logger.Error("Error copying file from container:", "error", err)
		return
	}

	// Clean up
	cmdRemove := exec.Command(containerRtEngine, "rm", containerID)
	cmdRemove.Stdout = os.Stdout
	cmdRemove.Stderr = os.Stderr
	if err := cmdRemove.Run(); err != nil {
		logger.Error("Error removing container:", "error", err)
		return
	}

	cmdRemoveImage := exec.Command(containerRtEngine, "rmi", msImg)
	cmdRemoveImage.Stdout = os.Stdout
	cmdRemoveImage.Stderr = os.Stderr
	if err := cmdRemoveImage.Run(); err != nil {
		logger.Error("Error removing image:", "error", err)
		return
	}
}

func ExtractMSFromDockerLocal(containerRtEngine, msImg, msExePath string) {

	containerIDCmd := exec.Command(containerRtEngine, "create", msImg, "bin/bash")
	containerIDOutput, err := containerIDCmd.Output()
	if err != nil {
		logger.Error("Error creating container:", "error", err)
		return
	}
	containerID := strings.TrimSpace(string(containerIDOutput))
	logger.Debug(fmt.Sprintf("containerID: %s", containerID))

	cmdCopy := exec.Command(containerRtEngine, "cp", fmt.Sprintf("%s:%s", containerID, msExePath), ".")
	cmdCopy.Stdout = os.Stdout
	cmdCopy.Stderr = os.Stderr
	if err := cmdCopy.Run(); err != nil {
		logger.Error("Executing cmd", "cmd", cmdCopy.String())
		logger.Error("Error copying file from container", "error", err)
		return
	}

	// Clean up
	cmdRemove := exec.Command(containerRtEngine, "rm", containerID)
	cmdRemove.Stdout = os.Stdout
	cmdRemove.Stderr = os.Stderr
	if err := cmdRemove.Run(); err != nil {
		logger.Error("Error removing container", "error", err)
		return
	}
}
