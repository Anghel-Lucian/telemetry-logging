package filedriver 

import (
	"os"

	"github.com/Anghel-Lucian/logger/config"
	"github.com/Anghel-Lucian/logger/driver"
	"github.com/Anghel-Lucian/logger/models"
)

type FileDriver struct {
    file *os.File
}

func (d *FileDriver) OnLogInfo(log models.Log) {
    d.file.WriteString(log.LogMsg + "\n")
}

func (d *FileDriver) OnLogWarn(log models.Log) {
    d.file.WriteString(log.LogMsg + "\n")
}

func (d *FileDriver) OnLogDebug(log models.Log) {
    d.file.WriteString(log.LogMsg + "\n")
}

func (d *FileDriver) OnLogError(log models.Log) {
    d.file.WriteString(log.LogMsg + "\n")
}

func (d *FileDriver) ReadConfig(config config.DriverConfig) {
    filename := config.Config["output-file"].(string)
    
    file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)

    if err != nil {
        panic("FileDriver could not open output file")
    }

    d.file = file
}

func (d *FileDriver) Shutdown() {
    d.file.Close()
}

func GetDriver(config config.DriverConfig) driver.Driver {
    fileDriver := &FileDriver{}

    fileDriver.ReadConfig(config)

    return fileDriver
}

