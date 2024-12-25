package utils

import (
	"devcd/logger"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func LoadViperConfigObj[T any](v *viper.Viper, configName, configType, filePath string, configObj *T) error {

	// Initialize Viper
	v.SetConfigType(configType)
	v.SetConfigName(configName)
	v.AddConfigPath(filePath)
	err := v.ReadInConfig()
	if err != nil {
		logger.Error("fatal error config file", "error", err)
		return err
	}

	if err := v.Unmarshal(configObj); err != nil {
		logger.Error("fatal error config file marshalling", "error", err)

		return err
	}
	return nil
}

func LoadConfigInViper(v *viper.Viper, configName, configType, configPath string) error {

	// Initialize Viper
	v.SetConfigType(configType)
	v.SetConfigName(configName)
	v.AddConfigPath(configPath)
	err := v.ReadInConfig()
	if err != nil {
		logger.Error("Fatal error config file", "error", err)
		return err
	}
	return nil

}

func MergeConfigInViper(v *viper.Viper, configName, configType, configPath string) error {

	// Initialize Viper
	v.SetConfigType("yaml")
	v.SetConfigName(configName)
	v.AddConfigPath(configPath)
	err := v.MergeInConfig()
	if err != nil {
		logger.Error("Fatal error config file", "error", err)
		return err
	}
	return nil

}

func SetViperConfigAsEnv(v *viper.Viper) {

	// Set environment variables for each key
	for key := range v.AllSettings() {
		logger.Debug("Setting environment variable", "key", key, "value", v.GetString(key))
		os.Setenv(strings.ToUpper(key), v.GetString(key))
		if err := v.BindEnv(key); err != nil {
			logger.Error("Error binding environment variable", "key", key, "error", err)
		}
	}

}
