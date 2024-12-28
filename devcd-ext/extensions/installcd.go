package extensions

import (
	"devcd/logger"
	"devcd_ext/services"
)

func InstallFulldevc() {

	logger.Info("###############################################")
	logger.Info("devcd installation STARTED")
	logger.Info("################################################")

	services.InstallJavaImage()
	services.CreateVolumesAndNetworks()
	//CreateCbCluster()

	ExtractMsJars()

	logger.Info("#################################################")
	logger.Info("devcd installation COMPLETED")
	logger.Info("#################################################")
}

func ExtractMsJars() {
	logger.Info("Extracting all MS jars")
	NewExtractMs()
}
