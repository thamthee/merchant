package configs

import (
	"io"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	Mongo struct {
		AuthMechanism  string   `json:"auth-mechanism" mapstructure:"auth-mechanism"`
		AuthSource     string   `json:"auth-source" mapstructure:"auth-source"`
		Username       string   `json:"username" mapstructure:"username"`
		Password       string   `json:"password" mapstructure:"password"`
		Hosts          []string `json:"hosts" mapstructure:"hosts"`
		ReplicaSetName string   `json:"replica-set-name" mapstructure:"replica-set-name"`
		IsDirect       bool     `json:"is-direct" mapstructure:"is-direct"`

		SSL *struct {
			Type string `json:"type" mapstructure:"type"`
			Cert string `json:"cert" mapstructure:"cert"`
		} `json:"ssl" mapstructure:"ssl"`
	} `json:"mongo" mapstructure:"mongo"`

	Web struct {
		APIHost string `json:"api-host" mapstructure:"api-host"`
	} `json:"web" mapstructure:"web"`
}

func ParseFrom(file io.Reader, fileType string) (Config, error) {
	viper.SetConfigType(fileType)
	if err := viper.ReadConfig(file); err != nil {
		return Config{}, errors.Wrap(err, "reading config")
	}

	viper.AutomaticEnv()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, errors.Wrap(err, "unable to decode config")
	}

	return config, nil
}
