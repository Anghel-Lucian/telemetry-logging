# Intro

Telemetry Logging is a Go package that allows:
1. client apps to log to multiple agents at the same time;
2. client app developers to configure which agents receive the logs without touching the core code of the app, being instead configured via a file;
3. prospective developers to develop new drivers for any agent that might receive logs by implementing an interface;

# Requirements, building and running

### Requirements

* Go v1.23.2 (You can probably get away with using almost any version of Go).

### Building and running

* `go build .`;
* `./<whichever-name-you-gave-to-the-package>`.
  
# Functional aspects

The main advantage of this package is that, once the logging code has been written, if the logging agents change (migrating from `.txt` file to `.json`, or from Redis to DynamoDB), the client application's code does not need to change. As for logging, the package supports multiple logging levels and attaching any kind of attribute to the log (be it isolated, or part of a transaction).

## Usage in the code
The package can be imported in other Go packages. It exposes a `GetDefaultLogger()` function that will return the instance of the `Logger` that is to be used throughout the application.

Using this instance, you can call the following functions for your logging needs:
* `Logger.Info()`;
* `Logger.Warn()`;
* `Logger.Debug()`;
* `Logger.Error()`.

Each of these functions take arguments for the actual log message, its attributes if any, and a trace identifier if any (used in the case where you want to correlate multiple logs, such as in a full request-response cycle).

## Configuration

Specifying log levels, different agents/drivers and their configuration options is done entirely through a configuration file (although it is also possible to inject the configuration as a Go object).

The configuration file must be named `logger-config.yml` and it has two main sections:
* `levels`, which is specifies which logging level is to be active;
* and `drivers`, which will determine the components that facilitate the logger's interaction with any writable source and their configuration.

An example of a configuration file:
```
# logger-config.yml
levels:
  info: true
  warn: true
  debug: false # this means that no Debug-level log will be dispatched to the drivers 
  error: true

# this particular configuration means that the logs will be dispatched to the stdout driver (which outputs to stdout) and file driver (which outputs to a text file),
# but the application code does not change regardless of how many drivers or the source they are writing to
drivers: 
  - name: "stdout"
    config:
      color: true
  - name: "file"
    config:
      output-file: "log.txt"
```

## Drivers

The Drivers are the components that are notified when a log was emitted and that write to any source that they are programmed to. The single restriction of a Driver is that it needs adheres to an interface defined in the `logger` package. The rest (meaning what logs it outputs, how it formats them, where it writes them, how it groups them) is determined solely by the driver, and its configuration specified in `logger-config.yml`.

### Adding drivers

We'll consider this a part of the functional specification.

Once a Driver has been developed and implements the `Driver` interface, it needs to be added to the `LoadDrivers` function as an option (i.e., an entry in a `switch` statement). To be registered as a Driver by the client application, its name needs to be specified in the `logger-config.yml` file, as described in the "Configuration" sub-section.

#### Another possible solution for adding Drivers - building them as plugins

This is not implemented, but it is only another possible approach of adding Drivers.

If we want the `logger` package to support addition of drivers without changing at all from a code standpoint, then the Drivers need to be built with the `go build` command as a plugin. Then, given the path to the plugins, the `logger` package can use the go `plugin` package to load these Drivers dynamically.

However, the limitation to this is that the Go version between the `logger` package and the plugins (the Drivers) needs to be the same.


# TODO: add the architecture and the UML diagrams from docs folder





