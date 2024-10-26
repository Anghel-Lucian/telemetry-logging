package mocks

import (
    "github.com/Anghel-Lucian/logger/config"
)

var ConfigAllLevels config.LoggerConfig = config.LoggerConfig{
    Levels: config.LogLevelsConfig{
        Info: true,
        Warn: true,
        Debug: true,
        Error: true,
    },
    Drivers: nil,
}

var ConfigErrorOnly config.LoggerConfig = config.LoggerConfig{
    Levels: config.LogLevelsConfig{
        Info: false,
        Warn: false,
        Debug: false,
        Error: true,
    },
    Drivers: nil,
}

