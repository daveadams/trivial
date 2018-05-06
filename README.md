# trivial

Trivial support libraries for golang, including a stdout logger and a
concurrency-safe counter.


## Logger

Usage example:

    package main

    import "github.com/daveadams/trivial"

    var log trivial.Logger = trivial.Logger{
        IncludeTimestamp: true,
        Level:            trivial.InfoLevel,
    }

    func main() {
        t := "once"
        log.Infof("This message will be logged %s.", t)
        log.Debug("This message will not be logged.")
    }

Newlines are automatically appended. That's basically it.


## Counter

Usage example:

    package main

    import "github.com/daveadams/trivial"

    var log trivial.Logger = trivial.Logger{}
    var counter trivial.Counter = trivial.Counter{}

    func main() {
        log.Infof("Initial count is %d", counter.Count())
        counter.Increment()
        log.Infof("New count is %d", counter.Count())
    }


## License

This software is public domain. No rights are reserved. See LICENSE for more
information.
