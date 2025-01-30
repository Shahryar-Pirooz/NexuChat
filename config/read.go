package config

import (
	"path/filepath"

	"github.com/spf13/viper"
)

func absPath(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}
	return filepath.Abs(path)

}

func Read(path string) (Config, error) {
	var cnfg Config
	absPath, err := absPath(path)
	if err != nil {
		return cnfg, err
	}
	viper.SetConfigFile(absPath)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return cnfg, err
	}
	return cnfg, viper.Unmarshal(&cnfg)
}

func MustRead(path string) Config {
	config, err := Read(path)
	if err != nil {
		panic(err)
	}
	return config
}
