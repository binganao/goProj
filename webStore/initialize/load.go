package initialize

import (
	"log"

	"mall/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Panicln(err)
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		log.Panicln("unable to decode into struct", err)
	}
}
