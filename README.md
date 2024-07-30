# Custom Logger

A simple go logger which is based on the go library log and extends it with standardized log outputs according to stout.

## Features

- Supports log levels: INFO, WARN, ERROR, DEBUG
- Log levels can be dynamically set and queried
- Timestamps for each log message
- Easy integration and usage

## Installation

Clone the repository and import the package into your Go project:

```sh
go get github.com/head-crash/logger
```

In your Go file:

```go
import "github.com/head-crash/logger"
```

## Usage

### Initialize Logger

Create a shortcut for the default logger instance:

```go
log := logger.Default
```

### Set Log Level

Set the desired log level:

```go
log.SetLogLevel(logger.DEBUG)
```

### Write Log Messages

Write log messages at various log levels:

```go
log.Info("This is an info message")
log.Warn("This is a warning")
log.Error("This is an error message")
log.Debug("This is a debug message")
```

### Query Current Log Level

Query the current log level:

```go
currentLogLevel := log.GetLogLevel()
fmt.Println("Current log level:", currentLogLevel)
```

### Formatted Log Messages

You can also write formatted log messages using `Printf` style formatting:

```go
log.Info("User %s has logged in", "Alice")
log.Warn("Disk space is low: %d%% remaining", 10)
log.Error("Failed to open file: %s", "example.txt")
log.Debug("Processing item %d of %d", 1, 10)
```

## Example

Here is a complete example:

```go
package main

import (
 "fmt"
 "github.com/head-crash/logger"
)

func main() {
 log := logger.Default
 log.SetLogLevel(logger.DEBUG)

 log.Info("This is an info message")
 log.Warn("This is a warning")
 log.Error("This is an error message")
 log.Debug("This is a debug message")

 log.Info("User %s has logged in", "Alice")
 log.Warn("Disk space is low: %d%% remaining", 10)
 log.Error("Failed to open file: %s", "example.txt")
 log.Debug("Processing item %d of %d", 1, 10)

 currentLogLevel := log.GetLogLevel()
 fmt.Println("Current log level:", currentLogLevel)
}
```

## License

This project is licensed under the GNU GENERAL PUBLIC LICENSE. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or a pull request for suggestions and improvements.
