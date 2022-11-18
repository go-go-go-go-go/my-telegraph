package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Configs struct {
	Host         string `yaml:"host"`
	ExternalHost string `yaml:"external_host"`
	Port         int    `yaml:"port"`
	StorageType  string `yaml:"storage_type"`
	DbUrl        string `yaml:"db_url"`
}

var configs_inited bool = false
var configs Configs

func GetConfigs() Configs {
	if !configs_inited {
		yaml_bytes, err := os.ReadFile("config/configs.yaml")
		if err != nil {
			panic(err)
		}
		println(string(yaml_bytes))
		err = yaml.Unmarshal(yaml_bytes, &configs)
		if err != nil {
			panic(err)
		}
		configs_inited = true
	}
	return configs
}
