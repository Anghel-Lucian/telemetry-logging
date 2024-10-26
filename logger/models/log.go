package models

import (
    "time"
)

type LogAttributes map[string]string

type LogTraceID string

type Log struct {
    LogMsg string
    Attributes LogAttributes
    TraceID LogTraceID
    Timestamp time.Time
}

