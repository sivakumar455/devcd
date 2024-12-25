package utils

import (
	"devcd/logger"
	"fmt"
	"os/exec"
	"time"
)

type IContainerRuntime interface {
	CheckContainerRtStatus() error
	PullImage(imageName string) error
	BuildImage(imageName string, imgFilePath string) error
	RemoveImage(imageName string) error

	CreateContainer(containerName string) error
	StartContainer(containerName string) error
	StopContainer(containerName string) error
	RemoveContainer(containerName string) error
	IsContainerRunning(containerName string) bool

	RunDockerComposeUp(containerName string, composeFile string) error
	RunDockerComposeDown(containerName string, composeFile string) error

	CreateVolume(volumeName string) error
	DeleteVolume(volumeName string) error

	CreateNetwork(networkName string) error
	DeleteNetwork(networkName string) error
}

type ContainerRuntime struct {
	CrtEngine string
}

func (c *ContainerRuntime) CheckContainerRtStatus() error {
	containerRtEngine := c.CrtEngine
	logger.Info("Checking Container Runtime Engine status", "containerRtEngine", containerRtEngine)

	args := []string{containerRtEngine, "info"}
	err := RunC(args)
	if err != nil {
		logger.Error("Err Checking status of Container Runtime Engine", "containerRtEngine", containerRtEngine)
		logger.Error("Please start the Container Runtime Engine", "error", err)
		return fmt.Errorf("%v is not running", containerRtEngine)
	}

	logger.Info("Container Runtime Engine is running", "containerRtEngine", containerRtEngine)
	return nil
}

func (c *ContainerRuntime) PullImage(imageName string) error {

	// Pull the Container Runtime image
	containerRtEngine := c.CrtEngine
	logger.Info("Pulling Container Runtime image...", "imageName", imageName)
	args := []string{containerRtEngine, "pull", imageName}
	err := RunC(args)
	if err != nil {
		logger.Error("Error pulling Container Runtime image...", "imageName", imageName, "error", err)
		return fmt.Errorf("err pulling image : %v", imageName)
	}

	logger.Info("Container Runtime image pulled successfully", "imageName", imageName)
	return nil
}

func (c *ContainerRuntime) BuildImage(imgName string, imgFilePath string) error {

	// Build the Container Image
	containerRtEngine := c.CrtEngine
	logger.Info("Building Container Runtime image...", "imgName", imgName)

	args := []string{containerRtEngine, "build", "-t", imgName, imgFilePath}
	err := RunC(args)
	if err != nil {
		logger.Error("Error Building Container Runtime image...", "imgName", imgName, "error", err)
		return fmt.Errorf("err Building image : %v", err)
	}

	logger.Info("Build Container Runtime image successfully.", "imgName", imgName)
	return nil
}

func (c *ContainerRuntime) RemoveImage(imageName string) error {
	containerRtEngine := c.CrtEngine
	// Clean up
	logger.Info("Removing Container Runtime image...", "imageName", imageName)
	args := []string{containerRtEngine, "rmi", imageName}
	err := RunC(args)
	if err != nil {
		logger.Info("Error removing Container Runtime image...", "imageName", imageName, "error", err)
		return fmt.Errorf("err removing image : %v", err)
	}
	logger.Info("Removed Container Runtime image successfully", "imageName", imageName)
	return nil
}

func (c *ContainerRuntime) CreateContainer(containerName string) error {
	panic("CreateContainer not implemented yet")
}

func (c *ContainerRuntime) StartContainer(containerName string) error {
	panic("StartContainer not implemented yet")
}

func (c *ContainerRuntime) StopContainer(containerName string) error {
	containerRtEngine := c.CrtEngine
	logger.Info("Stopping Container", "containerName", containerName)

	args := []string{containerRtEngine, "stop", containerName}
	err := RunC(args)
	if err != nil {
		logger.Error("Error stopping Container", "containerName", containerName, "error", err)
		return fmt.Errorf("err stopping container : %v", err)
	}

	return nil
}

func (c *ContainerRuntime) RemoveContainer(containerName string) error {
	containerRtEngine := c.CrtEngine
	logger.Info("Removing Container", "containerName", containerName)

	args := []string{containerRtEngine, "rm", containerName}
	err := RunC(args)
	if err != nil {
		logger.Error("Error removing Container", "containerName", containerName, "error", err)
		return fmt.Errorf("err removing container : %v", err)
	}

	return nil
}

func (c *ContainerRuntime) IsContainerRunning(containerName string) bool {
	containerRtEngine := c.CrtEngine
	for i := 0; i < 5; i++ { // Retry up to 5 times
		cmd := exec.Command(containerRtEngine, "ps", "-q", "-f", fmt.Sprintf("name=%s", containerName))
		output, err := cmd.Output()
		if err != nil || len(output) == 0 {
			logger.Info("Container is not running, retrying...")
			time.Sleep(10 * time.Second) // Wait for 10 seconds before retrying
			continue
		}
		logger.Info("Container is running.")
		return true
	}
	return false
}

func (c *ContainerRuntime) RunDockerComposeUp(containerName string, composeFile string) error {

	containerRtEngine := c.CrtEngine

	logger.Info("Creating Container", "containerName", containerName)
	args := []string{containerRtEngine, "compose", "-p", containerName, "-f", composeFile, "up", "-d"}
	err := RunC(args)
	if err != nil {
		logger.Error("Error creating Container", "containerName", containerName, "error", err)
		return fmt.Errorf("err creating container : %v", containerName)
	}

	logger.Info("Created Container successfully", "containerName", containerName)
	return nil
}

func (c *ContainerRuntime) RunDockerComposeDown(containerName string, composeFile string) error {

	containerRtEngine := c.CrtEngine

	logger.Info("Terminating Container", "containerName", containerName)
	args := []string{containerRtEngine, "compose", "-p", containerName, "-f", composeFile, "down"}
	err := RunC(args)
	if err != nil {
		logger.Error("Error Terminating Container", "containerName", containerName, "error", err)
		return fmt.Errorf("err terminating container : %s", containerName)
	}

	logger.Info("Terminated Container successfully", "containerName", containerName)
	return nil
}

func (c *ContainerRuntime) CreateVolume(volumeName string) error {
	containerRtEngine := c.CrtEngine

	logger.Info("Creating Docker volume", "volumeName", volumeName)
	args := []string{containerRtEngine, "volume", "create", volumeName}
	err := RunC(args)
	if err != nil {
		logger.Error("Error Creating Docker volume", "volumeName", volumeName, "error", err)
		return err
	}

	logger.Info("Created Docker volume successfully", "volumeName", volumeName)
	return nil
}

func (c *ContainerRuntime) DeleteVolume(volumeName string) error {
	containerRtEngine := c.CrtEngine
	logger.Info("Deleting Docker volume", "volumeName", volumeName)

	args := []string{containerRtEngine, "volume", "rm", volumeName}
	err := RunC(args)
	if err != nil {
		logger.Error("Error deleting Docker volume", "volumeName", volumeName, "error", err)
		return err
	}

	logger.Info("Deleted Docker volume successfully", "volumeName", volumeName)
	return nil
}

func (c *ContainerRuntime) CreateNetwork(networkName string) error {

	containerRtEngine := c.CrtEngine

	logger.Info("Creating Docker network", "networkName", networkName)
	args := []string{containerRtEngine, "network", "create", networkName}
	err := RunC(args)
	if err != nil {
		logger.Info("Error creating Docker network", "networkName", networkName, "error", err)
		return err
	}

	logger.Info("Created Docker network successfully", "networkName", networkName)
	return nil
}

func (c *ContainerRuntime) DeleteNetwork(networkName string) error {

	containerRtEngine := c.CrtEngine
	logger.Info("Deleting Docker network", "networkName", networkName)
	args := []string{containerRtEngine, "network", "rm", networkName}
	err := RunC(args)
	if err != nil {
		logger.Error("Error deleting Docker network", "networkName", networkName, "error", err)
		return err
	}

	logger.Info("Error deleting Docker network", "networkName", networkName)
	return nil
}
