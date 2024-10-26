package driver

import (
	"github.com/Anghel-Lucian/logger/config"
	"github.com/Anghel-Lucian/logger/models"
)

type Driver interface {
    OnLogInfo(log models.Log)
    OnLogWarn(log models.Log)
    OnLogDebug(log models.Log)
    OnLogError(log models.Log)
    ReadConfig(config config.DriverConfig)
    Shutdown()
}

