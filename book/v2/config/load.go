package config

import (
	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v3"
	"os"
)

var config *Config

func C() *Config {
	if config == nil {
		config = Default()
	}
	return config
}

// 把系统yaml文件映射到配置类
func LoadConfigfromYaml(configPath string) error {
	content, err := os.ReadFile(configPath)
	if config != nil {
		return err
	}
	config = C()
	return yaml.Unmarshal(content, config)
}

func LoadConfigFromEnv() error {
	config = C()
	return env.Parse(config)
}
