package viper_test

import (
	"go_test/viper_test/ini_config"
	"go_test/viper_test/json_config"
	"go_test/viper_test/yaml_config"
	"testing"
)

func TestViperYaml(t *testing.T) {
	yaml_config.InitViper("yaml_config")
}

func TestViperJson(t *testing.T) {
	json_config.InitViper("json_config")
}

func TestViperIni(t *testing.T) {
	ini_config.InitViper("ini_config")
}
