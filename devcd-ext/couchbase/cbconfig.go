package couchbase

import (
	"devcd/logger"
	"devcd/utils"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// ClusterInfo represents the cluster configuration
type CbConfigInfo struct {
	ClusterInfo struct {
		ClusterName      string `mapstructure:"clusterName"`
		StorageMode      string `mapstructure:"storageMode"`
		Username         string `mapstructure:"username"`
		Password         string `mapstructure:"password"`
		Port             string `mapstructure:"port"`
		MemoryQuota      int    `mapstructure:"memoryQuota"`
		IndexMemoryQuota int    `mapstructure:"indexMemoryQuota"`
		BaseURL          string `mapstructure:"baseUrl"`
		Services         string `mapstructure:"services"`
	} `mapstructure:"clusterInfo"`

	BucketInfo []struct {
		Name  string `mapstructure:"name"`
		Type  string `mapstructure:"type"`
		Scope []struct {
			Name       string `mapstructure:"name"`
			Collection []struct {
				Name  string   `mapstructure:"name"`
				Index []string `mapstructure:"index"`
			} `mapstructure:"collection"`
		} `mapstructure:"scope"`
	} `mapstructure:"bucketInfo"`
}

var CBConfig CbConfigInfo

const cbConfigPath string = "/devcd-ext/couchbase"

func LoadCbConfig() {
	v := viper.New()
	cwd, _ := os.Getwd()
	configPath := filepath.Join(cwd, cbConfigPath)
	cbConfigName := "cb_init_config"
	configType := "yaml"

	err := utils.LoadViperConfigObj[CbConfigInfo](v, cbConfigName, configType, configPath, &CBConfig)
	if err != nil {
		logger.Error("Error loading cb config")
	}

	// Access the parsed data
	logger.Debug("Cluster Name", "cluster", CBConfig.ClusterInfo.ClusterName)
	logger.Debug("Bucket 1 Name", CBConfig.BucketInfo[0].Name)
	logger.Debug("Bucket 2 Name", CBConfig.BucketInfo[1].Name)

}
