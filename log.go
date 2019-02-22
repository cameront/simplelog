package simplelog

import (
	"fmt"
	"os"
	"time"

	kitlog "github.com/go-kit/kit/log"
)

var logger kitlog.Logger

func Debug(a interface{}) {
	log("debug", a)
}

func Debugf(format string, a ...interface{}) {
	log("debug", fmt.Sprintf(format, a...))
}

func Info(a interface{}) {
	log("info", a)
}

func Infof(format string, a ...interface{}) {
	log("info", fmt.Sprintf(format, a...))
}

func Warn(a interface{}) {
	log("warn", a)
}

func Warnf(format string, a ...interface{}) {
	log("warn", fmt.Sprintf(format, a...))
}

func Error(a interface{}) {
	log("error", a)
}

func Errorf(format string, a ...interface{}) {
	log("error", fmt.Sprintf(format, a...))
}

func log(level, a interface{}) {
	if logger == nil {
		w := kitlog.NewSyncWriter(os.Stderr)
		logger = kitlog.NewJSONLogger(w)
	}
	logger.Log("level", level, "timestamp", time.Now().Format(time.RFC3339), "message", a)
}
