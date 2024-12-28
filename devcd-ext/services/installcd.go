package services

import (
	"devcd/config"
	"devcd/logger"
	"devcd/utils"
	"devcd_ext/couchbase"
	"fmt"
	"os"
)

func InstallJavaImage() {
	CheckContainerRtEngine()
	BuildJavaImage()
}

func CheckContainerRtEngine() {

	// Check if Container Runtime is installed
	// cmd := exec.Command("command", "-v", containerRtEngine)
	// err := cmd.Run()
	// if err != nil {
	// 	logger.Info("Container Runtime Engine [ docker/podman ] is not installed..")
	// }

	// Check if Container Runtime daemon is running

	utils.CheckContainerRtStatus(config.ContainerEngine)
}

func BuildJavaImage() {

	// Pull Alpine image
	alpineImage := "alpine:latest"
	//containerEngine.PullDockerImage(alpineImage)
	utils.PullContainerImage(config.ContainerEngine, alpineImage)

	// Pull Corretto image
	correttoImage := fmt.Sprintf("amazoncorretto:%s", os.Getenv("JAVA_TAG"))
	//containerEngine.PullDockerImage(correttoImage)
	utils.PullContainerImage(config.ContainerEngine, correttoImage)

	// Build the java Container Runtime image
	javacorrImage := fmt.Sprintf("javacorr:%s", os.Getenv("JAVA_TAG"))
	javacorrFilePath := os.Getenv("DOCKER_COMPOSE_JAVA_PATH")
	//containerEngine.BuildDockerImage(javacorrImage, javacorrFilePath)
	utils.BuildContainerImage(config.ContainerEngine, javacorrImage, javacorrFilePath)

}

func CreateVolumesAndNetworks() {
	create_volumes()
	create_network()
}

func create_volumes() {

	zooDataVolume := os.Getenv("ZOO_DATA_VOL")
	utils.CreateVolume(config.ContainerEngine, zooDataVolume)

	kafkaDataVolume := os.Getenv("KAFKA_DATA_VOL")
	utils.CreateVolume(config.ContainerEngine, kafkaDataVolume)

	cbDataVolume := os.Getenv("CB_DATA_VOL")
	utils.CreateVolume(config.ContainerEngine, cbDataVolume)

}

func create_network() {
	networkName := os.Getenv("MS_NET")
	utils.CreateNetwork(config.ContainerEngine, networkName)

}

func CreateCbCluster() {

	cbVolumeName := os.Getenv("CB_DATA_VOL")
	cbContainerName := "cb_container"
	utils.CreateCBContainer(config.CONTAINER_RTE, cbContainerName, cbVolumeName)

	if utils.IsContainerRunning(config.ContainerEngine, cbContainerName) {
		couchbase.CouchbaseInit()

	} else {
		logger.Info("Container is not running. Initialization aborted.")
	}

	utils.StopContainer(config.ContainerEngine, cbContainerName)
	utils.RemoveContainer(config.ContainerEngine, cbContainerName)

	logger.Info("CB Cluster Creation completed...")
}
