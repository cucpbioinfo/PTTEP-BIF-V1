package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Application ApplicationConfig
	Postgres    PgConfig
	Redis       RedisConfig
}

type ApplicationConfig struct {
	Environment string
	Port        int
}

type PgConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	PoolSize int
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	Database string
}

func ReadConfig() Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("[Config] Cannot read config file, %s", err))
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("[Config] Unable to decode into struct, %s", err))
	}

	return config
}
