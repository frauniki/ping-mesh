package config

import (
	"errors"

	"github.com/frauniki/ping-mesh/pkg/validator"
	"github.com/spf13/viper"
)

const (
	defaultConfigPath     = "configs/"
	defaultConfigFileName = "config"
)

type Ping struct {
	Count    int    `mapstructure:"count" validate:"omitempty,min=1,max=10"`
	Interval string `mapstructure:"interval" validate:"duration-string"`
	Timeout  string `mapstructure:"timeout" validate:"duration-string"`
}

type Host struct {
	Name string `mapstructure:"name" validate:"required"`
	Host string `mapstructure:"host" validate:"required"`
}

type GlobalConfig struct {
	ServerURL string `mapstructure:"server_url" validate:"omitempty,url"`
	Interval  string `mapstructure:"interval" validate:"duration-string"`
	LogLevel  string `mapstructure:"log_level"`
	Ping      Ping   `mapstructure:"ping" validate:"required"`
	Hosts     []Host `mapstructure:"hosts" validate:"required"`
}

var Config *GlobalConfig

func validate() error {
	if Config == nil {
		return errors.New("Config is not loaded")
	}

	v := validator.NewValidator()
	return v.Struct(Config)
}

func Load(file string) (*GlobalConfig, error) {
	viper.SetConfigType("yaml")
	setDefaultValues()

	if file != "" {
		viper.SetConfigFile(file)
	} else {
		viper.AddConfigPath(defaultConfigPath)
		viper.SetConfigName(defaultConfigFileName)
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&Config); err != nil {
		return nil, err
	}

	if err := validate(); err != nil {
		return nil, err
	}

	return Config, nil
}
