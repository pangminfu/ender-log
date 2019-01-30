# ender-log
Golang log packages that will broadcast message to list of logger(eg slack, default log) provided to implementation is being added.

## Quick Start
```go
package main

import (
  enderlog "github.com/pangminfu/ender-log"
)

var elog = enderlog.New()

func main() {
  elog.Add(enderlog.NewDefaultSlackRepository("SLACK_WEBHOOK_URL"))
  elog.Add(enderlog.NewDefaultLoggerRepository())
  
  elog.Info("Hello World")
  elog.Warn("Hello World")
  elog.Error("Hello World")
  elog.Fatal("Hello World")
}
```

ender-log have a ready to use default implementation of Slack(DefaultSlackRepository) and default log(DefaultLoggerRepository).
above sample demostrate how to use default provided logger implementaion.

## Custom Logger Implementation
you can have your custom implementation of logger with ender-log. before adding your custom logger implementation, you need to have a CustomLoggerRepository that implement LoggerRepository interface. Below example demonstrate how you can add your own implementation of custom logger:
```go
elog = enderlog.New()
elog.Add(CustomLoggerRepository())
```

Can refer to this [gist](https://gist.github.com/pangminfu/135417ef698d11734f00f29168f81430) for how to write a custom logger.
