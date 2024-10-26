package main

import (
    "github.com/Anghel-Lucian/logger"
    logModels "github.com/Anghel-Lucian/logger/models"
)

// As long as logger-config.yaml is present in the project root,
// the config will be read from there.
// However, there's also the option of initializing the logger with config
// object defined in the source code files.
// In the same manner, the log dispatcher (the component responsible with
// notifying drivers that a log has been emitted) is also injectable.

// The logger-config.yaml is also responsible with what drivers
// are loaded for the logger to emit to. If there are multiple drivers registered,
// all of them will receive the logs.
// For example, since this project's logger-config.yaml specifies both stdoutdriver
// and filedriver, the logs will be written both to stdout and also to a file with the same
// name that's specified in logger-config.yaml.

// In this manner, the client app developer can add and remove drivers as required.
func main() {
    mLog := logger.GetDefaultLogger(nil, nil)
   
    // There are 4 levels of logging: info, warn, debug and error
    mLog.Info("This is an informational message", nil, "")

    mLog.Debug(
        "This is a debug message which also has attributes",
        logModels.LogAttributes{
            "attribute": "value",
        },
        "",
    )

    // In addition to the log message and attributes, the client can also specify
    // a traceID which can be used to query logs that are related to the same operation.
    // How those logs are emitted and how they are grouped by the traceID is up to the
    // driver implementation.
    // Another language-specific pattern one could use is passing a Context object
    // instead of a TraceID.
    traceID := mLog.GetTraceID()

    mLog.Error(
        "An error occured when some transaction took place",
        logModels.LogAttributes{
            "transactionalLogsAreTheSameAsNormalLogs": "but with a trace ID",
        },
        traceID,
    )

    mLog.Info(
        "Rolling back transaction",
        nil,
        traceID,
    )

    // Some drivers could write to resources that need to be explicitly closed (such as
    // files, or database connections), and so the Logger interface has a Shutdown method
    // that in turn gracefully shuts down the dispatcher, and by extension any drivers that
    // are registered.
    mLog.Shutdown()
}

