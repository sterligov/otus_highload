package config

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	HTTP struct {
		Addr           string        `yaml:"addr"`
		ReadTimeout    time.Duration `yaml:"read_timeout"`
		WriteTimeout   time.Duration `yaml:"write_timeout"`
		HandlerTimeout time.Duration `yaml:"handler_timeout"`
	} `yaml:"http"`

	GRPC struct {
		Addr string `yaml:"addr"`
	} `yaml:"grpc"`

	Database struct {
		Addr            string        `yaml:"connection_addr"`
		Driver          string        `yaml:"driver"`
		MaxOpenConns    int           `yaml:"max_open_conns"`
		MaxIdleConns    int           `yaml:"max_idle_conns"`
		MaxConnLifetime time.Duration `yaml:"max_conn_lifetime"`
	} `yaml:"database"`

	Logger struct {
		Path  string `yaml:"path"`
		Level string `yaml:"level"`
	} `yaml:"logger"`

	StorageType string `yaml:"storage_type"`

	JWT struct {
		SecretKey string `yaml:"secret_key"`
	}
}

func New(cfgFilename string) (*Config, error) {
	f, err := os.Open(cfgFilename)
	if err != nil {
		return nil, fmt.Errorf("open config file failed: %w", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			logrus.Warnf("config file close: %s", err)
		}
	}()

	cfg := &Config{}

	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(cfg); err != nil {
		return nil, fmt.Errorf("decode config file failed: %w", err)
	}

	return cfg, err
}
