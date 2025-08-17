package config

import (
	"awesomeProject/book/v3/models"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
		Log: &Log{
			level: zerolog.DebugLevel,
		},
	}
}

func (c *Config) String() string {
	return pretty.ToJSON(c)
}

type Config struct {
	Application *application `toml:"app" yaml:"app" json:"app"`
	MySQL       *mySQL       `toml:"mysql" yaml:"mysql" json:"mysql"`
	Log         *Log         `toml:"log" yaml:"log" json:"log"`
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

	db   *gorm.DB
	lock sync.Mutex
}

func (m *mySQL) GetDB() *gorm.DB {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.db == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			m.Username,
			m.Password,
			m.Host,
			m.Port,
			m.DB,
		)
		L().Info().Msgf("database: %s", dsn)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		db.AutoMigrate(&models.Book{})
		m.db = db
	}

	return m.db
}

type Log struct {
	level zerolog.Level `json:"level" toml:"level" yaml:"level" env:"LOG_LEVEL"`

	logger *zerolog.Logger
	lock   sync.Mutex
}

// 为什么这里不加星号
func (l *Log) SetLogger(logger zerolog.Logger) {
	l.logger = &logger
}

func (l *Log) Logger() *zerolog.Logger {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.logger == nil {
		l.SetLogger(zerolog.New(l.ConsoleWriter()).Level(l.level).With().Caller().Timestamp().Logger())
	}

	return l.logger
}

func (c *Log) ConsoleWriter() io.Writer {
	output := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.NoColor = false
		w.TimeFormat = time.RFC3339
	})

	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%-6s", i))
	}

	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}

	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}

	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	return output
}
