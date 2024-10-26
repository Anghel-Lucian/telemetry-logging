package logger

import (
    "testing"

    "github.com/Anghel-Lucian/logger/internal/mocks"
)

func TestDefaultLoggerInstantiation(t *testing.T) {
    resetSingleton()

    dispatcher := &defaultLogDispatcher{}
    mockDriver := &mocks.MockDriver{}

    dispatcher.RegisterDriver(mockDriver)

    mLogger := GetDefaultLogger(dispatcher, &mocks.ConfigAllLevels).(*defaultLogger)

    if mLogger.dispatcher != dispatcher {
        t.Fatalf("mLogger.dispatcher is %v, expected %v", mLogger.dispatcher, dispatcher) 
    }

    mLoggerSingleton := GetDefaultLogger(dispatcher, &mocks.ConfigAllLevels).(*defaultLogger)

    if mLogger != mLoggerSingleton {
        t.Fatalf("GetDefaultLogger does not return the same instance. defaultLogger should be a Singleton")
    }
}

func TestLoadingLoggerConfigAllLevels(t *testing.T) {
    resetSingleton()

    dispatcher := &defaultLogDispatcher{}
    mockDriver := &mocks.MockDriver{}

    dispatcher.RegisterDriver(mockDriver)

    mLogger := GetDefaultLogger(dispatcher, &mocks.ConfigAllLevels).(*defaultLogger)

    if !mLogger.config.Levels.Info || !mLogger.config.Levels.Debug || !mLogger.config.Levels.Warn || !mLogger.config.Levels.Error {
        t.Fatalf("mLogger.config.Levels is %v; expected all levels to be true", mLogger.config.Levels)
    }
}

func TestLoadingLoggerConfigErrorOnly(t *testing.T) {
    resetSingleton()

    dispatcher := &defaultLogDispatcher{}
    mockDriver := &mocks.MockDriver{}

    dispatcher.RegisterDriver(mockDriver)

    mLogger := GetDefaultLogger(dispatcher, &mocks.ConfigErrorOnly).(*defaultLogger)

    if mLogger.config.Levels.Info || mLogger.config.Levels.Debug || mLogger.config.Levels.Warn || !mLogger.config.Levels.Error {
        t.Fatalf("mLogger.config.Levels is %v; expected only 'error' to be true", mLogger.config.Levels)
    }
}

func resetSingleton() {
    mu.Lock()
    defer mu.Unlock()

    defaultLoggerInstance = nil
}


