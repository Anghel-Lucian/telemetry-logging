package config 

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

const LOGGER_CONFIG_FILENAME = "logger-config.yml"

type DriverConfig struct {
    Name string `yaml:"name"`
    Config map[string]interface{} `yaml:"config"`
}

type LogLevelsConfig struct {
    Info bool `yaml:"info"`
    Warn bool `yaml:"warn"`
    Debug bool `yaml:"debug"`
    Error bool `yaml:"error"`
}

type LoggerConfig struct {
    Levels LogLevelsConfig `yaml:"levels"`
    Drivers []DriverConfig `yaml:"drivers"`
}

func ReadLoggerConfig() LoggerConfig {
    configFile, err := os.ReadFile(LOGGER_CONFIG_FILENAME)

    if err != nil {
        panic("Failed to read logger configuration file. The file should be '" + LOGGER_CONFIG_FILENAME + "'.")
    }

    config := &LoggerConfig{}

    err = yaml.Unmarshal(configFile, config)

    if err != nil {
        panic("Failed to marshall logger YAML configuration.")
    }

    return *config
}

