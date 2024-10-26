package logger

import (
    "github.com/Anghel-Lucian/logger/models"
)

type Logger interface {
    Info(logMsg string, attributes models.LogAttributes, traceID models.LogTraceID)
    Warn(logMsg string, attributes models.LogAttributes, traceID models.LogTraceID)
    Debug(logMsg string, attributes models.LogAttributes, traceID models.LogTraceID)
    Error(logMsg string, attributes models.LogAttributes, traceID models.LogTraceID)
    GetTraceID() models.LogTraceID 
    Shutdown()
}

