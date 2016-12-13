package test

import (
	"gopkg.in/yaml.v2"
)

type TestConfig struct {
	Name string `yaml:"name"`
}

func (config TestConfig) GetConfig(buf []byte) (interface{}, error) {
	err := yaml.Unmarshal(buf, &config)
	return config, err
}
