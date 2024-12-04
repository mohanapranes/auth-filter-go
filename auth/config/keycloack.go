package config

import (
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type KeycloakConfig struct {
	Realm         string
	ClientID      string
	ClientSecret  string
	IntrospectURL string
}

var (
	once           sync.Once
	configInstance *KeycloakConfig
)

func NewKeycloakConfig() *KeycloakConfig {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}
