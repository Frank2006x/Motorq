package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	POSTGRES_DB          string `mapstructure:"POSTGRES_DB"`

}

func LoadConfig(path string) (Config, error) {
	configPath, err := findConfigPath(path, ".env")
	if err != nil {
		return Config{}, err
	}

	viper.SetConfigFile(configPath)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	var config Config

	err = viper.Unmarshal(&config)

	return config, err

}

func findConfigPath(startDir, fileName string) (string, error) {
	currentDir, err := filepath.Abs(startDir)
	if err != nil {
		return "", err
	}

	for {
		candidate := filepath.Join(currentDir, fileName)
		if _, statErr := os.Stat(candidate); statErr == nil {
			return candidate, nil
		} else if !os.IsNotExist(statErr) {
			return "", statErr
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			break
		}
		currentDir = parentDir
	}

	return "", fmt.Errorf("config file %q not found starting from %q", fileName, startDir)
}