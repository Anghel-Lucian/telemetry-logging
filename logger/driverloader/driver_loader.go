package driverloader

import (
	"github.com/Anghel-Lucian/logger/config"
	"github.com/Anghel-Lucian/logger/driver"
	"github.com/Anghel-Lucian/logger/driver/filedriver"
	"github.com/Anghel-Lucian/logger/driver/stdoutdriver"
)

func LoadDrivers(configs []config.DriverConfig) []driver.Driver {
    drivers := []driver.Driver{}

    for _, c := range configs {
        switch c.Name {
        case "file":
            drivers = append(drivers, filedriver.GetDriver(c))
        case "stdout":
            drivers = append(drivers, stdoutdriver.GetDriver(c))
        }
    }

    return drivers
}
