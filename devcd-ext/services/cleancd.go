package services

import (
	"devcd/config"
	"devcd/logger"
	"devcd/utils"
	"os"
)

func CleanCdVolumes() {
	delete_volumes()
	delete_network()
}

func delete_volumes() {

	zooDataVolume := os.Getenv("ZOO_DATA_VOL")
	utils.DeleteVolume(config.ContainerEngine, zooDataVolume)

	kafkaDataVolume := os.Getenv("KAFKA_DATA_VOL")
	utils.DeleteVolume(config.ContainerEngine, kafkaDataVolume)

	cbDataVolume := os.Getenv("CB_DATA_VOL")
	utils.DeleteVolume(config.ContainerEngine, cbDataVolume)

}

func delete_network() {
	networkName := os.Getenv("MS_NET")
	utils.DeleteNetwork(config.ContainerEngine, networkName)

}

func CleanMsImg() {
	msImage := "ms-img-container"

	err := utils.RemoveContainerImage(config.ContainerEngine, msImage)
	if err != nil {
		logger.Info("Error removing ms docker image", "error", err)
	}
	logger.Info("Successfully removed docker image", "msImage", msImage)
}
