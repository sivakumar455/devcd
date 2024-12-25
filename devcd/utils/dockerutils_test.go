package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestContainerEgineStatus(t *testing.T) {
	fmt.Println("Testing container status")

	crt := ContainerRuntime{"docker"}

	err := crt.CheckContainerRtStatus()

	if err != nil {
		t.Error("Container is not running")
	}
}

func TestPullDockerImage(t *testing.T) {
	fmt.Println("Testing container status")

	crt := &ContainerRuntime{"docker"}

	err := PullContainerImage(crt, "alpine:latest")

	if err != nil {
		t.Error("Error while pulling image")
	}
}

func TestBuildDockerImage(t *testing.T) {

	crt := &ContainerRuntime{"docker"}

	// create tmp dir and docker file
	dirName := "testfiles"
	fileName := "Dockerfile"
	dockerfileContent := []byte("FROM alpine\nRUN echo 'test imagess'")
	err := createTempFile(dirName, fileName, dockerfileContent)
	if err != nil {
		t.Errorf("err creating temp file: %v", err)
	}

	// Build docker img
	imageName := "test-img"
	//dockerFilePath := filepath.Join(dirName, fileName)
	err = BuildContainerImage(crt, imageName, dirName)
	if err != nil {
		t.Error("file error")
	}

	// Validate docker img
	cmd := exec.Command("docker", "images", "-q", imageName)
	output, err := cmd.Output()
	if err != nil || len(output) == 0 {
		t.Errorf("Docker image was not built successfully")
	}

	// Step 3: Cleanup
	err = RemoveContainerImage(crt, imageName)
	if err != nil {
		t.Error("err removing image")
	}
	// remove tmp dir and img
	defer os.RemoveAll(dirName)

}

func TestCreateDockerFile(t *testing.T) {
	dockerfileContent := []byte("FROM alpine\nRUN echo 'test images'")
	err := createTempFile("testfiles", "Dockerfile", dockerfileContent)
	if err != nil {
		t.Errorf("err creating temp file: %v", err)
	}

	dfilePath := filepath.Join("testfiles", "Dockerfile")

	if _, err := os.Stat(dfilePath); os.IsNotExist(err) {
		t.Errorf("Dockerfile was not created: %v ", err)
	} else {
		fmt.Println("Temp file Successfully created")
		err = os.RemoveAll("testfiles")
		if err != nil {
			t.Errorf("err removing temp dir: %v", err)
		}
		fmt.Println("Temp file Successfully removed")

	}
}

func createTempFile(dirName, fileName string, fileContent []byte) error {

	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return fmt.Errorf("err creating temp dir: %v", err)
	}

	file := filepath.Join(dirName, fileName)

	fp, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("err creating temp file: %v", err)
	}
	defer fp.Close()

	_, err = fp.Write(fileContent)
	if err != nil {
		return fmt.Errorf("err writing to temp file: %v", err)
	}

	//time.Sleep(20 * time.Second)
	//defer os.RemoveAll(dirName)
	return nil
}
