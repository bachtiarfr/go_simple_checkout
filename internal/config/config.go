package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host         string `yaml:"db_host"`
	Port         string `yaml:"db_port"`
	Username     string `yaml:"db_username"`
	Password     string `yaml:"db_password"`
	Databasename string `yaml:"db_name"`
	SecretKey    string `yaml:"secret_key"`
}

func ReadConfig(filePath string) (*Config, error) {
	config := &Config{}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
