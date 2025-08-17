package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"os"
)

var config *Config

func C() *Config {
	if config == nil {
		config = Default()
	}
	return config
}

func DB() *gorm.DB {
	return C().MySQL.GetDB()
}

func L() *zerolog.Logger {
	return C().Log.Logger()
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
