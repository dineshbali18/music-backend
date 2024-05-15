package config

import "github.com/spf13/viper"

func InitializeConfig() {
	viper.SetConfigFile(`../config.yml`)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
