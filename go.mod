module github.com/Anghel-Lucian/telemetry-logging

go 1.23.2

require github.com/Anghel-Lucian/logger v1.0.0

require (
	github.com/google/uuid v1.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/Anghel-Lucian/logger => ./logger
