package logger

import (
    "github.com/Anghel-Lucian/logger/driver"    
    "github.com/Anghel-Lucian/logger/models"    
)

type LogDispatcher interface {
    RegisterDriver(driver driver.Driver)
    RemoveDriver(driver driver.Driver)
    NotifyLogInfo(log models.Log)
    NotifyLogWarn(log models.Log)
    NotifyLogDebug(log models.Log)
    NotifyLogError(log models.Log)
    Shutdown()
}

