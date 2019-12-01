// Created on Sat Nov 30 2019
// author: ctacampado

// Package toylog implements a logging package based on
// go's standard log library that features a configurable
// logging level.
package toylog

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
)

// LogLvl is a custom type based on int.
// It is used to differenciate a valid log level argument
// vs some number
type LogLvl int

//logging levels
const (
	DEBUG LogLvl = iota
	INFO
	NOTICE
	WARNING
	ERR
	CRIT
	ALERT
	EMRG
)

// A ToyLog represents a container object for a
// logger and its current configuration parameters
type ToyLog struct {
	name   string      // name of the logger used for labeling logs for easier tracing
	lvl    LogLvl      // logging level
	file   *os.File    // output file where logs are logged
	logger *log.Logger // pointer to the active logger object
}

// NewToyLog initializes ToyLog struct. Valid arguments are
// as follows:
// - type string for the logger name
// - type int for the log level
// - type *os.File for the output file
// If it is an unknown parameter type, it returns nil
func NewToyLog(args ...interface{}) (*ToyLog, error) {

	tl := &ToyLog{
		name: "logger",
		lvl:  0,
		file: os.Stdout,
	}

	for _, arg := range args {
		switch a := arg.(type) {
		case string:
			if a != "" {
				tl.name = fmt.Sprintf("[%s] ", a)
			}
		case LogLvl:
			if a >= 0 && a <= EMRG {
				tl.lvl = a
			}
		case *os.File:
			if a != nil {
				tl.file = a
			}
		default:
			err := fmt.Sprintf("unknown parameter '%v' of type %s\n", arg, reflect.TypeOf(arg))
			return nil, errors.New(err)
		}
	}

	tl.logger = log.New(tl.file, tl.name, log.LstdFlags)
	return tl, nil
}

// Info level is a logging level that uses log.LstdFlags only.
// Using Info log level also shows Debug level logs
func (t *ToyLog) Info(format string, v ...interface{}) {
	if t.lvl >= INFO {
		t.logger.SetFlags(log.LstdFlags)
		newfmt := fmt.Sprintf("[INFO]: " + format)
		t.logger.Output(2, fmt.Sprintf(newfmt, v...))
	}
}

// Debug level is a logging level that uses log.LstdFlags | log.Lshortfile flags.
// This is the default logging level
func (t *ToyLog) Debug(format string, v ...interface{}) {
	t.logger.SetFlags(log.LstdFlags | log.Lshortfile)
	newfmt := fmt.Sprintf("[DEBUG]: " + format)
	t.logger.Output(2, fmt.Sprintf(newfmt, v...))
}

// Error level is a logging level that uses log.LstdFlags | log.Lshortfile | log.Llongfile flags.
// Using this log level will show all log levels less than or equal to ERR level
func (t *ToyLog) Error(format string, v ...interface{}) {
	t.logger.SetFlags(log.LstdFlags | log.Lshortfile | log.Llongfile)
	if t.lvl >= ERR {
		newfmt := fmt.Sprintf("[ERROR]: " + format)
		t.logger.Output(2, fmt.Sprintf(newfmt, v...))
	}
}
