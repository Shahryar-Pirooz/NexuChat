package config

import (
	"path/filepath"

	"github.com/spf13/viper"
)

func absPath(configPath string) (string, error) {
	if filepath.IsAbs(configPath) {
		return configPath, nil
	}
	return filepath.Abs(configPath)
}

func ReadConfig(configPath string) (Config, error) {
	var cnfg Config
	var err error
	configPath, err = absPath(configPath)
	if err != nil {
		panic("1.unable to read config file: " + err.Error())
	}
	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		panic("2.unable to read config file: " + err.Error())
	}
	return cnfg, viper.Unmarshal(&cnfg)
}

func MustReadConfig(configPath string) Config {
	c, err := ReadConfig(configPath)
	if err != nil {
		panic("unable to read config file: " + err.Error())
	}
	return c
}
