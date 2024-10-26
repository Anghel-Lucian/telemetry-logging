package mocks

import (
	"github.com/Anghel-Lucian/logger/config"
	"github.com/Anghel-Lucian/logger/models"
)

type MockDriver struct {
    Logs []models.Log
}

func (d *MockDriver) OnLogInfo(log models.Log) {
    d.Logs = append(d.Logs, log)
}

func (d *MockDriver) OnLogWarn(log models.Log) {
    d.Logs = append(d.Logs, log)
}

func (d *MockDriver) OnLogDebug(log models.Log) {
    d.Logs = append(d.Logs, log)
}

func (d *MockDriver) OnLogError(log models.Log) {
    d.Logs = append(d.Logs, log)
}

func (d *MockDriver) ReadConfig(config config.DriverConfig) {
}

