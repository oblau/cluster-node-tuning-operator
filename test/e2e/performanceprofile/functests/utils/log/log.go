package log

import (
	"fmt"
	"time"

	"github.com/onsi/ginkgo/v2"
)

type Level int

const (
	LevelNone    Level = 0
	LevelError   Level = 1
	LevelWarning Level = 2
	LevelInfo    Level = 3
)

var minLevel = LevelInfo

// SetLevel sets the minimum log level
func SetLevel(level Level) {
	minLevel = level
}

// stringToLevel maps the hardcoded strings to levels
func stringToLevel(levelStr string) Level {
	switch levelStr {
	case "[ERROR]":
		return LevelError
	case "[WARNING]":
		return LevelWarning
	case "[INFO]":
		return LevelInfo
	default:
		return LevelInfo
	}
}

func nowStamp() string {
	return time.Now().Format(time.StampMilli)
}

func logf(level string, format string, args ...interface{}) {
	if stringToLevel(level) > minLevel {
		return
	}
	fmt.Fprintf(ginkgo.GinkgoWriter, nowStamp()+": "+level+": "+format+"\n", args...)
}

func log(level string, args ...interface{}) {
	if stringToLevel(level) > minLevel {
		return
	}
	fmt.Fprint(ginkgo.GinkgoWriter, nowStamp()+": "+level+": ")
	fmt.Fprint(ginkgo.GinkgoWriter, args...)
	fmt.Fprint(ginkgo.GinkgoWriter, "\n")
}

// Info logs the info
func Info(args ...interface{}) {
	log("[INFO]", args...)
}

// Infof logs the info with arguments
func Infof(format string, args ...interface{}) {
	logf("[INFO]", format, args...)
}

// Warning logs the warning
func Warning(args ...interface{}) {
	log("[WARNING]", args...)
}

// Warningf logs the warning with arguments
func Warningf(format string, args ...interface{}) {
	logf("[WARNING]", format, args...)
}

// Error logs the warning
func Error(args ...interface{}) {
	log("[ERROR]", args...)
}

// Errorf logs the warning with arguments
func Errorf(format string, args ...interface{}) {
	logf("[ERROR]", format, args...)
}

// Tagged Infof logs tagged info with arguments
func TaggedInfof(tag string, format string, args ...interface{}) {
	logf("[INFO]", fmt.Sprintf("[%s] %s", tag, format), args...)
}
