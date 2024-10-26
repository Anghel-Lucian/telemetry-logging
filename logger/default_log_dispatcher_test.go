package logger

import (
	"testing"

	"github.com/Anghel-Lucian/logger/internal/mocks"
	"github.com/Anghel-Lucian/logger/models"
    "github.com/Anghel-Lucian/logger/driver"
)

func TestDefaultDispatcherDriverRegistration(t *testing.T) {
    resetSingleton()

    dispatcher := &defaultLogDispatcher{}
    mockDriver := &mocks.MockDriver{}

    dispatcher.RegisterDriver(mockDriver)

    if len(dispatcher.drivers) != 1 || dispatcher.drivers[0] != mockDriver {
        t.Fatalf("dispatcher.drivers is %v, expected %v", dispatcher.drivers, []driver.Driver{mockDriver})
    }
}

func TestDefaultDispatcherDriverNotifiesDrivers(t *testing.T) {
    resetSingleton()

    dispatcher := &defaultLogDispatcher{}
    mockDriver := &mocks.MockDriver{}

    dispatcher.RegisterDriver(mockDriver)

    mLogger := GetDefaultLogger(dispatcher, &mocks.ConfigAllLevels).(*defaultLogger)
   
    infoMsg := "Test info msg"
    errorMsg := "Test error msg"
    debugTransaction1 := "Debug transaction"
    debugTransaction2 := "Debug transaction 2"
    warnMsg := "Warning info msg"

    mLogger.Info(infoMsg, nil, "")
    mLogger.Error(errorMsg, models.LogAttributes{"error": "attr"}, "")
    mLogger.Debug(debugTransaction1, models.LogAttributes{}, "1")
    mLogger.Debug(debugTransaction2, models.LogAttributes{}, "1")
    mLogger.Warn(warnMsg, nil, "")

    if len(mockDriver.Logs) != 5 {
        t.Fatalf("mockDriver.Logs length is %v, expected 5", len(mockDriver.Logs)) 
    }

    transactionalLogs := []models.Log{}

    for _, l := range mockDriver.Logs {
        if l.TraceID != "" {
            transactionalLogs = append(transactionalLogs, l)
        }
    }

    if len(transactionalLogs) != 2 {
        t.Fatalf("transactionalLogs length is %v, expected 2", len(transactionalLogs))
    }

    if transactionalLogs[0].LogMsg != debugTransaction1 && transactionalLogs[1].LogMsg != debugTransaction2 {
        t.Fatalf("transactionalLogs[0].LogMsg is %v, transactionalLogs[1].LogMsg is %v, expected %v and %v", transactionalLogs[0].LogMsg, transactionalLogs[1].LogMsg, debugTransaction1, debugTransaction2)
    }
}

