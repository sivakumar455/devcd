package config

import (
	"devcd/utils"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

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

func TestGetConfigFromFile(t *testing.T) {

	fmt.Println("TEST: TestGetConfigFromFile")

	testFileDir := "testfiles"
	configFile1 := "configfile.yaml"
	fileContent := []byte("# Installation ENV\nJAVA_TAG: '21-alpine'\nTEST_CHASIS_CD: '1.2.20'")

	err := createTempFile(testFileDir, configFile1, fileContent)
	if err != nil {
		t.Errorf("err creating temp file: %v", err)
	}

	defer os.RemoveAll(testFileDir)

	v := viper.New()
	cwd, _ := os.Getwd()
	miniConfigPath := filepath.Join(cwd, "testfiles")
	utils.LoadConfigInViper(v, "configfile", "yaml", miniConfigPath)
	utils.SetViperConfigAsEnv(v)

	//fmt.Println("JAVA_TAG: ", os.Getenv("JAVA_TAG"))
	//fmt.Println("JAVA_TAG: ", v.Get("JAVA_TAG"))

	result := os.Getenv("JAVA_TAG")
	expected := "21-alpine"
	if result != expected {
		t.Errorf("Expected JAVA_TAG: %s got %s", expected, result)
	}

}

func TestGetConfigMergeFromFile(t *testing.T) {

	fmt.Println("TEST: TestGetConfigMergeFromFile")

	testFileDir := "testfiles"
	configFile1 := "configfile.yaml"
	fileContent := []byte("# Installation ENV\nRUN_MODE: 'docker'\nJAVA_TAG: '21-alpine'\nTEST_CHASIS_CD: '1.2.20'")
	err := createTempFile(testFileDir, configFile1, fileContent)
	if err != nil {
		t.Errorf("err creating temp file: %v", err)
	}

	configfile2 := "configfile2.yaml"
	fileContent2 := []byte("RUN_MODE: 'helm' # helm/compose\nFEATURE_CD: 2309")
	err = createTempFile(testFileDir, configfile2, fileContent2)
	if err != nil {
		t.Errorf("err creating temp file: %v", err)
	}

	defer os.RemoveAll(testFileDir)

	v := viper.New()
	cwd, _ := os.Getwd()
	miniConfigPath := filepath.Join(cwd, "testfiles")
	utils.LoadConfigInViper(v, "configfile", "yaml", miniConfigPath)
	utils.MergeConfigInViper(v, "configfile2", "yaml", miniConfigPath)
	utils.SetViperConfigAsEnv(v)
	//fmt.Println("after merge RUN_MODE: ", v.GetString("RUN_MODE"))
	//fmt.Println("after bind env RUN_MODE: ", os.Getenv("RUN_MODE"))

	expected := "helm"
	result := os.Getenv("RUN_MODE")
	if result != expected {
		t.Errorf("Expected RUN_MODE: %s, got %s", expected, result)
	}

}

func TestComposeConfig(t *testing.T) {
	fmt.Println("TEST: TestComposeConfig")

	dirName := "testingfiles"
	fileName := "composefile.yaml"

	fileContent := []byte("composeCharge: \n  - '/ms1/docker-compose.yml'\n  - '/ms2/docker-compose.yml'")

	err := createTempFile(dirName, fileName, fileContent)
	if err != nil {
		t.Errorf("err creating temp file: %v", err)
	}

	defer os.RemoveAll(dirName)

	var composeCfg ComposeConfig
	v := viper.New()

	err = utils.LoadViperConfigObj(v, "composefile", "yaml", dirName, &composeCfg)
	if err != nil {
		t.Errorf("err creating temp file: %v", err)
	}

	expected := composeCfg.ComposeMs[0]
	result := "/ms1/docker-compose.yml"

	if result != expected {
		t.Errorf("Expected: %v got %v", expected, result)
	}

}
