package logger

import (
	"github.com/Anghel-Lucian/logger/driver"
	"github.com/Anghel-Lucian/logger/models"
)

type defaultLogDispatcher struct {
    drivers []driver.Driver
}

func (ld *defaultLogDispatcher) RegisterDriver(driver driver.Driver) {
    ld.drivers = append(ld.drivers, driver)
}

func (ld *defaultLogDispatcher) RemoveDriver(driver driver.Driver) {
    driverIndex := -1

    for i, d := range ld.drivers {
        if d == driver {
            driverIndex = i
            break
        }
    }

    if driverIndex == -1 {
        return
    }

    ld.drivers = append(ld.drivers[:driverIndex], ld.drivers[driverIndex + 1:]...)
}

func (ld *defaultLogDispatcher) NotifyLogInfo(log models.Log) {
    for _, d := range ld.drivers {
        d.OnLogInfo(log)
    }
}

func (ld *defaultLogDispatcher) NotifyLogWarn(log models.Log) {
    for _, d := range ld.drivers {
        d.OnLogWarn(log)
    }
}

func (ld *defaultLogDispatcher) NotifyLogDebug(log models.Log) {
    for _, d := range ld.drivers {
        d.OnLogDebug(log)
    }
}

func (ld *defaultLogDispatcher) NotifyLogError(log models.Log) {
    for _, d := range ld.drivers {
        d.OnLogError(log)
    }
}

func (ld *defaultLogDispatcher) Shutdown() {
    for _, d := range ld.drivers {
        d.Shutdown()
    }
}

