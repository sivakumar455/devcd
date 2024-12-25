package config

import (
	"devcd/logger"
	"devcd/utils"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type HelmConfig struct {
	HelmBs     []string
	HelmMs     []string
	HelmTestBs []string
	HelmTestMs []string
	ValuesMs   []string
	ValuesBs   []string
}

type ComposeConfig struct {
	ComposeBs     []string
	ComposeMs     []string
	ComposeTestBs []string
	ComposeTestMs []string
}

var v *viper.Viper
var ComposeCfg ComposeConfig
var HelmCfg HelmConfig

var (
	containerRte    string
	FEATURE_CD      string
	CONTAINER_RTE   string
	RUN_MODE        string
	BS_COMPOSE_HOME string
	MS_COMPOSE_HOME string
	BS_HELM_HOME    string
	MS_HELM_HOME    string
	DEVCD_RUNTIME   string
)

const helmConfigHome string = "/devcd-helm"
const composeConfigHome string = "/devcd-compose"
const devcHome string = "/devcd"

const core_config_values string = "core_config_values"
const cd_global_config string = "cd_global_config"

const helm_env_values string = "helm_env_values"

const ms_compose_files string = "ms_compose_files"
const helm_ms_config string = "helm_ms_config"

var ContainerEngine *utils.ContainerRuntime

func init() {
	v = viper.New()
}

// Load all global configs and helm and compose file configs
func LoadGlobalConfig() error {

	//cwd, _ := os.Getwd()
	exePath, err := os.Executable()
	if err != nil {
		logger.Error("Err getting exe path")
		return err
	}
	cwd := filepath.Dir(exePath)
	logger.Debug("Current working dir", "cwd", cwd)

	// internal config
	defaultConfigPath := filepath.Join(cwd, devcHome)
	err = utils.LoadConfigInViper(v, core_config_values, "yaml", defaultConfigPath)
	if err != nil {
		logger.Error("Err in LoadConfigInViper")
		return err
	}
	// user global config
	//configPath := filepath.Join(cwd, cdConfig)
	err = utils.MergeConfigInViper(v, cd_global_config, "yaml", cwd)
	if err != nil {
		logger.Error("Error in MergeConfigInViper")
		return err
	}
	// user helm level config
	configPath := filepath.Join(cwd, helmConfigHome, "env-helm")
	err = utils.MergeConfigInViper(v, helm_env_values, "yaml", configPath)
	if err != nil {
		logger.Error("Error in MergeConfigInViper")
		return err
	}

	utils.SetViperConfigAsEnv(v)

	initEnv()

	err = loadHelmComposeFiles(cwd)
	if err != nil {
		logger.Error("Error in Loading helm and composefiles")
		return err
	}

	return err

}

func initEnv() {

	ContainerEngine = &utils.ContainerRuntime{CrtEngine: os.Getenv("CONTAINER_RT_ENGINE")}
	checkForEmptyValue("CONTAINER_RT_ENGINE", ContainerEngine.CrtEngine)

	CONTAINER_RTE = os.Getenv("CONTAINER_RT_ENGINE")
	checkForEmptyValue("CONTAINER_RT_ENGINE", CONTAINER_RTE)

	containerRte = os.Getenv("CONTAINER_RT_ENGINE")
	checkForEmptyValue("CONTAINER_RT_ENGINE", containerRte)

	FEATURE_CD = os.Getenv("FEATURE_CD")
	checkForEmptyValue("FEATURE_CD", FEATURE_CD)

	RUN_MODE = os.Getenv("RUN_MODE")
	checkForEmptyValue("RUN_MODE", RUN_MODE)

	BS_COMPOSE_HOME = os.Getenv("BS_COMPOSE_HOME")
	checkForEmptyValue("BS_COMPOSE_HOME", BS_COMPOSE_HOME)

	MS_COMPOSE_HOME = os.Getenv("MS_COMPOSE_HOME")
	checkForEmptyValue("MS_COMPOSE_HOME", MS_COMPOSE_HOME)

	BS_HELM_HOME = os.Getenv("BS_HELM_HOME")
	checkForEmptyValue("BS_HELM_HOME", BS_HELM_HOME)

	MS_HELM_HOME = os.Getenv("MS_HELM_HOME")
	checkForEmptyValue("MS_HELM_HOME", MS_HELM_HOME)

	DEVCD_RUNTIME = os.Getenv("DEVCD_RUNTIME")
	checkForEmptyValue("DEVCD_RUNTIME", DEVCD_RUNTIME)
}

func checkForEmptyValue(key, value string) {
	if value == "" {
		logger.Warn("Environment variable is not set ", "key", key)
	}
}

func loadHelmComposeFiles(cwd string) error {
	// compose file config
	composeConfigPath := filepath.Join(cwd, composeConfigHome)
	err := utils.LoadViperConfigObj[ComposeConfig](v, ms_compose_files, "yaml", composeConfigPath, &ComposeCfg)
	if err != nil {
		logger.Error("fatal error config file", "error", err)
		return err
	}
	logger.Debug("compose MS files", "composeMs", ComposeCfg.ComposeMs)
	logger.Debug("Compose BS files", "composeBs", ComposeCfg.ComposeBs)

	// compose file config
	helmConfigPath := filepath.Join(cwd, helmConfigHome)
	err = utils.LoadViperConfigObj[HelmConfig](v, helm_ms_config, "yaml", helmConfigPath, &HelmCfg)
	if err != nil {
		logger.Error("fatal error config file", "error", err)
		return err
	}
	logger.Debug("helm MS files", "helmMs", HelmCfg.HelmMs)
	logger.Debug("helm BS files", "helmBs", HelmCfg.HelmBs)
	return err
}
