package stdoutdriver

import (
	"fmt"

	"github.com/Anghel-Lucian/logger/config"
	"github.com/Anghel-Lucian/logger/driver"
	"github.com/Anghel-Lucian/logger/models"
)

var red = "\033[31m" 
var green = "\033[32m" 
var yellow = "\033[33m" 
var white = "\033[97m"

type StdoutDriver struct {
    color bool
}

func (d *StdoutDriver) OnLogInfo(log models.Log) {
    concatenatedLogString := getConcatenatedLogStrings(log, "INFO")

    if d.color {
        concatenatedLogString = green + concatenatedLogString + "\033[0m"
    } 

    fmt.Println(concatenatedLogString)
}

func (d *StdoutDriver) OnLogWarn(log models.Log) {
    concatenatedLogString := getConcatenatedLogStrings(log, "WARN")

    if d.color {
        concatenatedLogString = white + concatenatedLogString + "\033[0m"
    } 

    fmt.Println(concatenatedLogString)
}

func (d *StdoutDriver) OnLogDebug(log models.Log) {
    concatenatedLogString := getConcatenatedLogStrings(log, "DEBUG")

    if d.color {
        concatenatedLogString = yellow + concatenatedLogString + "\033[0m"
    } 

    fmt.Println(concatenatedLogString)
}

func (d *StdoutDriver) OnLogError(log models.Log) {
    concatenatedLogString := getConcatenatedLogStrings(log, "ERROR")

    if d.color {
        concatenatedLogString = red + concatenatedLogString + "\033[0m"
    } 

    fmt.Println(concatenatedLogString)
}

func (d *StdoutDriver) ReadConfig(config config.DriverConfig) {
    if color, ok := config.Config["color"]; ok {
        if c, ok := color.(bool); ok {
            d.color = c
        }
    }
}

func (d *StdoutDriver) Shutdown() {
}

func GetDriver(config config.DriverConfig) driver.Driver {
    driver := &StdoutDriver{}

    driver.ReadConfig(config)

    return driver
}

func getConcatenatedLogStrings(log models.Log, logLevel string) string {
    logString := "[" + logLevel 

    if log.TraceID != "" {
        logString = logString + ":" + string(log.TraceID) + "]"
    } else {
        logString = logString + "]"
    }

    logString = logString + " "

    return logString + log.LogMsg + " // Attributes: " + getAttributesAsString(log.Attributes)
}

func getAttributesAsString(logAttributes models.LogAttributes) string {
    attributes := ""

    for key, value := range logAttributes {
        attributes = attributes + key + "=" + value + "; " 
    }

    return attributes
}

