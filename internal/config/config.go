package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	POSTGRES_DB string `mapstructure:"POSTGRES_DB"`
	Server      ServerConfig
}

type ServerConfig struct {
	Port string
}

func LoadConfig(path string) (Config, error) {
	configPath, err := findConfigPath(path, ".env")
	if err != nil {
		return Config{}, err
	}

	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func CheckConfig(cfg Config) error {
	if cfg.POSTGRES_DB == "" {
		return fmt.Errorf("POSTGRES_DB is required")
	}
	if cfg.Server.Port == "" {
		cfg.Server.Port = "8080"
	}
	return nil
}

func MustLoadConfig(path string) Config {
	cfg, err := LoadConfig(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	if err := CheckConfig(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "config validation failed: %v\n", err)
		os.Exit(1)
	}

	return cfg
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