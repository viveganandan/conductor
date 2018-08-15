package datadog

import (
	"fmt"
	"os"

	"github.com/DataDog/datadog-go/statsd"

	"github.com/Nextdoor/conductor/shared/logger"
)

var c = func() *statsd.Client {
	c, err := statsd.New(fmt.Sprintf("%v:%v", os.Getenv("STATSD_HOST"), 8125))
	if err != nil {
		logger.Error("Could not create statsd client: %s", err)
		return nil
	}
	return c
}()

func log(alertType statsd.EventAlertType, format string, args ...interface{}) {
	// Send event to statsd and log it too!
	if c != nil {
		e := statsd.NewEvent("conductor", fmt.Sprintf(format, args...))
		e.AlertType = alertType
		err := c.Event(e)
		if err != nil {
			logger.Error("Could not create datadog event: %v", err)
		}
	}
	switch alertType {
	case statsd.Info:
		logger.Info(format, args)
	default:
		logger.Error(format, args)
	}
}

func Info(format string, args ...interface{}) {
	log(statsd.Info, format, args)
}

func Error(format string, args ...interface{}) {
	log(statsd.Error, format, args)
}
