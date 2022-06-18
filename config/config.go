package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	MongoURI         string `mapstructure:"MONGO_URI"`
	MongoURITest     string `mapstructure:"MONGO_URI_TEST"`
	MongoDB          string `mapstructure:"MONGO_DATABASE"`
	MongoPoolMin     uint64 `mapstructure:"MONGO_POOL_MIN"`
	MongoPoolMax     uint64 `mapstructure:"MONGO_POOL_MAX"`
	MongoMaxIdleTime int64  `mapstructure:"MONGO_MAX_IDLE_TIME_SECOND"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
