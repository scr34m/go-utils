package log

import (
	"fmt"
	"log"
	"os"
)

func log_(lvl Level, format *string, args ...interface{}) {
	if format == nil {
		log.Printf(lvl.String() + " " + fmt.Sprint(args...))
	} else {
		log.Printf(lvl.String() + " " + fmt.Sprintf(*format, args...))
	}
}

// Fatal is equivalent to l.Critical(fmt.Sprint()) followed by a call to os.Exit(1).
func Fatal(args ...interface{}) {
	log_(CRITICAL, nil, args...)
	os.Exit(1)
}

// Fatalf is equivalent to l.Critical followed by a call to os.Exit(1).
func Fatalf(format string, args ...interface{}) {
	log_(CRITICAL, &format, args...)
	os.Exit(1)
}

// Panic is equivalent to l.Critical(fmt.Sprint()) followed by a call to panic().
func Panic(args ...interface{}) {
	log_(CRITICAL, nil, args...)
	panic(fmt.Sprint(args...))
}

// Panicf is equivalent to l.Critical followed by a call to panic().
func Panicf(format string, args ...interface{}) {
	log_(CRITICAL, &format, args...)
	panic(fmt.Sprintf(format, args...))
}

// Critical logs a message using CRITICAL as log level.
func Critical(args ...interface{}) {
	log_(CRITICAL, nil, args...)
}

// Criticalf logs a message using CRITICAL as log level.
func Criticalf(format string, args ...interface{}) {
	log_(CRITICAL, &format, args...)
}

// Error logs a message using ERROR as log level.
func Error(args ...interface{}) {
	log_(ERROR, nil, args...)
}

// Errorf logs a message using ERROR as log level.
func Errorf(format string, args ...interface{}) {
	log_(ERROR, &format, args...)
}

// Warning logs a message using WARNING as log level.
func Warning(args ...interface{}) {
	log_(WARNING, nil, args...)
}

// Warningf logs a message using WARNING as log level.
func Warningf(format string, args ...interface{}) {
	log_(WARNING, &format, args...)
}

// Notice logs a message using NOTICE as log level.
func Notice(args ...interface{}) {
	log_(NOTICE, nil, args...)
}

// Noticef logs a message using NOTICE as log level.
func Noticef(format string, args ...interface{}) {
	log_(NOTICE, &format, args...)
}

// Info logs a message using INFO as log level.
func Info(args ...interface{}) {
	log_(INFO, nil, args...)
}

// Infof logs a message using INFO as log level.
func Infof(format string, args ...interface{}) {
	log_(INFO, &format, args...)
}

// Debug logs a message using DEBUG as log level.
func Debug(args ...interface{}) {
	log_(DEBUG, nil, args...)
}

// Debugf logs a message using DEBUG as log level.
func Debugf(format string, args ...interface{}) {
	log_(DEBUG, &format, args...)
}
