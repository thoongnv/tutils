package config

import (
    "fmt"
    "github.com/spf13/viper"
)

func ReadConfig(configFileName string, configFilePath string) (*viper.Viper, error) {
    v := viper.New()
    v.SetConfigName(configFileName)
	v.AddConfigPath(configFilePath)

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
    }

    return v, err
}
