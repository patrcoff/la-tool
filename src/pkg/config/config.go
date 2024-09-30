package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config holds the global configuration
type Config struct {
	// Add configuration fields here
	DefaultAzureSubscription string
	DefaultResourceGroup     string
}

var globalConfig Config

// Init initializes the configuration
func Init(cfgFile string) error {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".logicappstool" (without extension)
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("error getting user home dir: %w", err)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".logicappstool")
	}

	viper.AutomaticEnv() // Read in environment variables that match

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// Unmarshal the configuration into our struct
	if err := viper.Unmarshal(&globalConfig); err != nil {
		return fmt.Errorf("unable to decode into config struct: %w", err)
	}

	return nil
}

// Get returns the global configuration
func Get() *Config {
	return &globalConfig
}
