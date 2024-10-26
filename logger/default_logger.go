package logger

import (
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/Anghel-Lucian/logger/config"
	"github.com/Anghel-Lucian/logger/driverloader"
	"github.com/Anghel-Lucian/logger/models"
)

type defaultLogger struct {
    dispatcher LogDispatcher
    config config.LoggerConfig
}

var mu *sync.Mutex = &sync.Mutex{}
var defaultLoggerInstance Logger;

func GetDefaultLogger(dispatcher LogDispatcher, lConfig *config.LoggerConfig) Logger {
    mu.Lock()
    defer mu.Unlock()

    if defaultLoggerInstance != nil {
        return defaultLoggerInstance
    }

    if lConfig == nil {
        configFromFile :=config.ReadLoggerConfig()
        lConfig = &configFromFile
    }

    drivers := driverloader.LoadDrivers(lConfig.Drivers)

    if dispatcher == nil {
        dispatcher = &defaultLogDispatcher{}
    }

    for _, d := range drivers {
        dispatcher.RegisterDriver(d)
    }

    defaultLoggerInstance = &defaultLogger{
        dispatcher: dispatcher,
        config: *lConfig,
    }

    return defaultLoggerInstance
}

func (l *defaultLogger) Info(logMsg string, attributes models.LogAttributes, traceID models.LogTraceID) {
    if !l.config.Levels.Info {
        return
    }

    l.dispatcher.NotifyLogInfo(getLog(logMsg, attributes, traceID))
}

func (l *defaultLogger) Warn(logMsg string, attributes models.LogAttributes, traceID models.LogTraceID) {
    if !l.config.Levels.Warn {
        return
    }

    l.dispatcher.NotifyLogWarn(getLog(logMsg, attributes, traceID))
}

func (l *defaultLogger) Debug(logMsg string, attributes models.LogAttributes, traceID models.LogTraceID) {
    if !l.config.Levels.Debug {
        return
    }
    
    l.dispatcher.NotifyLogDebug(getLog(logMsg, attributes, traceID))
}

func (l *defaultLogger) Error(logMsg string, attributes models.LogAttributes, traceID models.LogTraceID) {
    if !l.config.Levels.Error {
        return
    }

    l.dispatcher.NotifyLogError(getLog(logMsg, attributes, traceID))
}

func (l *defaultLogger) GetTraceID() models.LogTraceID {
    return models.LogTraceID(uuid.New().String())
}

func (l *defaultLogger) Shutdown() {
    l.dispatcher.Shutdown()
}

func getLog(logMsg string, attributes models.LogAttributes, traceID models.LogTraceID) models.Log {
    return models.Log{
        LogMsg: logMsg,
        Attributes: attributes,
        Timestamp: time.Now(),
        TraceID: traceID,
    }
}

