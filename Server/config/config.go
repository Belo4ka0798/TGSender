package config

import (
	"github.com/spf13/viper"
	"sync"
	"tgsender/pkg/logging"
)

func init() {
	viper.AddConfigPath("./config/")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
}

var once sync.Once

func GetConfigs() {
	once.Do(func() {
		logger := logging.GetLogger()
		err := viper.ReadInConfig()
		if err != nil {
			logger.Fatal(err)
		}
	})
}
