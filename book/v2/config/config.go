package config

import (
	"github.com/infraboard/mcube/v2/tools/pretty"
)

func Default() *Config {
	return &Config{
		Application: &application{
			Host: "127.0.0.1",
			Port: 8080,
		},
		MySQL: &mySQL{
			Host:     "127.0.0.1",
			Port:     3306,
			DB:       "go18",
			Username: "root",
			Password: "dywszz",
			Debug:    true,
		},
	}
}

func (c *Config) String() string {
	return pretty.ToJSON(c)
}

type Config struct {
	Application *application `toml:"app" yaml:"app" json:"app"`
	MySQL       *mySQL       `toml:"mysql" yaml:"mysql" json:"mysql"`
}
type application struct {
	Host string `toml:"host" yaml:"host" json:"host" env:"HOST"`
	Port int    `toml:"port" yaml:"port" json:"port" env:"PORT"`
}
type mySQL struct {
	Host     string `toml:"host" yaml:"host" json:"host" env:"DATASOURCE_HOST"`
	Port     int    `toml:"port" yaml:"port" json:"port" env:"DATASOURCE_PORT"`
	DB       string `toml:"database" yaml:"database" json:"database" env:"DATASOURCE_DB"`
	Username string `toml:"username" yaml:"username" json:"username" env:"DATASOURCE_USERNAME"`
	Password string `toml:"password" yaml:"password" json:"password" env:"DATASOURCE_PASSWORD"`
	Debug    bool   `toml:"debug" yaml:"debug" json:"debug" env:"DATASOURCE_DEBUG"`
}
